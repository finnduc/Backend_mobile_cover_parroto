export interface LearningHistory {
  id: number
  userId: string
  lessonId: number
  completedDictation: boolean
  completedPronunciation: boolean | null
  createdAt: string
  updatedAt: string
}

export interface LearningHistorySummary {
  completedCount: number
  unfinishedCount: number
}

export interface LearningHistoryLessonSummary {
  completed: number
  uncompleted: number
}
