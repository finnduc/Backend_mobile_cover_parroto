'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { User } from "@/types/users.models"

export type CreateUserInput = Omit<User, "id">
export type UpdateUserInput = Partial<Omit<User, "id">>

export async function createAdminUser(
  body: CreateUserInput
): Promise<BaseResponse<User>> {
  const res = await apiFetch<User>("/admin/users", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    const { updateTag, refresh } = await import("next/cache")
    updateTag(CACHE_TAGS.users)
    refresh()
  }
  return res
}

export async function updateAdminUser(
  id: number,
  body: UpdateUserInput
): Promise<BaseResponse<User>> {
  const res = await apiFetch<User>(`/admin/users/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    const { updateTag, refresh } = await import("next/cache")
    updateTag(CACHE_TAGS.users)
    updateTag(CACHE_TAGS.user(id))
    refresh()
  }
  return res
}

export async function deleteAdminUser(id: number): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/users/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    const { updateTag, refresh } = await import("next/cache")
    updateTag(CACHE_TAGS.users)
    updateTag(CACHE_TAGS.user(id))
    refresh()
  }
  return res
}
