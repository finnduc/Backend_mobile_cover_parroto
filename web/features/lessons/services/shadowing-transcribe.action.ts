'use server'

import { apiFetch } from "@/lib/api-fetch"
import type { BaseResponse } from "@/types/base-response"
import { ShadowingTranscribeRes } from "../dtos/shadowing-transcribe.dto"

export async function postShadowingTranscribe(
  audioBlob: Blob
): Promise<BaseResponse<ShadowingTranscribeRes>> {
  const formData = new FormData()
  formData.append("audio", audioBlob, "recording.webm")

  const res = await apiFetch<ShadowingTranscribeRes>("/shadowing-status/transcribe", {
    method: "POST",
    body: formData,
    withCredentials: true,
  })

  return res
}
