'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { ShadowingStatus } from "@/types/shadowing-status.models"

export async function postShadowingStatus(transcriptId: number, lessonId: number): Promise<BaseResponse<ShadowingStatus>> {
  const res = await apiFetch<ShadowingStatus>("/shadowing-status", {
    method: "POST",
    body: { transcript_id: transcriptId, lesson_id: lessonId },
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.shadowingStatus)
    refresh()
  }
  return res
}
