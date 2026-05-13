'use server'

import { apiFetch } from "@/lib/api-fetch";
import { BaseResponse } from "@/types/base-response";

export async function completeSignUp(token: string) {
  return apiFetch<void>("/auth/complete-signup", {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${token}`
    }
  });
}