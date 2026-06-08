import { NextRequest } from "next/server"
import { auth } from "@clerk/nextjs/server"

const BACKEND_URL = process.env.API_URL

async function proxyRequest(request: NextRequest, method: string) {
  if (!BACKEND_URL) {
    return new Response("API_URL not configured", { status: 500 })
  }

  const { getToken } = await auth()
  const token = await getToken()

  if (!token) {
    return new Response("Unauthorized", { status: 401 })
  }

  // Forward query parameters
  const searchParams = request.nextUrl.searchParams.toString()
  const backendUrl = `${BACKEND_URL}/chat/messages${searchParams ? `?${searchParams}` : ""}`

  // Prepare headers
  const headers: Record<string, string> = {
    Authorization: `Bearer ${token}`,
  }
  if (process.env.API_KEY) {
    headers.apikey = process.env.API_KEY
  }

  // Prepare body for POST
  let body: string | undefined
  if (method === "POST") {
    headers["Content-Type"] = "application/json"
    body = await request.text()
  }

  // Forward request to backend
  const controller = new AbortController()
  const timeout = setTimeout(() => controller.abort(), 15000)

  try {
    const backendResponse = await fetch(backendUrl, {
      method,
      headers,
      body,
      signal: controller.signal,
    })

    // Return backend response
    const responseData = await backendResponse.text()
    return new Response(responseData, {
      status: backendResponse.status,
      headers: {
        "Content-Type": backendResponse.headers.get("Content-Type") || "application/json",
      },
    })
  } catch (error) {
    const message = error instanceof Error ? error.message : "Unknown error"
    return new Response(`Backend request failed: ${message}`, { status: 502 })
  } finally {
    clearTimeout(timeout)
  }
}

export async function GET(request: NextRequest) {
  return proxyRequest(request, "GET")
}

export async function POST(request: NextRequest) {
  return proxyRequest(request, "POST")
}
