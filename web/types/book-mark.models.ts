export interface Bookmark {
  lessionId: number
  userId: string
  createdAt: string
  lesson: LessonInfo
}
export interface LessonInfo {
  id: number
  title: string
  thumbnailUrl: string
  level: string
  duration: number
}