'use server'

import { apiFetch } from "@/lib/api-fetch"
import type { BaseResponse } from "@/types/base-response"
import type { ChatHistory } from "@/types/chat.models"

export async function getChatHistory(params?: {
  beforeId?: number
  limit?: number
}): Promise<BaseResponse<ChatHistory>> {
  return apiFetch<ChatHistory>("/chat/messages", {
    withCredentials: true,
    query: {
      before_id: params?.beforeId,
      limit: params?.limit ?? 20,
    },
    cache: "no-store",
  })
}
