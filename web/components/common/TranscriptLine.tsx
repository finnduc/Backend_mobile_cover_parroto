import { cn } from "@/lib/utils"

export function TranscriptLine({
  text,
  active = false,
  completed = false,
  timestamp,
  onClick,
}: {
  text: string
  active?: boolean
  completed?: boolean
  timestamp?: string
  onClick?: () => void
}) {
  return (
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
  )
}
