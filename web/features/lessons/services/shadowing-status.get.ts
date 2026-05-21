'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { ShadowingStatus } from "@/types/shadowing-status.models"

export async function getShadowingStatus(lessonId: number): Promise<BaseResponse<ShadowingStatus[]>> {
  return apiFetch<ShadowingStatus[]>("/shadowing-status", {
    query: { lesson_id: lessonId },
    withCredentials: true,
    tags: [CACHE_TAGS.shadowingStatus],
  })
}
