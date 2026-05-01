"use client"

import { Button } from "@/components/ui/button"
import { SeekBar } from "@/components/common/SeekBar"
import { Play, Pause, SkipBack, SkipForward, Volume2, Settings } from "lucide-react"

export function LessonToolbar({
  playing = false,
  currentTime = 0,
  totalTime = 100,
}: {
  playing?: boolean
  currentTime?: number
  totalTime?: number
}) {
  return (
    <div className="flex items-center gap-3 rounded-xl border bg-background px-4 py-2">
      <Button variant="ghost" size="icon">
        <SkipBack className="size-4" />
      </Button>
      <Button variant="default" size="icon" className="size-8">
        {playing ? <Pause className="size-4" /> : <Play className="size-4" />}
      </Button>
      <Button variant="ghost" size="icon">
        <SkipForward className="size-4" />
      </Button>
      <SeekBar current={currentTime} total={totalTime} className="flex-1" />
      <span className="whitespace-nowrap font-mono text-xs text-muted-foreground">
        {formatTime(currentTime)} / {formatTime(totalTime)}
      </span>
      <Button variant="ghost" size="icon">
        <Volume2 className="size-4" />
      </Button>
      <Button variant="ghost" size="icon">
        <Settings className="size-4" />
      </Button>
    </div>
  )
}

function formatTime(seconds: number): string {
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`
}
