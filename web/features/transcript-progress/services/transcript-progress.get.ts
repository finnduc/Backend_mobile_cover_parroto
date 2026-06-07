'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { TranscriptProgress } from "@/types/transcript-progress.models"

export async function getTranscriptProgress(
  lessonId: number
): Promise<BaseResponse<TranscriptProgress[]>> {
  return apiFetch<TranscriptProgress[]>(
    `/transcript-progress/${lessonId}`,
    {
      withCredentials: true,
      tags: [CACHE_TAGS.transcriptProgress],
    }
  )
}
