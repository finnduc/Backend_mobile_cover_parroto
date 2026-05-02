package user

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	_ "go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{ svc services.IUserService }

func NewUserController(svc services.IUserService) *UserController {
	return &UserController{svc: svc}
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[res.UserRes]
// @Failure 401 {object} response.BaseResponse[any]
// @Failure 500 {object} response.BaseResponse[any]
// @Router /users/me [get]
// @Security BearerAuth
func (ctrl *UserController) GetProfile(c *gin.Context) {
	result, appErr := ctrl.svc.GetProfile(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
