'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { VocabularyCategory, VocabularyDeck, VocabularyItem } from "@/types/vocabulary.models"

export async function getVocabularyCategories(): Promise<BaseResponse<VocabularyCategory[]>> {
  return apiFetch<VocabularyCategory[]>("/vocabulary-categories")
}

export async function getAdminVocabularyCategories(): Promise<BaseResponse<VocabularyCategory[]>> {
  return apiFetch<VocabularyCategory[]>("/admin/vocabulary-categories", {
    withCredentials: true,
    tags: [CACHE_TAGS.vocabularyCategories],
  })
}

export async function getAdminVocabularyCategory(id: number): Promise<BaseResponse<VocabularyCategory>> {
  return apiFetch<VocabularyCategory>(`/admin/vocabulary-categories/${id}`, {
    withCredentials: true,
    tags: [CACHE_TAGS.vocabularyCategory(id)],
  })
}

export async function getSystemVocabularyDecks(
  categoryId?: number,
  level?: string
): Promise<BaseResponse<VocabularyDeck[]>> {
  const query: Record<string, unknown> = {}
  if (categoryId) query.categoryId = categoryId
  if (level) query.level = level
  return apiFetch<VocabularyDeck[]>("/vocabulary-system-decks", { query })
}

export async function getUserVocabularyDecks(): Promise<BaseResponse<VocabularyDeck[]>> {
  return apiFetch<VocabularyDeck[]>("/vocabulary-decks", {
    withCredentials: true,
    tags: [CACHE_TAGS.vocabularyDecks],
  })
}

export async function getVocabularyDeck(deckId: number): Promise<BaseResponse<VocabularyDeck>> {
  return apiFetch<VocabularyDeck>(`/vocabulary-decks/${deckId}`)
}

export async function getAdminVocabularyDecks(): Promise<BaseResponse<VocabularyDeck[]>> {
  return apiFetch<VocabularyDeck[]>("/admin/vocabulary-decks", {
    withCredentials: true,
    tags: [CACHE_TAGS.vocabularyDecks],
  })
}

export async function getVocabularyItems(
  deckId: number,
  page = 1,
  limit = 100
): Promise<BaseResponse<VocabularyItem[]>> {
  return apiFetch<VocabularyItem[]>(`/vocabulary-decks/${deckId}/items`, {
    query: { page, limit },
  })
}

export async function getAdminVocabularyItems(
  deckId: number,
  page = 1,
  limit = 100
): Promise<BaseResponse<VocabularyItem[]>> {
  return apiFetch<VocabularyItem[]>(`/admin/vocabulary-decks/${deckId}/items`, {
    withCredentials: true,
    query: { page, limit },
    tags: [CACHE_TAGS.vocabularyItems],
  })
}
