import type { Transcript } from "@/types/lessons.models"
import type { PronunciationAttempt } from "@/types/pronunciation.models"

export interface ExerciseControlProps {
  transcripts: Transcript[]
  initialCompletedIds: number[]
  lessonId: number
  pronunciationScores?: Map<number, PronunciationAttempt>
}
