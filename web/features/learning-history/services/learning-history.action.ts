'use server'
import { apiFetch } from "@/lib/api-fetch"
import type { BaseResponse } from "@/types/base-response"
import type { LearningHistory } from "@/types/learning-history.models"

export async function getHistoryByLesson(
  lessonId: number
): Promise<BaseResponse<LearningHistory>> {
  return apiFetch<LearningHistory>(`/learning-history/${lessonId}`, {
    withCredentials: true,
  })
}

export async function recordHistory(
  lessonId: number,
  durationWatched: number,
  completed: boolean
): Promise<BaseResponse<LearningHistory>> {
  return apiFetch<LearningHistory>("/learning-history", {
    method: "POST",
    body: { lessonId, durationWatched, completed },
    withCredentials: true,
  })
}
