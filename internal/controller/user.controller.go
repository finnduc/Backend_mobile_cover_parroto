package controller

import (
	"go-cover-parroto/internal/service"
	"go-cover-parroto/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var input service.UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	
	user, err := uc.userService.RegisterService(c.Request.Context(), input.Email, input.Purpose)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	
	if user.ID == 0 {
		response.ErrorResponseData(c, response.CodeUserAlreadyExists, nil)
		return
	}

	response.SuccessReponseData(c, gin.H{"id": user.ID, "email": user.Email, "name": user.Name})
}