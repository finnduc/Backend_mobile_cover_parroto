"use client"

import { useAuth } from "@clerk/nextjs"
import { useCallback, useEffect, useRef, useState } from "react"
import { snakeToCamel } from "@/lib/case"
import type { ChatMessage } from "@/types/chat.models"

type ConnectionStatus = "idle" | "connecting" | "open" | "closed" | "error"

interface UseChatSSEOptions {
  enabled: boolean
  onMessage?: (msg: ChatMessage) => void
  onError?: (reason: string) => void
}

export function useChatSSE({ enabled, onMessage, onError }: UseChatSSEOptions) {
  const { isSignedIn, getToken } = useAuth()
  const eventSourceRef = useRef<EventSource | null>(null)
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

    const token = await getToken()
    if (!token) {
      onErrorRef.current?.("Missing auth token")
      setStatus("error")
      return
    }

    setStatus("connecting")

    // Close existing connection
    if (eventSourceRef.current) {
      eventSourceRef.current.close()
      eventSourceRef.current = null
    }

    // Create new EventSource with token in URL
    // r3labs/sse requires ?stream= parameter to specify which stream to subscribe to
    const url = `/api/chat/events?stream=messages&token=${encodeURIComponent(token)}`
    const eventSource = new EventSource(url)
    eventSourceRef.current = eventSource

    eventSource.onopen = () => {
      reconnectAttemptRef.current = 0
      setStatus("open")
    }

    eventSource.addEventListener("chat.message.created", (event) => {
      try {
        const raw = JSON.parse(event.data)
        const message = snakeToCamel<ChatMessage>(raw)
        onMessageRef.current?.(message)
      } catch {
        // ignore malformed events
      }
    })

    eventSource.addEventListener("ping", () => {
      // keepalive only, ignore
      console.log("ping success")
    })

    eventSource.onerror = () => {
      setStatus("error")
      onErrorRef.current?.("SSE connection error")
      // Close to prevent native auto-reconnect with stale token
      eventSource.close()
      eventSourceRef.current = null
      // Schedule manual reconnect with fresh token
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
      if (eventSourceRef.current) {
        eventSourceRef.current.close()
        eventSourceRef.current = null
      }
      setStatus("idle")
    }
  }, [enabled, isSignedIn])

  const sendMessage = useCallback(
    async (content: string): Promise<boolean> => {
      const trimmed = content.trim()
      if (!trimmed) return false

      try {
        const token = await getToken()
        if (!token) return false

        const response = await fetch("/api/chat/messages", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify({ content: trimmed }),
        })

        return response.ok
      } catch {
        return false
      }
    },
    [getToken]
  )

  return { status, sendMessage }
}
