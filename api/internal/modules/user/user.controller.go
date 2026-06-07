package user

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/auth/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc services.IAuthService
}

func NewUserController(svc services.IAuthService) *UserController {
	return &UserController{svc: svc}
}

// GetProfile godoc
// @Summary Get current user profile from Clerk
// @Tags user
// @Success 200 {object} response.BaseResponse[any]
// @Router /user/profile [get]
// @Security BearerAuth
func (ctrl *UserController) GetProfile(c *gin.Context) {
	user, appErr := ctrl.svc.GetUserProfile(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(user))
}
