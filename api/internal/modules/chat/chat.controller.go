package chat

import (
	"context"
	"net/http"

	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/policy"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/chat/dtos/req"
	_ "go-cover-parroto/internal/modules/chat/dtos/res"
	"go-cover-parroto/internal/modules/chat/hub"
	"go-cover-parroto/internal/modules/chat/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatController struct {
	svc      services.IChatService
	hub      *hub.Hub
	upgrader websocket.Upgrader
}

func NewChatController(svc services.IChatService, h *hub.Hub) *ChatController {
	return &ChatController{
		svc: svc,
		hub: h,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// allow all origins; tighten if needed
				return true
			},
		},
	}
}

// GetHistory godoc
// @Summary List global chat history
// @Description Get global chat messages with cursor-based pagination (newest first)
// @Tags chat
// @Accept json
// @Produce json
// @Param before_id query int false "Return messages older than this id"
// @Param limit query int false "Number of messages to return (max 50)" default(20)
// @Success 200 {object} response.BaseResponse[res.HistoryRes]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /chat/messages [get]
// @Security BearerAuth
func (ctrl *ChatController) GetHistory(c *gin.Context) {
	var q req.ListMessagesQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.GetHistory(c.Request.Context(), q.BeforeID, q.Limit)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// Connect godoc
// @Summary Connect to global chat WebSocket
// @Description Upgrade HTTP connection to a WebSocket channel for global chat
// @Tags chat
// @Success 101 {string} string "Switching Protocols"
// @Failure 401 {object} response.BaseResponse[any]
// @Router /chat/ws [get]
// @Security BearerAuth
func (ctrl *ChatController) Connect(c *gin.Context) {
	userID, appErr := policy.GetUserID(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.S().Warnw("websocket upgrade failed", "error", err)
		return
	}

	// Sync user info from Clerk → local users table so message metadata
	// (name, avatar) is available when broadcasting / paginating history.
	ctrl.svc.SyncUser(c.Request.Context(), userID)

	// WS lifecycle outlives the HTTP request context (which is canceled as soon
	// as the gin handler returns after the upgrade), so use a fresh context for
	// background DB writes triggered by incoming frames.
	client := hub.NewClient(ctrl.hub, conn, userID, ctrl.svc.SendMessage, context.Background())
	client.Start()
}
