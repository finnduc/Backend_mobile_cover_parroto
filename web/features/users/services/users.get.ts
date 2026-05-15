'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { User } from "@/types/users.models"
import type { BaseResponse } from "@/types/base-response"

export async function getUsers(): Promise<BaseResponse<User[]>> {
  return apiFetch<User[]>("/users")
}

export async function getAdminUsers(): Promise<BaseResponse<User[]>> {
  return apiFetch<User[]>("/admin/users", {
    withCredentials: true,
    tags: [CACHE_TAGS.users],
  })
}

export async function getAdminUser(id: number): Promise<BaseResponse<User>> {
  return apiFetch<User>(`/admin/users/${id}`, {
    withCredentials: true,
    tags: [CACHE_TAGS.user(id)],
  })
}
