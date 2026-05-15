'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { Category } from "@/types/categories.models"
import { updateTag, refresh } from "next/cache"

export type CreateCategoryInput = Omit<Category, "id">
export type UpdateCategoryInput = Partial<Omit<Category, "id">>

export async function createAdminCategory(
  body: CreateCategoryInput
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
  body: UpdateCategoryInput
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
