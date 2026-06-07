'server-only'
import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type {
  LearningHistory,
  LearningHistorySummary,
  LearningHistoryLessonSummary,
} from "@/types/learning-history.models"

export async function getLearningHistory(
  page = 1,
  limit = 20
): Promise<BaseResponse<LearningHistory[]>> {
  return apiFetch<LearningHistory[]>("/learning-history", {
    query: { page, limit },
    withCredentials: true,
    tags: [CACHE_TAGS.learningHistory],
  })
}

export async function getFinishedLessons(): Promise<BaseResponse<LearningHistory[]>> {
  return apiFetch<LearningHistory[]>("/learning-history/finished", {
    withCredentials: true,
    tags: [CACHE_TAGS.learningHistory],
  })
}

export async function getUnfinishedLessons(): Promise<BaseResponse<LearningHistory[]>> {
  return apiFetch<LearningHistory[]>("/learning-history/unfinished", {
    withCredentials: true,
    tags: [CACHE_TAGS.learningHistory],
  })
}

export async function getLearningHistorySummary(): Promise<BaseResponse<LearningHistorySummary>> {
  return apiFetch<LearningHistorySummary>("/learning-history/summary", {
    withCredentials: true,
    tags: [CACHE_TAGS.learningHistory],
  })
}

export async function getLessonSummary(
  lessonId: number
): Promise<BaseResponse<LearningHistoryLessonSummary>> {
  return apiFetch<LearningHistoryLessonSummary>(
    `/learning-history/lessons/${lessonId}/summary`,
    {
      withCredentials: true,
      tags: [CACHE_TAGS.learningHistory],
    }
  )
}
