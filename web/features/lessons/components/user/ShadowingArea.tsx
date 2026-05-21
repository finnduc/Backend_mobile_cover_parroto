"use client"

import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import { Button } from "@/components/ui/button"
import { Mic, Square, Play, Pause, SkipBack, SkipForward, RotateCcw, Check } from "lucide-react"
import { useState, useTransition } from "react"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { toast } from "sonner"
import type { Transcript } from "@/types/lessons.models"

export interface ExerciseControlProps {
  transcripts: Transcript[]
  activeIndex: number
  highlightedIndex: number
  paused: boolean
  initialCompletedIds: number[]
  lessonId: number
  onPlay: () => void
  onPause: () => void
  onNext: () => void
  onPrev: () => void
  onReplay: () => void
  onTranscriptClick: (index: number) => void
}

export function ShadowingArea({
  transcripts,
  activeIndex,
  paused,
  initialCompletedIds,
  lessonId,
  onPlay,
  onPause,
  onNext,
  onPrev,
  onReplay,
}: Partial<ExerciseControlProps>) {
  const [recording, setRecording] = useState(false)
  const [completed, setCompleted] = useState<number[]>(initialCompletedIds ?? [])
  const [, startTransition] = useTransition()

  const lines = (transcripts ?? []).map((t) => t.content)
  const safeActiveIndex = (activeIndex ?? -1) >= 0 && (activeIndex ?? -1) < lines.length ? activeIndex! : 0

  return (
    <div className="space-y-4">
      <LessonProgressBar completed={completed.length} total={lines.length} />
      <div className="rounded-xl bg-gradient-to-b from-muted/50 to-background p-6 text-center">
        <div className="mb-4 flex items-center justify-center gap-2">
          <div className="flex h-16 w-64 items-center justify-center gap-1 rounded-full bg-muted px-4">
            {Array.from({ length: 40 }).map((_, i) => (
              <div
                key={i}
                className="w-0.5 rounded-full bg-primary"
                style={{
                  height: `${Math.random() * 40 + 8}px`,
                  opacity: recording ? 1 : 0.3,
                }}
              />
            ))}
          </div>
        </div>
        <p className="mb-6 text-lg font-medium">{lines[safeActiveIndex]}</p>
        <div className="flex items-center justify-center gap-3">
          <Button
            size="lg"
            variant={recording ? "destructive" : "default"}
            className="size-14 rounded-full"
            onClick={() => setRecording(!recording)}
          >
            {recording ? <Square className="size-6" /> : <Mic className="size-6" />}
          </Button>
          <Button
            size="icon"
            variant="outline"
            onClick={onPrev ?? (() => {})}
            disabled={(activeIndex ?? -1) <= 0}
          >
            <SkipBack className="size-4" />
          </Button>
          <Button
            size="icon"
            variant="outline"
            onClick={() => {
              if (paused) {
                onPlay?.()
              } else {
                onPause?.()
              }
            }}
          >
            {paused ? <Play className="size-4" /> : <Pause className="size-4" />}
          </Button>
          <Button
            size="icon"
            variant="outline"
            onClick={() => {
              if (!completed.includes(safeActiveIndex)) {
                setCompleted([...completed, safeActiveIndex])
                const seg = (transcripts ?? [])[safeActiveIndex]
                if (seg) {
                  startTransition(async () => {
                    const res = await postShadowingStatus(seg.id, lessonId!)
                    if (res.error) {
                      toast.error(res.error.message)
                    }
                  })
                }
              }
              onNext?.()
            }}
            disabled={(activeIndex ?? -1) >= lines.length - 1}
          >
            <SkipForward className="size-4" />
          </Button>
          <Button size="icon" variant="ghost" onClick={onReplay ?? (() => {})}>
            <RotateCcw className="size-4" />
          </Button>
        </div>
      </div>
      <div className="space-y-1">
        {(transcripts ?? []).map((seg, i) => {
          const isCompleted = completed.includes(i) && i !== safeActiveIndex
          const isSaved = (initialCompletedIds ?? []).includes(i)
          return (
            <div
              key={seg.id}
              className={`flex items-center gap-2 rounded px-3 py-1.5 text-sm ${
                isSaved ? "bg-green-100 text-green-700 line-through" : ""
              } ${isCompleted && !isSaved ? "bg-amber-50 text-amber-700" : ""} ${i === safeActiveIndex ? "bg-transcript-active font-medium" : ""}`}
            >
              {isSaved && <Check className="size-3.5 shrink-0 text-green-600" />}
              {isCompleted && !isSaved && <span className="size-3.5 shrink-0 rounded-full border border-amber-400" />}
              {seg.content}
            </div>
          )
        })}
      </div>
    </div>
  )
}
