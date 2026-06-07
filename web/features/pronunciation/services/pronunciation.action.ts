'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type {
  PronunciationAttempt,
  PronunciationProgress,
} from "@/types/pronunciation.models"

export async function assessPronunciation(
  formData: FormData
): Promise<BaseResponse<PronunciationAttempt>> {
  const res = await apiFetch<PronunciationAttempt>("/pronunciation-attempts", {
    method: "POST",
    body: formData,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.pronunciation)
    updateTag(CACHE_TAGS.pronunciationProgress)
    refresh()
  }
  return res
}

export async function deletePronunciationAttempt(
  attemptId: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(
    `/pronunciation/attempts/${attemptId}`,
    {
      method: "DELETE",
      withCredentials: true,
    }
  )
  if (!res.error) {
    updateTag(CACHE_TAGS.pronunciation)
    refresh()
  }
  return res
}

export async function updatePronunciationProgress(
  transcriptId: number
): Promise<BaseResponse<PronunciationProgress>> {
  const res = await apiFetch<PronunciationProgress>(
    `/pronunciation/progress/update/${transcriptId}`,
    {
      method: "POST",
      withCredentials: true,
    }
  )
  if (!res.error) {
    updateTag(CACHE_TAGS.pronunciationProgress)
    refresh()
  }
  return res
}
