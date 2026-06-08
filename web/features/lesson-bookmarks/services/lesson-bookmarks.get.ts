'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { LessonBookmark } from "@/types/lesson-bookmark.models"
import type { BaseResponse } from "@/types/base-response"

export async function getLessonBookmarks(
  page = 1,
  limit = 100
): Promise<BaseResponse<LessonBookmark[]>> {
  return apiFetch<LessonBookmark[]>("/lesson-bookmarks", {
    query: { page, limit },
    withCredentials: true,
    tags: [CACHE_TAGS.lessonBookmarks],
  })
}
