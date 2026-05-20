import type { VocabularyCategory, VocabularyDeck, VocabularyItem } from "@/types/vocabulary.models"

export type CreateVocabularyCategoryDto = Omit<VocabularyCategory, "id" | "createdAt" | "updatedAt">
export type UpdateVocabularyCategoryDto = Partial<CreateVocabularyCategoryDto>

export type CreateVocabularyDeckDto = Omit<VocabularyDeck, "id" | "userId" | "isDefault" | "createdAt" | "updatedAt" | "category">
export type UpdateVocabularyDeckDto = Partial<Omit<CreateVocabularyDeckDto, "categoryId"> & { categoryId?: number | null }>

export type CreateVocabularyItemDto = Omit<VocabularyItem, "id" | "deckId" | "lessonId" | "transcriptId" | "createdAt" | "updatedAt">
export type UpdateVocabularyItemDto = Partial<CreateVocabularyItemDto>
