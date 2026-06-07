'use server'

import { apiFetch } from "@/lib/api-fetch"
import type { BaseResponse } from "@/types/base-response"
import type { UserProfile } from "@/features/profile/services/profile.get"

export async function completeSignUp() {
  return apiFetch<void>("/auth/complete-signup", {
    method: "POST",
    withCredentials: true,
  })
}

export async function syncAuth(): Promise<BaseResponse<UserProfile>> {
  return apiFetch<UserProfile>("/auth/sync", {
    method: "POST",
    withCredentials: true,
  })
}
