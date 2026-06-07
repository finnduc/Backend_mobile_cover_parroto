'use server'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { LearningHistory } from "@/types/learning-history.models"

export async function recordLearningHistory(
  lessonId: number,
  completedDictation: boolean,
  completedPronunciation: boolean | null
): Promise<BaseResponse<LearningHistory>> {
  const res = await apiFetch<LearningHistory>("/learning-history", {
    method: "POST",
    body: {
      lessonId,
      completedDictation,
      completedPronunciation,
    },
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.learningHistory)
    refresh()
  }
  return res
}
