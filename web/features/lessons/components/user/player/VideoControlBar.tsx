"use client"

import { Button } from "@/components/ui/button"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import { Pause, Play } from "lucide-react"
import { usePlayerContext } from "@/features/lessons/context/player-context"

export function VideoControlBar() {
  const { paused, autoStop, setAutoStop, onPlay, onPause } = usePlayerContext()

  return (
    <div className="flex items-center justify-between gap-4 rounded-xl border bg-card p-3">
      <div className="flex flex-col gap-3">
        <div className="flex items-center gap-2">
          <Switch
            id="auto-stop"
            size="sm"
            checked={autoStop}
            onCheckedChange={setAutoStop}
          />
          <Label htmlFor="auto-stop" className="cursor-pointer text-sm">
            Tự động dừng
          </Label>
        </div>
      </div>

      <Button
        size="lg"
        onClick={() => (paused ? onPlay() : onPause())}
        className="min-w-32"
      >
        {paused ? (
          <>
            <Play data-icon="inline-start" />
            BẮT ĐẦU
          </>
        ) : (
          <>
            <Pause data-icon="inline-start" />
            TẠM DỪNG
          </>
        )}
      </Button>
    </div>
  )
}
