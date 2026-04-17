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

func (ctrl *UserController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
		return
	}

	result, appErr := ctrl.svc.GetProfile(c.Request.Context(), userID.(uint))
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}
