package user

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc services.IUserService
}

func NewUserController(svc services.IUserService) *UserController {
	return &UserController{svc: svc}
}

// GetProfile godoc
// @Summary Get current user profile
// @Tags user
// @Produce json
// @Success 200 {object} response.BaseResponse[models.User]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 404 {object} response.BaseResponse[any]
// @Router /user/profile [get]
// @Security BearerAuth
func (ctrl *UserController) GetProfile(c *gin.Context) {
	user, appErr := ctrl.svc.GetProfile(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(user))
}
