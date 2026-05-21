'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { DictationStatus } from "@/types/dictation-status.models"

export async function getDictationStatus(lessonId: number): Promise<BaseResponse<DictationStatus[]>> {
  return apiFetch<DictationStatus[]>("/dictation-status", {
    query: { lesson_id: lessonId },
    withCredentials: true,
    tags: [CACHE_TAGS.dictationStatus],
  })
}
