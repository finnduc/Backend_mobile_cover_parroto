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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.RegisterReq true "Register request"
// @Success 200 {object} response.BaseResponse[res.RegisterRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 409 {object} response.BaseResponse[any]
// @Router /auth/register [post]
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

// Login godoc
// @Summary Login user
// @Description Authenticate user and return access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.LoginReq true "Login request"
// @Success 200 {object} response.BaseResponse[res.LoginRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /auth/login [post]
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

// Refresh godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body req.RefreshReq true "Refresh request"
// @Success 200 {object} response.BaseResponse[res.RefreshRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /auth/refresh [post]
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

// Logout godoc
// @Summary Logout user
// @Description Logout current user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[res.LogoutRes]
// @Router /auth/logout [post]
// @Security BearerAuth
func (ctrl *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(res.LogoutRes{Message: "logged out"}))
}
