"use client"

import { Button } from "@/components/ui/button"
import { ScrollText, Play } from "lucide-react"

export function ModeToggle({
  mode,
  onChange,
}: {
  mode: "normal" | "transcript"
  onChange: (mode: "normal" | "transcript") => void
}) {
  return (
    <div className="flex items-center gap-2 rounded-xl border bg-background px-3 py-2">
      <Button
        variant={mode === "normal" ? "default" : "ghost"}
        size="sm"
        onClick={() => onChange("normal")}
        className="gap-1.5"
      >
        <Play className="size-3.5" />
        Normal
      </Button>
      <Button
        variant={mode === "transcript" ? "default" : "ghost"}
        size="sm"
        onClick={() => onChange("transcript")}
        className="gap-1.5"
      >
        <ScrollText className="size-3.5" />
        Transcript
      </Button>
    </div>
  )
}
