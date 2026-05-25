"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import type { ExerciseControlProps } from "@/features/lessons/components/user/ShadowingArea"
import { Lightbulb, Mic, Pause, Play, RotateCcw, SkipBack, SkipForward, Square, Check } from "lucide-react"
import { useState, useEffect, useTransition } from "react"
import { postDictationStatus } from "@/features/lessons/services/dictation-status.action"
import { toast } from "sonner"

export function DictationArea({
  transcripts,
  activeIndex,
  highlightedIndex,
  paused,
  initialCompletedIds,
  lessonId,
  onPlay,
  onPause,
  onNext,
  onPrev,
  onReplay,
  onTranscriptClick,
}: Partial<ExerciseControlProps>) {
  const [recording, setRecording] = useState(false)
  const initialMaxLine = (initialCompletedIds ?? []).length > 0
    ? Math.max(...(initialCompletedIds ?? [])) + 1
    : 0
  const [maxLine, setMaxLine] = useState(initialMaxLine)
  const [inputs, setInputs] = useState<string[]>((transcripts ?? []).map(() => ""))
  const [showHints, setShowHints] = useState(false)
  const [, startTransition] = useTransition()

  // Sync current line with active playback position
  const currentLine =
    (activeIndex ?? -1) >= 0
      ? activeIndex!
      : (highlightedIndex ?? -1) >= 0
        ? highlightedIndex!
        : 0

  // Update max progress when activeIndex advances explicitly
  useEffect(() => {
    if ((activeIndex ?? -1) >= 0) {
      setMaxLine((prev) => Math.max(prev, activeIndex! + 1))
    }
  }, [activeIndex])

  return (
    <div className="space-y-4">
      <LessonProgressBar completed={(initialCompletedIds ?? []).length} total={(transcripts ?? []).length} />

      <div className="flex items-center justify-center gap-2 rounded-xl border bg-background px-3 py-2">
        <Button
          size="lg"
          variant={recording ? "destructive" : "default"}
          className="size-12 rounded-full"
          onClick={() => setRecording(!recording)}
        >
          {recording ? <Square className="size-5" /> : <Mic className="size-5" />}
        </Button>
        <Button size="icon" variant="outline" onClick={onPrev ?? (() => {})} disabled={(activeIndex ?? -1) <= 0 && (highlightedIndex ?? -1) <= 0}>
          <SkipBack className="size-4" />
        </Button>
        <Button
          size="icon"
          variant="outline"
          onClick={() => (paused ? onPlay?.() : onPause?.())}
        >
          {paused ? <Play className="size-4" /> : <Pause className="size-4" />}
        </Button>
        <Button
          size="icon"
          variant="outline"
          onClick={onNext ?? (() => {})}
          disabled={
            ((activeIndex ?? -1) >= 0 ? activeIndex! : highlightedIndex ?? 0) >=
            (transcripts ?? []).length - 1
          }
        >
          <SkipForward className="size-4" />
        </Button>
        <Button size="icon" variant="ghost" onClick={onReplay ?? (() => {})}>
          <RotateCcw className="size-4" />
        </Button>
      </div>

      <div className="space-y-2">
        {(transcripts ?? []).map((seg, i) => {
          const isComplete = i < maxLine
          const isCurrent = i === currentLine
          const isSaved = (initialCompletedIds ?? []).includes(i)
          return (
            <div
              key={seg.id}
              onClick={() => {
                if (isComplete && !isCurrent) {
                  onTranscriptClick?.(i)
                }
              }}
              className={`rounded-lg p-3 transition-colors ${
                isCurrent ? "border-2 border-primary/20 bg-transcript-active" : ""
              } ${isSaved ? "bg-green-100" : ""} ${isComplete && !isSaved && !isCurrent ? "bg-amber-50" : ""} ${
                isComplete && !isCurrent ? "cursor-pointer hover:bg-transcript-active/50" : ""
              }`}
            >
              {isComplete && !isCurrent ? (
                <div className="flex items-center gap-2">
                  {isSaved ? (
                    <Check className="size-4 shrink-0 text-green-600" />
                  ) : (
                    <span className="size-4 shrink-0 rounded-full border border-amber-400" />
                  )}
                  <p className="text-sm text-foreground">{seg.content}</p>
                </div>
              ) : isCurrent ? (
                <div className="space-y-2">
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">
                      {Math.floor(seg.startTimestamp / 60)}:{(seg.startTimestamp % 60).toString().padStart(2, "0")}
                      {" - "}
                      {Math.floor(seg.endTimestamp / 60)}:{(seg.endTimestamp % 60).toString().padStart(2, "0")}
                    </span>
                  </div>
                  {showHints && (
                    <p className="text-xs text-muted-foreground">{seg.vietnamese || seg.content.slice(0, 20) + "..."}</p>
                  )}
                  <Input
                    value={inputs[i]}
                    onChange={(e) => {
                      const next = [...inputs]
                      next[i] = e.target.value
                      setInputs(next)
                    }}
                    placeholder="Gõ những gì bạn nghe được..."
                    className="w-full"
                    autoFocus
                  />
                  <div className="flex gap-2">
                    <Button
                      size="xs"
                      onClick={() => {
                        const newMaxLine = Math.max(maxLine, currentLine + 1)
                        setMaxLine(newMaxLine)
                        const seg = (transcripts ?? [])[currentLine]
                        const userInput = (inputs[currentLine] ?? "").trim().toLowerCase()
                        const correctAnswer = seg?.content.trim().toLowerCase()
                        if (seg && userInput === correctAnswer) {
                          startTransition(async () => {
                            const res = await postDictationStatus(seg.id, lessonId!)
                            if (res.error) {
                              toast.error(res.error.message)
                            }
                          })
                        }
                        onNext?.()
                      }}
                    >
                      Kiểm tra
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setShowHints(!showHints)}>
                      <Lightbulb className="mr-1 size-3" />
                      Gợi ý
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setInputs((transcripts ?? []).map(() => ""))}>
                      <RotateCcw className="mr-1 size-3" />
                      Làm lại
                    </Button>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground line-through">
                  {seg.content.replace(/[a-zA-Z0-9]/g, "\u25B8")}
                </p>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}
