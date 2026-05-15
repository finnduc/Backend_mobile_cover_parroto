'server-only'
import type { BaseResponse } from "@/types/base-response";
import { snakeToCamel, camelToSnake } from "./case";
import { auth } from "@clerk/nextjs/server";

type ApiFetchOptions = {
  baseUrl?: string;
  withCredentials?: boolean;
  query?: Record<string, any>;
  body?: Record<string, any> | BodyInit | null;
  tags?: string[];
} & Omit<RequestInit, "body">;

export async function apiFetch<T = any>(
  url: string,
  options?: ApiFetchOptions
): Promise<BaseResponse<T>> {

  try {
    const {
      withCredentials = false,
      baseUrl = process.env.API_URL,
      query,
      tags,
      ...fetchOptions
    } = options || {};

    if (!baseUrl) {
      throw new Error(
        "Server API_URL is not configured. Please set API_URL environment variable."
      );
    }

    const headers: Record<string, any> = {
      ...fetchOptions?.headers,
      apikey: process.env.API_KEY || "",
    };

    if (withCredentials) {
      const { getToken } = await auth();
      const token = await getToken();
      if (token) headers["Authorization"] = `Bearer ${token}`;
    }

    const searchParams = new URLSearchParams();
    if (query) {
      Object.entries(query).forEach(([key, value]) => {
        if (value === undefined || value === null || value === "") return;

        if (typeof value === "object") {
          // If it's a DynamoDB object, we MUST stringify it properly
          searchParams.append(key, JSON.stringify(value));
        } else {
          searchParams.append(key, String(value));
        }
      });
    }

    let body = fetchOptions.body;
    if (body && typeof body === "object" && !(body instanceof FormData) && !(body instanceof Blob) && !(body instanceof URLSearchParams) && !(body instanceof ArrayBuffer) && !(body instanceof ReadableStream)) {
      headers["Content-Type"] = "application/json";
      body = JSON.stringify(camelToSnake(body));
    }

    const queryString = searchParams.toString();

    const fullUrl = `${baseUrl}${url}${queryString ? `?${queryString}` : ""}`;

    console.log("🚀 Calling API:", fullUrl);

    const response = await fetch(fullUrl, {
      method: fetchOptions.method || "GET",
      ...fetchOptions,
      headers,
      body,
      ...(tags && { next: { tags } }),
    });

    if (!response.ok) {
      let message = "Unknown error";
      try {
        const errorData = await response.json();
        message = errorData.error?.message || errorData.message || message;
      } catch (_) { }
      throw { code: response.status, message };
    }

    let rawData: any;
    const contentType = response.headers.get("content-type");

    if (contentType && contentType.includes("application/json")) {
      const text = await response.text();
      rawData = text ? JSON.parse(text) : {};
    } else {
      rawData = {};
    }

    const data = snakeToCamel<any>(rawData);
    const resultData = data.data !== undefined ? data.data : data;
    const resultMeta = data.pagination || data.meta
      ? { pagination: data.pagination, ...data.meta }
      : undefined;

    return {
      data: resultData as T,
      error: null,
      ...(resultMeta && { meta: resultMeta }),
    } as BaseResponse<T>;
  } catch (error: any) {
    console.error(error);
    return {
      data: null,
      error: {
        code: error.code || 500,
        message: error.message || "Unknown error",
      },
    } as BaseResponse<T>;
  }
}