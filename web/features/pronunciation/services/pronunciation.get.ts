'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { PronunciationProgress, PronunciationProgressDetail } from "@/types/pronunciation.models"

export async function getPronunciationProgress(
  lessonId: number
): Promise<BaseResponse<PronunciationProgress[]>> {
  return apiFetch<PronunciationProgress[]>(
    `/pronunciation/progress/${lessonId}`,
    {
      withCredentials: true,
      tags: [CACHE_TAGS.pronunciationProgress],
    }
  )
}

export async function getPronunciationProgressDetail(
  lessonId: number
): Promise<BaseResponse<PronunciationProgressDetail[]>> {
  return apiFetch<PronunciationProgressDetail[]>(
    `/pronunciation/progress/${lessonId}/detail`,
    {
      withCredentials: true,
      tags: [CACHE_TAGS.pronunciationProgress],
    }
  )
}
