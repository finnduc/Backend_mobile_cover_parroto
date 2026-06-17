import type { Transcript } from "@/types/lessons.models"

export interface ExerciseControlProps {
  transcripts: Transcript[]
  initialCompletedIds: number[]
  lessonId: number
  isAuthenticated?: boolean
}
