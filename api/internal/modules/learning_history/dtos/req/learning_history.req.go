package req

type CreateLearningHistoryReq struct {
	LessonID               uint  `json:"lesson_id" binding:"required"`
	CompletedDictation     *bool `json:"completed_dictation"`
	CompletedPronunciation *bool `json:"completed_pronunciation"`
}
