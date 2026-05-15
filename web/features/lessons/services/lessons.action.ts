'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import { updateTag, refresh } from "next/cache"
import type { BaseResponse } from "@/types/base-response"
import type { Lesson } from "@/types/lessons.models"
import type { CreateLessonDto, UpdateLessonDto } from "@/features/lessons/dtos/lesson.dto"

export async function createAdminLesson(
  body: CreateLessonDto
): Promise<BaseResponse<Lesson>> {
  const res = await apiFetch<Lesson>("/admin/lessons", {
    method: "POST",
    body,
    withCredentials: true,
  })

  if (!res.error) {
    updateTag(CACHE_TAGS.lessons)
    refresh()
  }
  return res
}

export async function updateAdminLesson(
  id: number,
  body: UpdateLessonDto
): Promise<BaseResponse<Lesson>> {
  const res = await apiFetch<Lesson>(`/admin/lessons/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    updateTag(CACHE_TAGS.lessons)
    updateTag(CACHE_TAGS.lesson(id))
    refresh()
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
