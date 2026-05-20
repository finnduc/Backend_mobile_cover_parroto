export interface VocabularyCategory {
  id: number
  name: string
  description: string
  createdAt: string
  updatedAt: string
}

export interface VocabularyDeck {
  id: number
  userId: string | null
  categoryId: number | null
  name: string
  description: string
  thumbnailUrl: string
  level: string
  isDefault: boolean
  createdAt: string
  updatedAt: string
  category?: VocabularyCategory
}

export interface VocabularyItem {
  id: number
  deckId: number
  lessonId: number | null
  transcriptId: number | null
  phrase: string
  normalizedPhrase: string
  meaning: string
  exampleSentence: string
  note: string
  createdAt: string
  updatedAt: string
}
