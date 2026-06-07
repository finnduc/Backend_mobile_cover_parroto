'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"

export interface UserProfile {
  id: string
  email: string
  name: string
  userRole: string
  avatarUrl: string
  createdAt: string
}

export async function getUserProfile(): Promise<BaseResponse<UserProfile>> {
  return apiFetch<UserProfile>("/user/profile", {
    withCredentials: true,
    tags: [CACHE_TAGS.users],
  })
}
