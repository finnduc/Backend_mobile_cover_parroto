"use client";

import { auth } from "@/lib/firebase/client-app";
import type { BaseResponse } from "@/types/base-response";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001/api/v1";

export async function clientApiFetch<T = any>(
  url: string,
  options?: RequestInit
): Promise<BaseResponse<T>> {
  try {
    const user = auth.currentUser;
    if (!user) {
      return {
        data: null,
        error: { code: 401, message: "Not authenticated" },
      } as BaseResponse<T>;
    }

    const token = await user.getIdToken();

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...(options?.headers as Record<string, string>),
      Authorization: `Bearer ${token}`,
    };

    const response = await fetch(`${BASE_URL}${url}`, {
      ...options,
      headers,
    });

    if (!response.ok) {
      let message = "Unknown error";
      try {
        const errorData = await response.json();
        message = errorData.error?.message || errorData.message || message;
      } catch (_) {}
      return {
        data: null,
        error: { code: response.status, message },
      } as BaseResponse<T>;
    }

    const data = await response.json();
    return {
      data: data.data ?? data,
      error: null,
    } as BaseResponse<T>;
  } catch (error: any) {
    return {
      data: null,
      error: { code: error.code || 500, message: error.message || "Unknown error" },
    } as BaseResponse<T>;
  }
}
