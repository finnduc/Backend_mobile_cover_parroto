export interface PronunciationAttempt {
  text: string
  overallScore: number
  scores: PronunciationScores
  feedback: string
  words: PronunciationWord[]
  attempt: AttemptInfo
}

export interface PronunciationScores {
  accuracy: number
  fluency: number
  completeness: number
  prosody: number
}

export interface PronunciationWord {
  word: string
  score: number
  feedback: string
  weakPhonemes: WeakPhoneme[]
}

export interface WeakPhoneme {
  phoneme: string
  score: number
}

export interface AttemptInfo {
  id: number
  userId: string
  lessonId: number
  transcriptId: number
  createdAt: string
}

export interface PronunciationProgress {
  userId: string
  lessonId: number
  transcriptId: number
  bestAttemptId: number | null
  bestScore: number | null
  feedback: string
  createdAt: string
  updatedAt: string
}
