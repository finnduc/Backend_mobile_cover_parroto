package req

type RecordHistoryReq struct {
	LessonID        uint    `json:"lesson_id" binding:"required"`
	DurationWatched float64 `json:"duration_watched" binding:"required,min=0"`
	Completed       bool    `json:"completed"`
}
