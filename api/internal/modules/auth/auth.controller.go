package auth

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/auth/dtos/req"
	_ "go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc services.IAuthService
}

func NewAuthController(svc services.IAuthService) *AuthController {
	return &AuthController{svc: svc}
}

// Sync godoc
// @Summary Sync user with Firebase
// @Description Verify Firebase ID token and create or get user in database
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.SyncReq true "Firebase token"
// @Success 200 {object} response.BaseResponse[res.SyncRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /auth/sync [post]
func (ctrl *AuthController) Sync(c *gin.Context) {
	var body req.SyncReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.SyncUser(c.Request.Context(), body.FirebaseToken)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
