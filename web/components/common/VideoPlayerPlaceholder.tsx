import { Play } from "lucide-react"

function formatDuration(seconds: number): string {
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`
}

export function VideoPlayerPlaceholder({ duration }: { duration?: number }) {
  return (
    <div className="relative aspect-video w-full overflow-hidden rounded-xl bg-muted">
      <div className="flex size-full items-center justify-center bg-gradient-to-br from-muted-foreground/10 to-muted-foreground/20">
        <div className="flex size-16 items-center justify-center rounded-full bg-primary/90 shadow-lg">
          <Play className="ml-0.5 size-8 fill-primary-foreground text-primary-foreground" />
        </div>
      </div>
      {duration !== undefined && (
        <div className="absolute bottom-2 right-2 rounded bg-black/70 px-1.5 py-0.5 text-xs text-white">
          {formatDuration(duration)}
        </div>
      )}
    </div>
  )
}
