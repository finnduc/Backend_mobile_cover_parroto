package auth

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/auth/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc services.IAuthService
}

func NewAuthController(svc services.IAuthService) *AuthController {
	return &AuthController{svc: svc}
}

// CompleteSignUp godoc
// @Summary Finalize user registration
// @Description Sets initial user claims and roles after the first Firebase authentication
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[any] "Successfully finalized profile"
// @Failure 401 {object} response.BaseResponse[any] "Unauthorized - Valid session required"
// @Failure 500 {object} response.BaseResponse[any] "Internal Server Error"
// @Router /auth/complete-signup [post]
// @Security BearerAuth
func (ctrl *AuthController) Complete(c *gin.Context) {
	appErr := ctrl.svc.CompleteSignUp(c.Request.Context())

	if appErr != nil {

		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}

	c.JSON(http.StatusOK, response.Success("User profile finalized successfully"))
}

// Sync godoc
// @Summary Sync/get current Clerk user
// @Tags auth
// @Success 200 {object} response.BaseResponse[res.AuthUserRes]
// @Router /auth/sync [post]
// @Security BearerAuth
func (ctrl *AuthController) Sync(c *gin.Context) {
	user, appErr := ctrl.svc.SyncUser(c.Request.Context())
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(user))
}
