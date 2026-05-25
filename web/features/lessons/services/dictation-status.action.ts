'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { DictationStatus } from "@/types/dictation-status.models"

export async function postDictationStatus(transcriptId: number, lessonId: number): Promise<BaseResponse<DictationStatus>> {
  const res = await apiFetch<DictationStatus>("/dictation-status", {
    method: "POST",
    body: { transcript_id: transcriptId, lesson_id: lessonId },
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.dictationStatus)
    refresh()
  }
  return res
}
