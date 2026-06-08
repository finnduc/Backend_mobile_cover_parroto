'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { LessonBookmark } from "@/types/lesson-bookmark.models"

export async function toggleLessonBookmark(
  lessonId: number
): Promise<BaseResponse<LessonBookmark | null>> {
  const res = await apiFetch<LessonBookmark | null>(`/lesson-bookmarks/${lessonId}/toggle`, {
    method: "POST",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.lessonBookmarks)
    refresh()
  }
  return res
}

export async function deleteLessonBookmark(
  id: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/lesson-bookmarks/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.lessonBookmarks)
    refresh()
  }
  return res
}
