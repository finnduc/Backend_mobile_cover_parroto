import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from "@/components/ui/context-menu"
import { cn } from "@/lib/utils"
import { BookOpen, Eye, EyeOff } from "lucide-react"

function formatTimestamp(timestamp: string) {
  if (timestamp.includes(":")) {
    const [minRaw, secRaw] = timestamp.split(":")
    const minutes = Number(minRaw)
    const seconds = Number(secRaw)
    if (!Number.isFinite(minutes) || !Number.isFinite(seconds)) {
      return timestamp
    }
    return `${minutes}:${seconds.toFixed(2).padStart(5, "0")}`
  }
  const totalSeconds = Number(timestamp)
  if (!Number.isFinite(totalSeconds)) {
    return timestamp
  }
  const minutes = Math.floor(totalSeconds / 60)
  const seconds = totalSeconds % 60
  return `${minutes}:${seconds.toFixed(2).padStart(5, "0")}`
}

function MaskedWord({ word }: { word: string }) {
  const width = Math.max(word.length * 7, 24)
  return (
    <span
      aria-hidden
      className="inline-block h-3 rounded bg-muted align-middle"
      style={{ width: `${width}px` }}
    />
  )
}

export function TranscriptLine({
  text,
  phonetic,
  active = false,
  completed = false,
  timestamp,
  transcriptId,
  onClick,
  onAddToDeck,
  masked = false,
  revealed = false,
  onReveal,
}: {
  text: string
  phonetic?: string
  active?: boolean
  completed?: boolean
  timestamp?: string
  transcriptId?: number
  onClick?: () => void
  onAddToDeck?: (transcriptId: number, text: string) => void
  masked?: boolean
  revealed?: boolean
  onReveal?: () => void
}) {
  const words = masked ? text.split(/\s+/) : []
  return (
    <ContextMenu>
      <ContextMenuTrigger asChild>
        <button
          type="button"
          onClick={onClick}
          className={cn(
            "flex w-full items-center gap-3 rounded-lg px-3 py-2 text-left text-sm transition-colors",
            active && "border-2 border-primary bg-transcript-active",
            !active && completed && "bg-transcript-complete",
            !active && !completed && "hover:bg-muted/50",
          )}
        >
          {timestamp && (
            <span className="mt-0.5 shrink-0 font-mono text-xs text-muted-foreground">
              {formatTimestamp(timestamp)}
            </span>
          )}
          <div className="flex min-w-0 flex-1 flex-wrap items-center gap-1">
            {masked && !revealed ? (
              words.map((w, i) => <MaskedWord key={i} word={w} />)
            ) : (
              <span className={cn("block", completed && "line-through opacity-60")}>
                {text}
              </span>
            )}
          </div>
          {masked && onReveal && (
            <span
              role="button"
              tabIndex={0}
              onClick={(e) => {
                e.stopPropagation()
                onReveal()
              }}
              onKeyDown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                  e.preventDefault()
                  e.stopPropagation()
                  onReveal()
                }
              }}
              className="shrink-0 rounded p-1 text-muted-foreground hover:bg-muted hover:text-foreground"
            >
              {revealed ? <EyeOff className="size-4" /> : <Eye className="size-4" />}
            </span>
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
          <BookOpen data-icon="inline-start" />
          Thêm vào danh sách từ vựng
        </ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  )
}
