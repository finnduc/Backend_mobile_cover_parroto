package user

import (
	"net/http"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/modules/user/dtos/req"
	_ "go-cover-parroto/internal/modules/user/dtos/res"
	"go-cover-parroto/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type UserAdminController struct {
	svc services.IUserService
}

func NewUserAdminController(svc services.IUserService) *UserAdminController {
	return &UserAdminController{svc: svc}
}

// List godoc
// @Summary List users
// @Description List all users with pagination (admin only)
// @Tags admin-users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.BaseResponse[response.PaginatedResponse[res.UserRes]]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/users [get]
// @Security BearerAuth
func (ctrl *UserAdminController) List(c *gin.Context) {
	var q req.ListUserQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
		return
	}
	result, appErr := ctrl.svc.List(c.Request.Context(), q)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get a user by their ID (admin only)
// @Tags admin-users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.BaseResponse[res.UserRes]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/users/{id} [get]
// @Security BearerAuth
func (ctrl *UserAdminController) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	result, appErr := ctrl.svc.GetByID(c.Request.Context(), id)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success(result))
}

// // Update godoc
// // @Summary Update user
// // @Description Update a user by ID (admin only)
// // @Tags admin-users
// // @Accept json
// // @Produce json
// // @Param id path int true "User ID"
// // @Param body body req.UpdateUserReq true "User data"
// // @Success 200 {object} response.BaseResponse[res.UserRes]
// // @Failure 400 {object} response.BaseResponse[any]
// // @Failure 401 {object} response.BaseResponse[any]
// // @Router /admin/users/{id} [put]
// // @Security BearerAuth
// func (ctrl *UserAdminController) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	if id == "" {
// 		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
// 		return
// 	}
// 	var body req.UpdateUserReq
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest(err.Error())))
// 		return
// 	}
// 	result, appErr := ctrl.svc.Update(c.Request.Context(), id, body)
// 	if appErr != nil {
// 		c.JSON(appErr.Code, response.Fail(appErr))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.Success(result))
// }

// Delete godoc
// @Summary Delete user
// @Description Delete a user by ID (admin only)
// @Tags admin-users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.BaseResponse[any]
// @Failure 400 {object} response.BaseResponse[any]
// @Failure 401 {object} response.BaseResponse[any]
// @Router /admin/users/{id} [delete]
// @Security BearerAuth
func (ctrl *UserAdminController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.Fail(response.BadRequest("invalid id")))
		return
	}
	appErr := ctrl.svc.Delete(c.Request.Context(), id)
	if appErr != nil {
		c.JSON(appErr.Code, response.Fail(appErr))
		return
	}
	c.JSON(http.StatusOK, response.Success("user deleted"))
}
