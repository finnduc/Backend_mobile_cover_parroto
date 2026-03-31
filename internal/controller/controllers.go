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

// ─── Attempt Controller ──────────────────────────────────────────────────────

type AttemptController struct {
	attemptSvc service.IAttemptService
}

func NewAttemptController(attemptSvc service.IAttemptService) *AttemptController {
	return &AttemptController{attemptSvc: attemptSvc}
}

// Create godoc
// @Summary Tao lan luyen tap moi
// @Description Bat dau mot lan luyen tap moi cho bai hoc duoc chon.
// @Tags attempts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{lesson_id=uint} true "Attempt request"
// @Success 200 {object} response.ResponseData{data=models.Attempt} "Tao attempt thanh cong"
// @Router /attempts [post]
func (ctrl *AttemptController) Create(c *gin.Context) {
	var req struct {
		LessonID uint `json:"lesson_id" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	userID := c.GetUint("user_id")
	attempt, err := ctrl.attemptSvc.CreateAttempt(c.Request.Context(), userID, req.LessonID)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, attempt)
}

// Get godoc
// @Summary Chi tiet lan luyen tap
// @Description Lay thong tin chi tiet cua mot attempt theo ID.
// @Tags attempts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Attempt ID"
// @Success 200 {object} response.ResponseData{data=models.Attempt} "Lay chi tiet attempt thanh cong"
// @Router /attempts/{id} [get]
func (ctrl *AttemptController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	attempt, err := ctrl.attemptSvc.GetAttempt(c.Request.Context(), uint(id))
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, attempt)
}

// ─── Answer Controller ───────────────────────────────────────────────────────

type AnswerController struct {
	answerSvc service.IAnswerService
}

func NewAnswerController(answerSvc service.IAnswerService) *AnswerController {
	return &AnswerController{answerSvc: answerSvc}
}

// Submit godoc
// @Summary Nop 1 cau tra loi
// @Description Gui cau tra loi cho mot transcript cu the va tinh diem.
// @Tags answers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body service.SubmitAnswerInput true "Answer info"
// @Success 200 {object} response.ResponseData{data=service.AnswerResult} "Nop cau tra loi thanh cong"
// @Router /answers [post]
func (ctrl *AnswerController) Submit(c *gin.Context) {
	var input service.SubmitAnswerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	userID := c.GetUint("user_id")
	result, err := ctrl.answerSvc.SubmitAnswer(c.Request.Context(), userID, input)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		case errors.Is(err, pkgerrors.ErrInvalidInput):
			response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, result)
}

// BulkSubmit godoc
// @Summary Nop nhieu cau tra loi
// @Description Gui nhieu cau tra loi cung luc (phu hop dong bo du lieu).
// @Tags answers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body service.BulkSubmitInput true "Bulk answers info"
// @Success 200 {object} response.ResponseData{data=service.BulkAnswerResult} "Nop nhieu cau tra loi thanh cong"
// @Router /answers/bulk [post]
func (ctrl *AnswerController) BulkSubmit(c *gin.Context) {
	var input service.BulkSubmitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	userID := c.GetUint("user_id")
	result, err := ctrl.answerSvc.BulkSubmit(c.Request.Context(), userID, input)
	if err != nil {
		switch {
		case errors.Is(err, pkgerrors.ErrNotFound):
			response.ErrorResponseData(c, response.CodeNotFound, nil)
		case errors.Is(err, pkgerrors.ErrInvalidInput):
			response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		default:
			response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		}
		return
	}
	response.SuccessReponseData(c, result)
}

// ─── Progress Controller ─────────────────────────────────────────────────────

type ProgressController struct {
	progressSvc service.IProgressService
}

func NewProgressController(progressSvc service.IProgressService) *ProgressController {
	return &ProgressController{progressSvc: progressSvc}
}

// Get godoc
// @Summary Tien do hoc cua nguoi dung
// @Description Lay tien do hien tai cua nguoi dung tren bai hoc (sequence cuoi, diem trung binh).
// @Tags progress
// @Security BearerAuth
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Success 200 {object} response.ResponseData{data=models.UserProgress} "Lay tien do thanh cong"
// @Router /progress/{lesson_id} [get]
func (ctrl *ProgressController) Get(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lesson_id"), 10, 64)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	userID := c.GetUint("user_id")
	progress, err := ctrl.progressSvc.GetProgress(c.Request.Context(), userID, uint(lessonID))
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, progress)
}

// ─── Bookmark Controller ─────────────────────────────────────────────────────

type BookmarkController struct {
	bookmarkSvc service.IBookmarkService
}

func NewBookmarkController(bookmarkSvc service.IBookmarkService) *BookmarkController {
	return &BookmarkController{bookmarkSvc: bookmarkSvc}
}

// Toggle godoc
// @Summary Them/bo danh dau yeu thich
// @Description Bat/tat trang thai bookmark cho bai hoc cua nguoi dung.
// @Tags bookmarks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{lesson_id=uint} true "Bookmark request"
// @Success 200 {object} response.ResponseData{data=object{bookmarked=bool}} "Cap nhat bookmark thanh cong"
// @Router /bookmarks [post]
func (ctrl *BookmarkController) Toggle(c *gin.Context) {
	var req struct {
		LessonID uint `json:"lesson_id" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	userID := c.GetUint("user_id")
	bookmarked, err := ctrl.bookmarkSvc.Toggle(c.Request.Context(), userID, req.LessonID)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, gin.H{"bookmarked": bookmarked})
}

// List godoc
// @Summary Danh sach bai hoc da bookmark
// @Description Lay danh sach bai hoc da danh dau yeu thich co phan trang.
// @Tags bookmarks
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (default 10, max 100)"
// @Success 200 {object} response.ResponseData{data=utils.ListResponse} "Lay danh sach bookmark thanh cong"
// @Router /bookmarks [get]
func (ctrl *BookmarkController) List(c *gin.Context) {
	userID := c.GetUint("user_id")
	q := pagination(c)
	lessons, total, err := ctrl.bookmarkSvc.ListBookmarks(c.Request.Context(), userID, q)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, utils.ListResponse{
		Items: lessons,
		Meta:  utils.PaginationMeta{Total: total, Page: q.Page, Limit: q.Limit},
	})
}

// ─── Category Controller ─────────────────────────────────────────────────────

type CategoryController struct {
	categorySvc service.ICategoryService
}

func NewCategoryController(categorySvc service.ICategoryService) *CategoryController {
	return &CategoryController{categorySvc: categorySvc}
}

// List godoc
// @Summary Danh sach category
// @Description Lay toan bo danh muc bai hoc.
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.ResponseData{data=[]models.Category} "Lay danh sach category thanh cong"
// @Router /categories [get]
func (ctrl *CategoryController) List(c *gin.Context) {
	cats, err := ctrl.categorySvc.ListCategories(c.Request.Context())
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, cats)
}

// GetLessons godoc
// @Summary Danh sach bai hoc theo category
// @Description Lay danh sach bai hoc theo category co phan trang.
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (default 10, max 100)"
// @Success 200 {object} response.ResponseData{data=utils.ListResponse} "Lay danh sach bai hoc theo category thanh cong"
// @Router /categories/{id}/lessons [get]
func (ctrl *CategoryController) GetLessons(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInvalidParams, nil)
		return
	}
	q := pagination(c)
	lessons, total, err := ctrl.categorySvc.GetLessonsByCategory(c.Request.Context(), uint(id), q)
	if err != nil {
		response.ErrorResponseData(c, response.CodeInternalServerError, nil)
		return
	}
	response.SuccessReponseData(c, utils.ListResponse{
		Items: lessons,
		Meta:  utils.PaginationMeta{Total: total, Page: q.Page, Limit: q.Limit},
	})
}
