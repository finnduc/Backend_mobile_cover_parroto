'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { Bookmark } from "@/types/book-mark.models"

export async function addBookmark(lessonId: number): Promise<BaseResponse<Bookmark>> {
  const res = await apiFetch<Bookmark>(`/bookmarks/${lessonId}`, {
    method: "POST",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.bookmarks)
    refresh()
  }
  return res
}

export async function removeBookmark(lessonId: number): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/bookmarks/${lessonId}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.bookmarks)
    refresh()
  }
  return res
}
