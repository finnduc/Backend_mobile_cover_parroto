'use server'

import { apiFetch } from "@/lib/api-fetch";


export async function completeSignUp(){
  return apiFetch<void>("/auth/complete-signup", {
    method: "POST",
    withCredentials: true
  });
}