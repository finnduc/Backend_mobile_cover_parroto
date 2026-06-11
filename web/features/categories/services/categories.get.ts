'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { Category } from "@/types/categories.models"
import type { BaseResponse } from "@/types/base-response"

export async function getCategories(): Promise<BaseResponse<Category[]>> {
  return apiFetch<Category[]>("/categories")
}

export async function getAdminCategories(
  page = 1,
  limit = 10
): Promise<BaseResponse<Category[]>> {
  return apiFetch<Category[]>("/admin/categories", {
    withCredentials: true,
    query: { page, limit },
    tags: [CACHE_TAGS.categories],
  })
}
