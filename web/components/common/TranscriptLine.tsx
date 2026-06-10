import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from "@/components/ui/context-menu"
import { cn } from "@/lib/utils"
import { BookOpen, Bookmark } from "lucide-react"

function formatTimestamp(timestamp: string) {
  // Case: "1:9.944000000000003"
  if (timestamp.includes(":")) {
    const [minRaw, secRaw] = timestamp.split(":")

    const minutes = Number(minRaw)
    const seconds = Number(secRaw)

    if (!Number.isFinite(minutes) || !Number.isFinite(seconds)) {
      return timestamp
    }

    return `${minutes}:${seconds.toFixed(2).padStart(5, "0")}`
  }

  // Case: "69.944000000000003"
  const totalSeconds = Number(timestamp)

  if (!Number.isFinite(totalSeconds)) {
    return timestamp
  }

  const minutes = Math.floor(totalSeconds / 60)
  const seconds = totalSeconds % 60

  return `${minutes}:${seconds.toFixed(2).padStart(5, "0")}`
}

export function TranscriptLine({
  text,
  active = false,
  completed = false,
  bookmarked = false,
  timestamp,
  transcriptId,
  onClick,
  onAddToDeck,
  onBookmark,
}: {
  text: string
  active?: boolean
  completed?: boolean
  bookmarked?: boolean
  timestamp?: string
  transcriptId?: number
  onClick?: () => void
  onAddToDeck?: (transcriptId: number, text: string) => void
  onBookmark?: () => void
}) {

  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <button
          type="button"
          onClick={onClick}
          className={cn(
            "flex w-full gap-3 rounded-lg px-3 py-2 text-left text-sm transition-colors",
            active && "bg-transcript-active",
            completed && "bg-transcript-complete",
            !active && !completed && "hover:bg-muted/50",
          )}
        >
          {timestamp && (
            <span className="mt-0.5 shrink-0 font-mono text-xs text-muted-foreground">
              {formatTimestamp(timestamp)}
            </span>
          )}
          <span className={cn("flex-1", completed && "line-through opacity-60")}>{text}</span>
          {onBookmark && (
            <Bookmark
              className={cn(
                "size-4 shrink-0 cursor-pointer transition-colors",
                bookmarked ? "fill-current text-yellow-500" : "text-muted-foreground hover:text-foreground"
              )}
              onClick={(e) => {
                e.stopPropagation()
                onBookmark()
              }}
            />
          )}
        </button>
      </ContextMenuTrigger>
      <ContextMenuContent>
        <ContextMenuItem
          disabled={!onAddToDeck || transcriptId == null}
          onSelect={() => {
            if (transcriptId == null) return
            onAddToDeck?.(transcriptId, text)
          }}
        >
          <BookOpen className="size-4" />
          Thêm vào danh sách từ vựng
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
}
