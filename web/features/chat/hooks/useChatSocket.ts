"use client"

import { useAuth } from "@clerk/nextjs"
import { useCallback, useEffect, useRef, useState } from "react"
import { snakeToCamel } from "@/lib/case"
import type { ChatMessage } from "@/types/chat.models"

type ConnectionStatus = "idle" | "connecting" | "open" | "closed" | "error"

interface WSEventRaw {
  type: string
  data?: Record<string, unknown>
}

interface UseChatSocketOptions {
  enabled: boolean
  onMessage?: (msg: ChatMessage) => void
  onError?: (reason: string) => void
}

function resolveWsUrl(): string | null {
  const httpBase = process.env.NEXT_PUBLIC_API_URL
  if (!httpBase) return null
  try {
    const u = new URL(httpBase)
    u.protocol = u.protocol === "https:" ? "wss:" : "ws:"
    const path = u.pathname.replace(/\/$/, "")
    return `${u.origin}${path}/chat/ws`
  } catch {
    return null
  }
}

export function useChatSocket({ enabled, onMessage, onError }: UseChatSocketOptions) {
  const { isSignedIn, getToken } = useAuth()
  const wsRef = useRef<WebSocket | null>(null)
  const reconnectTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null)
  const reconnectAttemptRef = useRef(0)
  const enabledRef = useRef(enabled)
  const onMessageRef = useRef(onMessage)
  const onErrorRef = useRef(onError)
  const connectRef = useRef<() => Promise<void>>(async () => {})

  const [status, setStatus] = useState<ConnectionStatus>("idle")

  useEffect(() => {
    enabledRef.current = enabled
  }, [enabled])
  useEffect(() => {
    onMessageRef.current = onMessage
  }, [onMessage])
  useEffect(() => {
    onErrorRef.current = onError
  }, [onError])

  const connect = useCallback(async () => {
    if (!enabledRef.current || !isSignedIn) return
    const baseWsUrl = resolveWsUrl()
    if (!baseWsUrl) {
      onErrorRef.current?.("NEXT_PUBLIC_API_URL is not set")
      setStatus("error")
      return
    }

    const token = await getToken()
    if (!token) {
      onErrorRef.current?.("Missing auth token")
      setStatus("error")
      return
    }

    const url = `${baseWsUrl}?token=${encodeURIComponent(token)}`
    setStatus("connecting")
    const ws = new WebSocket(url)
    wsRef.current = ws

    ws.onopen = () => {
      reconnectAttemptRef.current = 0
      setStatus("open")
    }

    ws.onmessage = (ev) => {
      try {
        const raw = JSON.parse(ev.data) as WSEventRaw
        if (raw.type === "message" && raw.data) {
          onMessageRef.current?.(snakeToCamel<ChatMessage>(raw.data))
        } else if (raw.type === "error" && typeof raw.data?.message === "string") {
          onErrorRef.current?.(raw.data.message)
        }
      } catch {
        // ignore malformed frames
      }
    }

    ws.onerror = () => {
      setStatus("error")
    }

    ws.onclose = () => {
      wsRef.current = null
      setStatus("closed")
      if (!enabledRef.current) return
      const attempt = ++reconnectAttemptRef.current
      const delay = Math.min(30_000, 1000 * 2 ** Math.min(attempt, 5))
      reconnectTimerRef.current = setTimeout(() => {
        void connectRef.current()
      }, delay)
    }
  }, [isSignedIn, getToken])

  useEffect(() => {
    connectRef.current = connect
  }, [connect])

  useEffect(() => {
    if (!enabled || !isSignedIn) return
    const handle = setTimeout(() => {
      void connectRef.current()
    }, 0)
    return () => {
      clearTimeout(handle)
      if (reconnectTimerRef.current) {
        clearTimeout(reconnectTimerRef.current)
        reconnectTimerRef.current = null
      }
      const ws = wsRef.current
      wsRef.current = null
      if (ws && ws.readyState === WebSocket.OPEN) {
        ws.close(1000, "client unmount")
      } else if (ws) {
        ws.onopen = null
        ws.onmessage = null
        ws.onerror = null
        ws.onclose = null
      }
      setStatus("idle")
    }
  }, [enabled, isSignedIn])

  const sendMessage = useCallback((content: string): boolean => {
    const ws = wsRef.current
    if (!ws || ws.readyState !== WebSocket.OPEN) return false
    const trimmed = content.trim()
    if (!trimmed) return false
    ws.send(JSON.stringify({ content: trimmed }))
    return true
  }, [])

  return { status, sendMessage }
}
