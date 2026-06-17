"use client"

import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import { usePlayerContext } from "@/features/lessons/context/player-context"

export function DictationControlBar() {
  const {
    autoStop,
    setAutoStop,
    transcriptMode,
    setTranscriptMode,
    sidebarVisible,
    setSidebarVisible,
  } = usePlayerContext()

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
        <div className="flex items-center gap-2">
          <Switch
            id="transcript-mode"
            size="sm"
            checked={transcriptMode === "full"}
            onCheckedChange={(checked) =>
              setTranscriptMode(checked ? "full" : "masked")
            }
          />
          <Label htmlFor="transcript-mode" className="cursor-pointer text-sm">
            Hiện bản chép
          </Label>
        </div>
      </div>
      <div className="flex items-center gap-2">
        <Switch
          id="sidebar-visible"
          size="sm"
          checked={sidebarVisible}
          onCheckedChange={setSidebarVisible}
        />
        <Label htmlFor="sidebar-visible" className="cursor-pointer text-sm">
          Thanh bên
        </Label>
      </div>
    </div>
  )
}
