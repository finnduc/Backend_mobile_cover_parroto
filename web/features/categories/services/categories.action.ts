'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { Category } from "@/types/categories.models"
import type { CreateCategoryDto, UpdateCategoryDto } from "@/features/categories/dtos/category.dto"

export async function createAdminCategory(
  body: CreateCategoryDto
): Promise<BaseResponse<Category>> {
  const res = await apiFetch<Category>("/admin/categories", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.categories)
    refresh()
  }
  return res
}

export async function updateAdminCategory(
  id: number,
  body: UpdateCategoryDto
): Promise<BaseResponse<Category>> {
  const res = await apiFetch<Category>(`/admin/categories/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.categories)
    updateTag(CACHE_TAGS.category(id))
    refresh()
  }
  return res
}

export async function deleteAdminCategory(id: number): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/categories/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.categories)
    updateTag(CACHE_TAGS.category(id))
    refresh()
  }
  return res
}
