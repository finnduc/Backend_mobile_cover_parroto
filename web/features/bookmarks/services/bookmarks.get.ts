'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { TranscriptBookmark } from "@/types/book-mark.models"
import type { BaseResponse } from "@/types/base-response"

export async function getTranscriptBookmarks(
  lessonId: number,
  page = 1,
  limit = 20
): Promise<BaseResponse<TranscriptBookmark[]>> {
  return apiFetch<TranscriptBookmark[]>(`/transcript-bookmarks/${lessonId}`, {
    query: { page, limit },
    withCredentials: true,
    tags: [CACHE_TAGS.bookmarks],
  })
}
