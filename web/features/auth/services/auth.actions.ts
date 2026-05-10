'use server'

import { apiFetch } from "@/lib/api-fetch";
import { BaseResponse } from "@/types/base-response";


export async function completeSignUp(){
  return apiFetch<BaseResponse<void>>("/auth/complete-signup", {
    method: "POST",
  });
}