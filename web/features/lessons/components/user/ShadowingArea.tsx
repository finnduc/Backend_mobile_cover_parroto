"use client"

import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import { Button } from "@/components/ui/button"
import { Mic, Square, Play, Pause, SkipBack, SkipForward, RotateCcw, Check } from "lucide-react"
import { useState, useTransition } from "react"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { assessPronunciation, updatePronunciationProgress } from "@/features/pronunciation/services/pronunciation.action"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { PronunciationScore } from "@/features/lessons/components/user/PronunciationScore"
import { Badge } from "@/components/ui/badge"
import { toast } from "sonner"
import type { Transcript } from "@/types/lessons.models"
import type { PronunciationAttempt } from "@/types/pronunciation.models"

export interface ExerciseControlProps {
  transcripts: Transcript[]
  activeIndex: number
  highlightedIndex: number
  paused: boolean
  initialCompletedIds: number[]
  lessonId: number
  pronunciationScores?: Map<number, PronunciationAttempt>
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
  pronunciationScores,
  onPlay,
  onPause,
  onNext,
  onPrev,
  onReplay,
}: Partial<ExerciseControlProps>) {
  const [completed, setCompleted] = useState<number[]>(initialCompletedIds ?? [])
  const [scores, setScores] = useState<Map<number, PronunciationAttempt>>(
    pronunciationScores ?? new Map()
  )
  const [assessingIndex, setAssessingIndex] = useState<number | null>(null)
  const { isRecording, startRecording, stopRecording } = useAudioRecorder()
  const [, startTransition] = useTransition()

  const lines = (transcripts ?? []).map((t) => t.content)
  const safeActiveIndex = (activeIndex ?? -1) >= 0 && (activeIndex ?? -1) < lines.length ? activeIndex! : 0

  const handleRecord = async () => {
    if (isRecording) {
      try {
        const blob = await stopRecording()
        setAssessingIndex(safeActiveIndex)
        const seg = (transcripts ?? [])[safeActiveIndex]
        if (seg && lessonId) {
          const formData = new FormData()
          formData.append("audio", blob, "recording.webm")
          formData.append("referenceText", seg.content)
          formData.append("lessonId", String(lessonId))
          formData.append("transcriptId", String(seg.id))

          const res = await assessPronunciation(formData)
          if (res.error) {
            toast.error(res.error.message)
          } else if (res.data) {
            setScores((prev) => new Map(prev).set(safeActiveIndex, res.data!))
            await updatePronunciationProgress(seg.id)
            if (!completed.includes(safeActiveIndex)) {
              setCompleted([...completed, safeActiveIndex])
              await postShadowingStatus(seg.id, lessonId)
            }
            if (res.data.overallScore >= 80) {
              toast.success("Pronunciation score: " + res.data.overallScore.toFixed(0) + " - Grade A")
            }
          }
        }
      } catch {
        toast.error("Failed to process recording")
      } finally {
        setAssessingIndex(null)
      }
    } else {
      await startRecording()
    }
  }

  const currentScore = scores.get(safeActiveIndex)

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
                  opacity: isRecording ? 1 : 0.3,
                }}
              />
            ))}
          </div>
        </div>
        <p className="mb-6 text-lg font-medium">{lines[safeActiveIndex]}</p>
        <div className="flex items-center justify-center gap-3">
          <Button
            size="lg"
            variant={isRecording ? "destructive" : "default"}
            className="size-14 rounded-full"
            onClick={handleRecord}
            disabled={assessingIndex !== null}
          >
            {assessingIndex !== null ? (
              <span className="animate-pulse text-xs">...</span>
            ) : isRecording ? (
              <Square className="size-6" />
            ) : (
              <Mic className="size-6" />
            )}
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

      {currentScore && <PronunciationScore result={currentScore} />}

      <div className="space-y-1">
        {(transcripts ?? []).map((seg, i) => {
          const isCompleted = completed.includes(i) && i !== safeActiveIndex
          const isSaved = (initialCompletedIds ?? []).includes(i)
          const segScore = scores.get(i)
          const scoreBadgeClass = segScore
            ? segScore.overallScore >= 80
              ? "bg-green-500"
              : segScore.overallScore >= 60
                ? "bg-amber-500"
                : "bg-red-500"
            : ""
          return (
            <div
              key={seg.id}
              className={`flex items-center gap-2 rounded px-3 py-1.5 text-sm ${
                isSaved ? "bg-green-100 text-green-700 line-through" : ""
              } ${isCompleted && !isSaved ? "bg-amber-50 text-amber-700" : ""} ${i === safeActiveIndex ? "bg-transcript-active font-medium" : ""}`}
            >
              {isSaved && <Check className="size-3.5 shrink-0 text-green-600" />}
              {isCompleted && !isSaved && <span className="size-3.5 shrink-0 rounded-full border border-amber-400" />}
              {segScore && (
                <Badge className={scoreBadgeClass + " text-white text-xs"}>
                  {segScore.overallScore.toFixed(0)}
                </Badge>
              )}
              {seg.content}
            </div>
          )
        })}
      </div>
    </div>
  )
}
