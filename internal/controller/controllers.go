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
// @Summary Create a new attempt
// @Description Start a new learning attempt for a specific lesson
// @Tags attempts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{lesson_id=uint} true "Attempt request"
// @Success 200 {object} response.ResponseData{data=models.Attempt} "Attempt created"
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
// @Summary Get attempt details
// @Description Get detailed information about a specific attempt by ID
// @Tags attempts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Attempt ID"
// @Success 200 {object} response.ResponseData{data=models.Attempt} "Attempt details"
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
// @Summary Submit a single answer
// @Description Submit an answer for a specific transcript and calculate score
// @Tags answers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body service.SubmitAnswerInput true "Answer info"
// @Success 200 {object} response.ResponseData{data=service.AnswerResult} "Submit successful"
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
// @Summary Bulk submit answers
// @Description Submit multiple answers at once (e.g. for sync)
// @Tags answers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body service.BulkSubmitInput true "Bulk answers info"
// @Success 200 {object} response.ResponseData{data=service.BulkAnswerResult} "Bulk submit successful"
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
// @Summary Get user progress for a lesson
// @Description Get current progress (last sequence, average score) for a specific lesson
// @Tags progress
// @Security BearerAuth
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Success 200 {object} response.ResponseData{data=models.UserProgress} "User progress"
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
// @Summary Toggle bookmark
// @Description Add or remove a lesson from user bookmarks
// @Tags bookmarks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body object{lesson_id=uint} true "Bookmark request"
// @Success 200 {object} response.ResponseData{data=object{bookmarked=bool}} "Toggle successful"
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
// @Summary List bookmarked lessons
// @Description Get a paginated list of lessons marked as favorite by the user
// @Tags bookmarks
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (default 10, max 100)"
// @Success 200 {object} response.ResponseData{data=utils.ListResponse} "List of bookmarked lessons"
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
// @Summary List all categories
// @Description Get a list of all lesson categories
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.ResponseData{data=[]models.Category} "List of categories"
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
// @Summary Get lessons by category
// @Description Get a paginated list of lessons belonging to a specific category
// @Tags categories
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Page size (default 10, max 100)"
// @Success 200 {object} response.ResponseData{data=utils.ListResponse} "List of lessons"
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
