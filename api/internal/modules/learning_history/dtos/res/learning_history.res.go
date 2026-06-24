package res

type LearningHistoryRes struct {
	ID                     uint   `json:"id"`
	UserID                 string `json:"user_id"`
	LessonID               uint   `json:"lesson_id"`
	CompletedDictation     bool   `json:"completed_dictation"`
	CompletedPronunciation bool   `json:"completed_pronunciation"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type LearningHistorySummaryRes struct {
	CompletedCount  int64 `json:"completed_count"`
	UnfinishedCount int64 `json:"unfinished_count"`
}

type LessonProgressSummaryRes struct {
	Completed   int64 `json:"completed"`
	Uncompleted int64 `json:"uncompleted"`
}
