'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { TranscriptProgress } from "@/types/transcript-progress.models"

export async function createTranscriptProgress(
  lessonId: number,
  transcriptId: number
): Promise<BaseResponse<TranscriptProgress>> {
  const res = await apiFetch<TranscriptProgress>(
    `/transcript-progress/${lessonId}`,
    {
      method: "POST",
      body: { transcriptId },
      withCredentials: true,
    }
  )
  if (!res.error) {
    updateTag(CACHE_TAGS.transcriptProgress)
    refresh()
  }
  return res
}
