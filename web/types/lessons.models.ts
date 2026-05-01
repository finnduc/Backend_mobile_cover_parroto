export interface Lesson {
  id: number
  title: string
  description: string
  thumbnailUrl: string
  videoUrl: string
  duration: number
  level: string
  categoryId: number
  createdAt: string
}

export interface Transcript {
  id: number
  lessonId: number
  sequence: number
  content: string
  phonetic: string
  vietnamese: string
  startTimestamp: number
  endTimestamp: number
}
