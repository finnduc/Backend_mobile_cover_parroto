package controller

import (
	"errors"
	"go-familytree/internal/service"
	pkgerrors "go-familytree/pkg/errors"
	"go-familytree/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSvc service.IAuthService
}

func NewAuthController(authSvc service.IAuthService) *AuthController {
	return &AuthController{authSvc: authSvc}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account with email, password and name
// @Tags auth
// @Accept json
// @Produce json
// @Param body body service.RegisterInput true "Registration info"
// @Success 200 {object} response.ResponseData "User created"
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var input service.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	user, err := ctrl.authSvc.Register(c.Request.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrConflict):
			response.ErrorResponseData(c, response.CodeUserAlreadyExists, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, gin.H{"id": user.ID, "email": user.Email, "name": user.Name})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return access + refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param body body service.LoginInput true "Login credentials"
// @Success 200 {object} response.ResponseData{data=service.TokenPair} "Login successful"
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var input service.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	pair, err := ctrl.authSvc.Login(c.Request.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeUserNotFound, nil)
		case errors.Is(err, pkgerrors.ErrUnauthorized):
			response.ErrorResponseData(c, response.CodePasswordWrong, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, pair)
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate the current access token (blacklist) and delete refresh token
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.ResponseData "Logout successful"
// @Router /auth/logout [post]
func (ctrl *AuthController) Logout(c *gin.Context) {
	userID := c.GetUint("user_id")
	rawToken := c.GetString("raw_token")
	_ = ctrl.authSvc.Logout(c.Request.Context(), rawToken, userID)
	response.SuccessReponseData(c, nil)
}

// Refresh godoc
// @Summary Refresh access token
// @Description Exchange a valid refresh token for a new access + refresh token pair
// @Tags auth
// @Accept json
// @Produce json
// @Param body body object{refresh_token=string,user_id=uint} true "Refresh request"
// @Success 200 {object} response.ResponseData{data=service.TokenPair} "Tokens refreshed"
// @Router /auth/refresh [post]
func (ctrl *AuthController) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		UserID       uint   `json:"user_id"       binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	pair, err := ctrl.authSvc.Refresh(c.Request.Context(), req.RefreshToken, req.UserID)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrUnauthorized):
			response.ErrorResponseData(c, response.CodeUnauthorized, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, pair)
}
