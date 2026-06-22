package user

import (
	"net/http"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	authreq "go-cover-parroto/internal/modules/auth/dtos/req"
	"go-cover-parroto/internal/modules/auth/services"
	"go-cover-parroto/internal/utils"

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
	userID, appErr := utils.GetFromContext[string](c.Request.Context(), enums.ContextKeyUserID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	user, appErr := ctrl.svc.GetUserProfile(c.Request.Context(), userID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(user))
}

func (ctrl *UserController) UpdateProfile(c *gin.Context) {
	userID, appErr := utils.GetFromContext[string](c.Request.Context(), enums.ContextKeyUserID)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	var body authreq.UpdateProfileReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	user, appErr := ctrl.svc.UpdateUserProfile(c.Request.Context(), userID, body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(user))
}
