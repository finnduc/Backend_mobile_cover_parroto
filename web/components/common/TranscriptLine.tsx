import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from "@/components/ui/context-menu"
import { cn } from "@/lib/utils"
import { BookOpen } from "lucide-react"

export function TranscriptLine({
  text,
  active = false,
  completed = false,
  timestamp,
  transcriptId,
  onClick,
  onAddToDeck,
}: {
  text: string
  active?: boolean
  completed?: boolean
  timestamp?: string
  transcriptId?: number
  onClick?: () => void
  onAddToDeck?: (transcriptId: number, text: string) => void
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
            <span className="mt-0.5 shrink-0 font-mono text-xs text-muted-foreground">{timestamp}</span>
          )}
          <span className={cn(completed && "line-through opacity-60")}>{text}</span>
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
