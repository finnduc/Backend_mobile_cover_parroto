'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { VocabularyCategory, VocabularyDeck, VocabularyItem } from "@/types/vocabulary.models"
import type { CreateVocabularyCategoryDto, UpdateVocabularyCategoryDto, CreateVocabularyDeckDto, UpdateVocabularyDeckDto, CreateVocabularyItemDto, UpdateVocabularyItemDto } from "@/features/vocabulary/dtos/vocabulary.dto"

export async function createVocabularyCategory(
  body: CreateVocabularyCategoryDto
): Promise<BaseResponse<VocabularyCategory>> {
  const res = await apiFetch<VocabularyCategory>("/admin/vocabulary-categories", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyCategories)
    refresh()
  }
  return res
}

export async function updateVocabularyCategory(
  id: number,
  body: UpdateVocabularyCategoryDto
): Promise<BaseResponse<VocabularyCategory>> {
  const res = await apiFetch<VocabularyCategory>(`/admin/vocabulary-categories/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyCategories)
    updateTag(CACHE_TAGS.vocabularyCategory(id))
    refresh()
  }
  return res
}

export async function deleteVocabularyCategory(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/vocabulary-categories/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyCategories)
    updateTag(CACHE_TAGS.vocabularyCategory(id))
    refresh()
  }
  return res
}

export async function createVocabularyDeck(
  body: CreateVocabularyDeckDto
): Promise<BaseResponse<VocabularyDeck>> {
  const res = await apiFetch<VocabularyDeck>("/vocabulary-decks", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    refresh()
  }
  return res
}

export async function updateVocabularyDeck(
  id: number,
  body: UpdateVocabularyDeckDto
): Promise<BaseResponse<VocabularyDeck>> {
  const res = await apiFetch<VocabularyDeck>(`/vocabulary-decks/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    updateTag(CACHE_TAGS.vocabularyDeck(id))
    refresh()
  }
  return res
}

export async function deleteVocabularyDeck(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/vocabulary-decks/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    updateTag(CACHE_TAGS.vocabularyDeck(id))
    refresh()
  }
  return res
}

export async function createVocabularyItem(
  deckId: number,
  body: CreateVocabularyItemDto
): Promise<BaseResponse<VocabularyItem>> {
  const res = await apiFetch<VocabularyItem>(`/vocabulary-decks/${deckId}/items`, {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    updateTag(CACHE_TAGS.vocabularyDeck(deckId))
    refresh()
  }
  return res
}

export async function updateVocabularyItem(
  id: number,
  body: UpdateVocabularyItemDto
): Promise<BaseResponse<VocabularyItem>> {
  const res = await apiFetch<VocabularyItem>(`/vocabulary-items/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    refresh()
  }
  return res
}

export async function deleteVocabularyItem(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/vocabulary-items/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    refresh()
  }
  return res
}

export async function createAdminVocabularyDeck(
  body: CreateVocabularyDeckDto
): Promise<BaseResponse<VocabularyDeck>> {
  const res = await apiFetch<VocabularyDeck>("/admin/vocabulary-decks", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    refresh()
  }
  return res
}

export async function updateAdminVocabularyDeck(
  id: number,
  body: UpdateVocabularyDeckDto
): Promise<BaseResponse<VocabularyDeck>> {
  const res = await apiFetch<VocabularyDeck>(`/admin/vocabulary-decks/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    updateTag(CACHE_TAGS.vocabularyDeck(id))
    refresh()
  }
  return res
}

export async function deleteAdminVocabularyDeck(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/vocabulary-decks/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyDecks)
    updateTag(CACHE_TAGS.vocabularyDeck(id))
    refresh()
  }
  return res
}

export async function createAdminVocabularyItem(
  deckId: number,
  body: CreateVocabularyItemDto
): Promise<BaseResponse<VocabularyItem>> {
  const res = await apiFetch<VocabularyItem>(`/admin/vocabulary-decks/${deckId}/items`, {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    updateTag(CACHE_TAGS.vocabularyDeck(deckId))
    refresh()
  }
  return res
}

export async function updateAdminVocabularyItem(
  id: number,
  body: UpdateVocabularyItemDto
): Promise<BaseResponse<VocabularyItem>> {
  const res = await apiFetch<VocabularyItem>(`/admin/vocabulary-items/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    refresh()
  }
  return res
}

export async function deleteAdminVocabularyItem(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/vocabulary-items/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.vocabularyItems)
    refresh()
  }
  return res
}
