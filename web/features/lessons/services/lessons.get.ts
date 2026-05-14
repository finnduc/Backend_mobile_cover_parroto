'server-only'
import { apiFetch } from "@/lib/api-fetch"
import type { Lesson, Transcript } from "@/types/lessons.models"
import type { BaseResponse } from "@/types/base-response"

export async function getLessons(
  page = 1,
  limit = 10,
  filters?: { categoryId?: number; level?: string }
): Promise<BaseResponse<Lesson[]>> {
  const query: Record<string, any> = { page, limit }
  if (filters?.categoryId) query.category_id = filters.categoryId
  if (filters?.level) query.level = filters.level
  return apiFetch<Lesson[]>("/lessons", { query })
}

export async function getLesson(id: number): Promise<BaseResponse<Lesson>> {
  return apiFetch<Lesson>(`/lessons/${id}`)
}

export async function getTranscripts(lessonId: number): Promise<BaseResponse<Transcript[]>> {
  return apiFetch<Transcript[]>(`/lessons/${lessonId}/transcripts`)
}

export async function getAdminLessons(
  page = 1,
  limit = 10,
  filters?: { categoryId?: number; level?: string }
): Promise<BaseResponse<Lesson[]>> {
  const query: Record<string, any> = { page, limit }
  if (filters?.categoryId) query.category_id = filters.categoryId
  if (filters?.level) query.level = filters.level
  return apiFetch<Lesson[]>("/admin/lessons", { query, withCredentials: true })
}

export async function getAdminLesson(id: number): Promise<BaseResponse<Lesson>> {
  return apiFetch<Lesson>(`/admin/lessons/${id}`, { withCredentials: true })
}

export async function getAdminTranscripts(lessonId: number): Promise<BaseResponse<Transcript[]>> {
  return apiFetch<Transcript[]>(`/admin/lessons/${lessonId}/transcripts`, { withCredentials: true })
}
