'use server'

import { apiFetch } from "@/lib/api-fetch"
import { CACHE_TAGS } from "@/lib/tags"
import type { BaseResponse } from "@/types/base-response"
import type { Transcript } from "@/types/lessons.models"
import { updateTag, refresh } from "next/cache"

export type CreateTranscriptInput = Omit<Transcript, "id">
export type UpdateTranscriptInput = Partial<Omit<Transcript, "id">>

export async function createAdminTranscript(
  body: CreateTranscriptInput
): Promise<BaseResponse<Transcript>> {
  const res = await apiFetch<Transcript>("/admin/transcripts", {
    method: "POST",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    
    updateTag(CACHE_TAGS.transcripts)
    updateTag(CACHE_TAGS.lesson(body.lessonId))
    refresh()
  }
  return res
}

export async function updateAdminTranscript(
  id: number,
  body: UpdateTranscriptInput
): Promise<BaseResponse<Transcript>> {
  const res = await apiFetch<Transcript>(`/admin/transcripts/${id}`, {
    method: "PUT",
    body,
    withCredentials: true,
  })
  if (!res.error) {
    
    updateTag(CACHE_TAGS.transcripts)
    if (body.lessonId) {
      updateTag(CACHE_TAGS.lesson(body.lessonId))
    }
    refresh()
  }
  return res
}

export async function deleteAdminTranscript(
  id: number,
  lessonId?: number
): Promise<BaseResponse<void>> {
  const res = await apiFetch<void>(`/admin/transcripts/${id}`, {
    method: "DELETE",
    withCredentials: true,
  })
  if (!res.error) {
    
    updateTag(CACHE_TAGS.transcripts)
    if (lessonId) {
      updateTag(CACHE_TAGS.lesson(lessonId))
    }
    refresh()
  }
  return res
}
