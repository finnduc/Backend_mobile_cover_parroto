package chat

import (
	"net/http"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/chat/dtos/req"
	_ "go-cover-parroto/internal/modules/chat/dtos/res"
	"go-cover-parroto/internal/modules/chat/hub"
	"go-cover-parroto/internal/modules/chat/services"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	svc    services.IChatService
	sseHub *hub.SSEHub
}

func NewChatController(svc services.IChatService, sse *hub.SSEHub) *ChatController {
	return &ChatController{
		svc:    svc,
		sseHub: sse,
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

// SendMessage godoc
// @Summary Send a chat message
// @Description Send a message to the global chat
// @Tags chat
// @Accept json
// @Produce json
// @Param body body req.SendMessageReq true "Message content"
// @Success 201 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /chat/messages [post]
// @Security BearerAuth
func (ctrl *ChatController) SendMessage(c *gin.Context) {
	var body req.SendMessageReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	userID := c.Request.Context().Value(enums.ContextKeyUserID).(string)

	if appErr := ctrl.svc.SendMessage(c.Request.Context(), userID, body.Content); appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	c.JSON(http.StatusCreated, response.Success[any](nil))
}
