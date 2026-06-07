"use client"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import { PronunciationScore } from "@/features/lessons/components/user/PronunciationScore"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { assessPronunciation, updatePronunciationProgress } from "@/features/pronunciation/services/pronunciation.action"
import type { Transcript } from "@/types/lessons.models"
import type { PronunciationAttempt } from "@/types/pronunciation.models"
import { Check, Mic, Pause, Play, RotateCcw, Send, SkipBack, SkipForward, Square, Volume2 } from "lucide-react"
import { useEffect, useRef, useState, useTransition } from "react"
import { toast } from "sonner"

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

const WAVEFORM_BARS = [
  31, 22, 45, 34, 26, 36, 28, 15, 18, 38,
  19, 45, 29, 28, 42, 20, 27, 35, 40, 37,
  15, 16, 14, 16, 40, 43, 33, 25, 21, 27,
  44, 10, 26, 17, 22, 12, 17, 9, 27, 33,
]

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
  const [recordedBlob, setRecordedBlob] = useState<Blob | null>(null)
  const [recordedUrl, setRecordedUrl] = useState<string | null>(null)
  const [replaying, setReplaying] = useState(false)
  const replayAudioRef = useRef<HTMLAudioElement | null>(null)
  const { isRecording, startRecording, stopRecording } = useAudioRecorder()
  const [, startTransition] = useTransition()

  const lines = (transcripts ?? []).map((t) => t.content)
  const safeActiveIndex = (activeIndex ?? -1) >= 0 && (activeIndex ?? -1) < lines.length ? activeIndex! : 0

  useEffect(() => {
    return () => {
      if (recordedUrl) URL.revokeObjectURL(recordedUrl)
      if (replayAudioRef.current) {
        replayAudioRef.current.pause()
        replayAudioRef.current = null
      }
    }
  }, [])

  const handleRecordToggle = async () => {
    if (isRecording) {
      const blob = await stopRecording()
      if (recordedUrl) URL.revokeObjectURL(recordedUrl)
      const url = URL.createObjectURL(blob)
      setRecordedBlob(blob)
      setRecordedUrl(url)
    } else {
      setRecordedBlob(null)
      if (recordedUrl) {
        URL.revokeObjectURL(recordedUrl)
        setRecordedUrl(null)
      }
      await startRecording()
    }
  }

  const handleReplay = () => {
    if (!recordedUrl) return
    if (replaying && replayAudioRef.current) {
      replayAudioRef.current.pause()
      replayAudioRef.current = null
      setReplaying(false)
      return
    }
    const audio = new Audio(recordedUrl)
    audio.onended = () => {
      setReplaying(false)
      replayAudioRef.current = null
    }
    audio.onerror = () => {
      toast.error("Failed to play recording")
      setReplaying(false)
      replayAudioRef.current = null
    }
    replayAudioRef.current = audio
    setReplaying(true)
    audio.play()
  }

  const handleSubmit = async () => {
    if (!recordedBlob) return
    setAssessingIndex(safeActiveIndex)
    const seg = (transcripts ?? [])[safeActiveIndex]
    if (seg && lessonId) {
      const formData = new FormData()
      formData.append("audio", recordedBlob, "recording.webm")
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
        setRecordedBlob(null)
        if (recordedUrl) {
          URL.revokeObjectURL(recordedUrl)
          setRecordedUrl(null)
        }
      }
    }
    setAssessingIndex(null)
  }

  const currentScore = scores.get(safeActiveIndex)

  return (
    <div className="space-y-4">
      <LessonProgressBar completed={completed.length} total={lines.length} />
      <div className="rounded-xl bg-gradient-to-b from-muted/50 to-background p-6 text-center">
        <div className="mb-4 flex items-center justify-center gap-2">
          <div className="flex h-16 w-64 items-center justify-center gap-1 rounded-full bg-muted px-4">
            {WAVEFORM_BARS.map((height, i) => (
              <div
                key={i}
                className="w-0.5 rounded-full bg-primary"
                style={{
                  height: `${height}px`,
                  opacity: isRecording ? 1 : recordedBlob ? 0.6 : 0.3,
                }}
              />
            ))}
          </div>
        </div>
        <p className="mb-6 text-lg font-medium">{lines[safeActiveIndex]}</p>
        <div className="flex items-center justify-center gap-3">
          <Button
            size="lg"
            variant={isRecording ? "destructive" : recordedBlob ? "secondary" : "default"}
            className="size-14 rounded-full"
            onClick={handleRecordToggle}
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
            onClick={onPrev ?? (() => { })}
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
          <Button size="icon" variant="ghost" onClick={onReplay ?? (() => { })}>
            <RotateCcw className="size-4" />
          </Button>
        </div>

        {recordedBlob && !isRecording && (
          <div className="mt-4 flex items-center justify-center gap-2">
            <Button
              size="sm"
              variant="outline"
              onClick={handleReplay}
            >
              {replaying ? (
                <><Square className="mr-1 size-3" /> Stop</>
              ) : (
                <><Volume2 className="mr-1 size-3" /> Replay</>
              )}
            </Button>
            <Button
              size="sm"
              onClick={handleSubmit}
              disabled={assessingIndex !== null}
            >
              <Send className="mr-1 size-3" />
              Get Score
            </Button>
          </div>
        )}
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
              className={`flex items-center gap-2 rounded px-3 py-1.5 text-sm ${isSaved ? "bg-green-100 text-green-700 line-through" : ""
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
