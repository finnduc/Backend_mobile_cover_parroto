package auth

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/auth/dtos/req"
	"go-cover-parroto/internal/modules/auth/dtos/res"
	"go-cover-parroto/internal/modules/auth/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc services.IAuthService
}

func NewAuthController(svc services.IAuthService) *AuthController {
	return &AuthController{svc: svc}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var body req.RegisterReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.Register(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var body req.LoginReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.Login(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *AuthController) Refresh(c *gin.Context) {
	var body req.RefreshReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}

	result, appErr := ctrl.svc.RefreshToken(c.Request.Context(), body)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(res.LogoutRes{Message: "logged out"}))
}
