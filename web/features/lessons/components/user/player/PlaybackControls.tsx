"use client"

import { Button } from "@/components/ui/button"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { Pause, Play, RotateCcw, SkipBack, SkipForward } from "lucide-react"
import type { ReactNode } from "react"

export function PlaybackControls({
  children,
  totalLines,
}: {
  children?: ReactNode
  totalLines: number
}) {
  const { paused, activeIndex, highlightedIndex, onPlay, onPause, onNext, onPrev, onReplay } = usePlayerContext()

  const idx = activeIndex >= 0 ? activeIndex : highlightedIndex

  return (
    <div className="flex items-center justify-center gap-2 rounded-xl border bg-background px-3 py-2">
      {children}
      <Button
        size="icon"
        variant="outline"
        onClick={onPrev}
        disabled={idx <= 0}
      >
        <SkipBack className="size-4" />
      </Button>
      <Button
        size="icon"
        variant="outline"
        onClick={() => (paused ? onPlay() : onPause())}
      >
        {paused ? <Play className="size-4" /> : <Pause className="size-4" />}
      </Button>
      <Button
        size="icon"
        variant="outline"
        onClick={onNext}
        disabled={idx >= totalLines - 1}
      >
        <SkipForward className="size-4" />
      </Button>
      <Button size="icon" variant="ghost" onClick={onReplay}>
        <RotateCcw className="size-4" />
      </Button>
    </div>
  )
}
