package controller

import (
	"errors"
	"go-familytree/internal/models"
	"go-familytree/internal/service"
	pkgerrors "go-familytree/pkg/errors"
	"go-familytree/pkg/response"
	"go-familytree/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ = models.Base{}

type LessonController struct {
	lessonSvc service.ILessonService
}

func NewLessonController(lessonSvc service.ILessonService) *LessonController {
	return &LessonController{lessonSvc: lessonSvc}
}

// CreateLesson godoc
// @Summary Tao bai hoc
// @Description Tao bai hoc moi tu payload frontend (url + transcripts), va co the gan category.
// @Tags lessons
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body service.CreateLessonInput true "Du lieu bai hoc tu frontend"
// @Success 200 {object} response.ResponseData{data=models.Lesson} "Tao bai hoc thanh cong"
// @Router /lessons [post]
func (ctrl *LessonController) CreateLesson(c *gin.Context) {
	var input service.CreateLessonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}

	lesson, err := ctrl.lessonSvc.CreateLesson(c.Request.Context(), input)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, lesson)
}

// ListLessons godoc
// @Summary Danh sach bai hoc
// @Description Lay danh sach bai hoc co phan trang; ho tro loc theo do kho va category.
// @Tags lessons
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (default 10, max 100)"
// @Param level query string false "Filter by level (easy|medium|hard)"
// @Param category_id query int false "Filter by category ID"
// @Success 200 {object} response.ResponseData{data=utils.ListResponse} "Lay danh sach bai hoc thanh cong"
// @Router /lessons [get]
func (ctrl *LessonController) ListLessons(c *gin.Context) {
	var q utils.PaginationQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	level := c.Query("level")
	var categoryID *uint
	if catStr := c.Query("category_id"); catStr != "" {
		v, err := strconv.ParseUint(catStr, 10, 64)
		if err == nil {
			id := uint(v)
			categoryID = &id
		}
	}

	lessons, total, err := ctrl.lessonSvc.ListLessons(c.Request.Context(), level, categoryID, q)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	q.Normalize()
	response.SuccessReponseData(c, utils.ListResponse{
		Items: lessons,
		Meta:  utils.PaginationMeta{Total: total, Page: q.Page, Limit: q.Limit},
	})
}

// GetLesson godoc
// @Summary Chi tiet bai hoc
// @Description Lay thong tin chi tiet cua mot bai hoc theo ID.
// @Tags lessons
// @Security BearerAuth
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} response.ResponseData{data=models.Lesson} "Lay chi tiet bai hoc thanh cong"
// @Router /lessons/{id} [get]
func (ctrl *LessonController) GetLesson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	lesson, err := ctrl.lessonSvc.GetLesson(c.Request.Context(), uint(id))
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, lesson)
}

// GetTranscripts godoc
// @Summary Danh sach transcript cua bai hoc
// @Description Lay toan bo transcript cua bai hoc, sap xep theo thu tu sequence.
// @Tags lessons
// @Security BearerAuth
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} response.ResponseData{data=[]models.Transcript} "Lay transcript thanh cong"
// @Router /lessons/{id}/transcripts [get]
func (ctrl *LessonController) GetTranscripts(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	transcripts, err := ctrl.lessonSvc.GetTranscripts(c.Request.Context(), uint(id))
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, transcripts)
}
