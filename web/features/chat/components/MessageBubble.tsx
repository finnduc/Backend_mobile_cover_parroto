"use client"

import { cn } from "@/lib/utils"
import type { ChatMessage } from "@/types/chat.models"

function formatTime(iso: string): string {
  try {
    const d = new Date(iso)
    return d.toLocaleTimeString("vi-VN", { hour: "2-digit", minute: "2-digit" })
  } catch {
    return ""
  }
}

function fallbackInitial(name?: string, userId?: string): string {
  const source = name?.trim() || userId || "?"
  return source.charAt(0).toUpperCase()
}

export function MessageBubble({
  message,
  isOwn,
}: {
  message: ChatMessage
  isOwn: boolean
}) {
  return (
    <div
      className={cn(
        "flex w-full items-end gap-2",
        isOwn ? "justify-end" : "justify-start"
      )}
    >
      {!isOwn && (
        <div className="flex size-7 shrink-0 items-center justify-center overflow-hidden rounded-full bg-muted text-xs font-medium text-muted-foreground">
          {message.avatarUrl ? (
            // eslint-disable-next-line @next/next/no-img-element
            <img
              src={message.avatarUrl}
              alt={message.userName ?? message.userId}
              className="size-full object-cover"
            />
          ) : (
            fallbackInitial(message.userName, message.userId)
          )}
        </div>
      )}
      <div className={cn("flex max-w-[70%] flex-col gap-0.5", isOwn ? "items-end" : "items-start")}>
        {!isOwn && (
          <span className="px-1 text-[11px] text-muted-foreground">
            {message.userName?.trim() || "Người dùng"}
          </span>
        )}
        <div
          className={cn(
            "rounded-2xl px-3 py-1.5 text-sm leading-relaxed whitespace-pre-wrap break-words",
            isOwn
              ? "rounded-br-sm bg-primary text-primary-foreground"
              : "rounded-bl-sm bg-muted text-foreground"
          )}
        >
          {message.content}
        </div>
        <span className="px-1 text-[10px] text-muted-foreground">
          {formatTime(message.createdAt)}
        </span>
      </div>
    </div>
  )
}
