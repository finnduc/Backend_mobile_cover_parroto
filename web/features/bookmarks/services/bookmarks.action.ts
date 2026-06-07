'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { TranscriptBookmark } from "@/types/book-mark.models"

export async function createTranscriptBookmark(
  transcriptId: number,
  note: string
): Promise<BaseResponse<TranscriptBookmark>> {
  const res = await apiFetch<TranscriptBookmark>("/transcript-bookmarks", {
    method: "POST",
    body: { transcriptId, note },
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.bookmarks)
    refresh()
  }
  return res
}

export async function updateTranscriptBookmarkNote(
  id: number,
  note: string
): Promise<BaseResponse<TranscriptBookmark>> {
  const res = await apiFetch<TranscriptBookmark>(`/transcript-bookmarks/${id}`, {
    method: "PATCH",
    body: { note },
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.bookmarks)
    refresh()
  }
  return res
}

export async function deleteTranscriptBookmark(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/transcript-bookmarks/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.bookmarks)
    refresh()
  }
  return res
}
