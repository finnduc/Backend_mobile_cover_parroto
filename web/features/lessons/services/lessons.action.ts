'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { redirect } from "next/navigation"
import type { BaseResponse } from "@/types/base-response"
import type { Lesson } from "@/types/lessons.models"
import { updateTag, refresh } from "next/cache"
export type CreateLessonInput = Omit<Lesson, "id" | "createdAt">
export type UpdateLessonInput = Partial<Omit<Lesson, "id" | "createdAt">>

export async function createAdminLesson(
  body: CreateLessonInput
): Promise<BaseResponse<Lesson>> {
  const res = await apiFetch<Lesson>("/admin/lessons", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    const { updateTag } = await import("next/cache")
    updateTag(CACHE_TAGS.lessons)
  }
  return res
}

export async function updateAdminLesson(
  id: number,
  body: UpdateLessonInput
): Promise<BaseResponse<Lesson>> {
  const res = await apiFetch<Lesson>(`/admin/lessons/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    const { updateTag } = await import("next/cache")
    updateTag(CACHE_TAGS.lessons)
    updateTag(CACHE_TAGS.lesson(id))
  }
  return res
}

export async function deleteAdminLesson(id: number): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/lessons/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    
    updateTag(CACHE_TAGS.lessons)
    updateTag(CACHE_TAGS.lesson(id))
    refresh()
  }
  return res
}
