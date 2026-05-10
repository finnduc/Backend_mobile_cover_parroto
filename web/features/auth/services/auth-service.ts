"use client";

import { clientApiFetch } from "@/lib/client-api";

export async function completeSignUp() {
  return clientApiFetch<{ message: string }>("/auth/complete-signup", {
    method: "POST",
  });
}
