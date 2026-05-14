'use server'
import { apiFetch } from "@/lib/api-fetch";

export async function completeSignUp(values: any) {
  return apiFetch<void>("/auth/complete-signup", {
    method: "POST",
    withCredentials: true,
    body: JSON.stringify(values)
  });
}