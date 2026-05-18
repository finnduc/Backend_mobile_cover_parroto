'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { Bookmark } from "@/types/book-mark.models"
import type { BaseResponse } from "@/types/base-response"

export async function getBookmarks(): Promise<BaseResponse<Bookmark[]>> {
  return apiFetch<Bookmark[]>("/bookmarks", {
    withCredentials: true,
    tags: [CACHE_TAGS.bookmarks],
  })
}
