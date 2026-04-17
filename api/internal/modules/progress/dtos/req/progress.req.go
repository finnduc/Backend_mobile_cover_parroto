package req

type UpdateProgressReq struct {
	LessonID        uint    `json:"lesson_id" binding:"required"`
	DurationWatched float64 `json:"duration_watched"`
	Completed       bool    `json:"completed"`
}

type GetProgressReq struct {
	LessonID uint `json:"lesson_id" binding:"required"`
}
