'server-only'
import { apiFetch } from "@/lib/api-fetch"
import type { BaseResponse } from "@/types/base-response"
import type { LearningHistory } from "@/types/learning-history.models"

export async function getLearningHistory(
  page = 1,
  limit = 20
): Promise<BaseResponse<LearningHistory[]>> {
  return apiFetch<LearningHistory[]>("/learning-history", {
    query: { page, limit },
    withCredentials: true,
  })
}

export async function getLearningHistoryByLesson(
  lessonId: number
): Promise<BaseResponse<LearningHistory>> {
  return apiFetch<LearningHistory>(`/learning-history/${lessonId}`, {
    withCredentials: true,
  })
}
