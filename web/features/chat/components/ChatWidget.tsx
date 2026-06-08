"use client"

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { useAuth, useUser } from "@clerk/nextjs"
import { Loader2, MessageCircle, Send, X } from "lucide-react"
import {
  FormEvent,
  KeyboardEvent,
  useCallback,
  useEffect,
  useLayoutEffect,
  useRef,
  useState,
} from "react"
import { toast } from "sonner"
import { getChatHistory } from "@/features/chat/services/chat.get"
import { useChatSSE } from "@/features/chat/hooks/useChatSSE"
import type { ChatMessage } from "@/types/chat.models"
import { MessageBubble } from "./MessageBubble"

const PAGE_LIMIT = 20

function mergeAsc(existing: ChatMessage[], incoming: ChatMessage[]): ChatMessage[] {
  const seen = new Map<number, ChatMessage>()
  for (const m of existing) seen.set(m.id, m)
  for (const m of incoming) seen.set(m.id, m)
  return Array.from(seen.values()).sort((a, b) => a.id - b.id)
}

export function ChatWidget() {
  const { isSignedIn, isLoaded } = useAuth()
  const { user } = useUser()
  const [open, setOpen] = useState(false)
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [hasMore, setHasMore] = useState(false)
  const [nextId, setNextId] = useState<number | undefined>(undefined)
  const [loading, setLoading] = useState(false)
  const [loadingMore, setLoadingMore] = useState(false)
  const [sending, setSending] = useState(false)
  const [draft, setDraft] = useState("")
  const [hasLoadedOnce, setHasLoadedOnce] = useState(false)
  const [unread, setUnread] = useState(0)

  const scrollerRef = useRef<HTMLDivElement | null>(null)
  const stickToBottomRef = useRef(true)
  const restoreScrollRef = useRef<{ prevHeight: number } | null>(null)
  const openRef = useRef(open)
  const userIdRef = useRef<string | undefined>(user?.id)

  useEffect(() => {
    openRef.current = open
  }, [open])
  useEffect(() => {
    userIdRef.current = user?.id
  }, [user?.id])

  const handleIncomingMessage = useCallback((msg: ChatMessage) => {
    setMessages((prev) => mergeAsc(prev, [msg]))
    if (!openRef.current && msg.userId !== userIdRef.current) {
      setUnread((u) => u + 1)
    }
  }, [])

  const handleSocketError = useCallback((reason: string) => {
    toast.error(reason)
  }, [])

  const { status, sendMessage } = useChatSSE({
    enabled: !!isSignedIn,
    onMessage: handleIncomingMessage,
    onError: handleSocketError,
  })

  const loadInitial = useCallback(async () => {
    setLoading(true)
    const res = await getChatHistory({ limit: PAGE_LIMIT })
    setLoading(false)
    setHasLoadedOnce(true)
    if (res.error) {
      toast.error(res.error.message || "Không tải được lịch sử chat")
      return
    }
    const data = res.data
    if (!data) return
    setMessages((prev) => mergeAsc(prev, data.messages))
    setHasMore(data.hasMore)
    setNextId(data.nextId)
    stickToBottomRef.current = true
  }, [])

  const loadOlder = useCallback(async () => {
    if (!hasMore || loadingMore || nextId == null) return
    const scroller = scrollerRef.current
    if (scroller) {
      restoreScrollRef.current = { prevHeight: scroller.scrollHeight }
    }
    setLoadingMore(true)
    const res = await getChatHistory({ beforeId: nextId, limit: PAGE_LIMIT })
    setLoadingMore(false)
    if (res.error) {
      toast.error(res.error.message || "Không tải được tin cũ hơn")
      return
    }
    const data = res.data
    if (!data) return
    setMessages((prev) => mergeAsc(data.messages, prev))
    setHasMore(data.hasMore)
    setNextId(data.nextId)
  }, [hasMore, loadingMore, nextId])

  const handleToggle = useCallback(() => {
    const next = !open
    setOpen(next)
    if (next) {
      setUnread(0)
      if (!hasLoadedOnce) void loadInitial()
    }
  }, [open, hasLoadedOnce, loadInitial])

  // Auto-scroll on new messages / preserve position when prepending older messages
  useLayoutEffect(() => {
    const scroller = scrollerRef.current
    if (!scroller) return

    if (restoreScrollRef.current) {
      const delta = scroller.scrollHeight - restoreScrollRef.current.prevHeight
      scroller.scrollTop = delta
      restoreScrollRef.current = null
      return
    }

    if (stickToBottomRef.current) {
      scroller.scrollTop = scroller.scrollHeight
    }
  }, [messages, open])

  const onScroll = useCallback(() => {
    const el = scrollerRef.current
    if (!el) return
    stickToBottomRef.current = el.scrollHeight - el.scrollTop - el.clientHeight < 60
    if (el.scrollTop < 40 && hasMore && !loadingMore) {
      void loadOlder()
    }
  }, [hasMore, loadingMore, loadOlder])

  const handleSend = async (e?: FormEvent) => {
    e?.preventDefault()
    if (sending || !draft.trim()) return
    if (status !== "open") {
      toast.error("Đang mất kết nối, vui lòng thử lại")
      return
    }
    setSending(true)
    const ok = await sendMessage(draft)
    setSending(false)
    if (ok) {
      setDraft("")
      stickToBottomRef.current = true
    }
  }

  const onKeyDown = (e: KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault()
      handleSend()
    }
  }

  if (!isLoaded || !isSignedIn) return null

  return (
    <div className="pointer-events-none fixed right-4 bottom-4 z-50 flex flex-col items-end gap-3 sm:right-6 sm:bottom-6">
      {open && (
        <div className="pointer-events-auto flex h-[520px] w-[92vw] max-w-[380px] flex-col overflow-hidden rounded-2xl border bg-popover text-popover-foreground shadow-2xl">
          <header className="flex items-center justify-between border-b px-4 py-3">
            <div className="flex items-center gap-2">
              <div
                className={cn(
                  "size-2 rounded-full",
                  status === "open" ? "bg-emerald-500" : "bg-amber-500"
                )}
              />
              <div>
                <h2 className="text-sm font-semibold">Phòng chat chung</h2>
                <p className="text-[11px] text-muted-foreground">
                  {status === "open"
                    ? "Đang trực tuyến"
                    : status === "connecting"
                      ? "Đang kết nối..."
                      : "Mất kết nối"}
                </p>
              </div>
            </div>
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setOpen(false)}
              aria-label="Đóng chat"
            >
              <X className="size-4" />
            </Button>
          </header>

          <div
            ref={scrollerRef}
            onScroll={onScroll}
            className="flex-1 space-y-3 overflow-y-auto px-3 py-3"
          >
            {loadingMore && (
              <div className="flex items-center justify-center py-1">
                <Loader2 className="size-4 animate-spin text-muted-foreground" />
              </div>
            )}
            {loading && messages.length === 0 ? (
              <div className="flex h-full items-center justify-center">
                <Loader2 className="size-5 animate-spin text-muted-foreground" />
              </div>
            ) : messages.length === 0 ? (
              <div className="flex h-full items-center justify-center px-6 text-center text-sm text-muted-foreground">
                Chưa có tin nhắn nào. Hãy là người đầu tiên chào mọi người!
              </div>
            ) : (
              messages.map((m) => (
                <MessageBubble key={m.id} message={m} isOwn={m.userId === user?.id} />
              ))
            )}
          </div>

          <form
            onSubmit={handleSend}
            className="flex items-end gap-2 border-t bg-background px-3 py-2"
          >
            <textarea
              value={draft}
              onChange={(e) => setDraft(e.target.value)}
              onKeyDown={onKeyDown}
              placeholder="Nhập tin nhắn..."
              rows={1}
              maxLength={1000}
              className="max-h-28 min-h-9 flex-1 resize-none rounded-lg border border-input bg-transparent px-3 py-2 text-sm outline-none transition-colors placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50 dark:bg-input/30"
            />
            <Button
              type="submit"
              size="icon"
              disabled={sending || !draft.trim() || status !== "open"}
              aria-label="Gửi"
            >
              <Send className="size-4" />
            </Button>
          </form>
        </div>
      )}

      <Button
        type="button"
        size="icon-lg"
        className={cn(
          "pointer-events-auto relative size-12 rounded-full shadow-lg transition-transform",
          open && "rotate-90"
        )}
        onClick={handleToggle}
        aria-label={open ? "Đóng chat" : "Mở chat"}
      >
        {open ? <X className="size-5" /> : <MessageCircle className="size-5" />}
        {!open && unread > 0 && (
          <span className="absolute -top-1 -right-1 flex h-5 min-w-5 items-center justify-center rounded-full bg-destructive px-1.5 text-[10px] font-semibold text-destructive-foreground">
            {unread > 99 ? "99+" : unread}
          </span>
        )}
      </Button>
    </div>
  )
}
