"use client"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { PronunciationScore } from "@/features/lessons/components/user/PronunciationScore"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { PlaybackControls } from "@/features/lessons/components/user/PlaybackControls"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { assessPronunciation, updatePronunciationProgress } from "@/features/pronunciation/services/pronunciation.action"
import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import type { PronunciationAttempt } from "@/types/pronunciation.models"
import { Check, ChevronDown, ChevronUp, Mic, Send, Square, Volume2 } from "lucide-react"
import { useEffect, useRef, useState } from "react"
import { toast } from "sonner"

const WAVEFORM_BARS = [
  31, 22, 45, 34, 26, 36, 28, 15, 18, 38,
  19, 45, 29, 28, 42, 20, 27, 35, 40, 37,
  15, 16, 14, 16, 40, 43, 33, 25, 21, 27,
  44, 10, 26, 17, 22, 12, 17, 9, 27, 33,
]

export function ShadowingArea({
  transcripts,
  initialCompletedIds,
  lessonId,
  isAuthenticated,
  pronunciationScores,
}: ExerciseControlProps) {
  const [completed, setCompleted] = useState<number[]>(initialCompletedIds ?? [])
  const [scores, setScores] = useState<Map<number, PronunciationAttempt>>(
    pronunciationScores ?? new Map()
  )
  const [assessingIndex, setAssessingIndex] = useState<number | null>(null)
  const [recordedBlob, setRecordedBlob] = useState<Blob | null>(null)
  const [recordedUrl, setRecordedUrl] = useState<string | null>(null)
  const [replaying, setReplaying] = useState(false)
  const [expandedScoreIndex, setExpandedScoreIndex] = useState<number | null>(null)
  const replayAudioRef = useRef<HTMLAudioElement | null>(null)
  const { isRecording, startRecording, stopRecording } = useAudioRecorder()
  const { activeIndex } = usePlayerContext()

  const lines = transcripts.map((t) => t.content)
  const safeActiveIndex = activeIndex >= 0 && activeIndex < lines.length ? activeIndex : 0

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
    if (!recordedBlob || !isAuthenticated) return
    setAssessingIndex(safeActiveIndex)
    const seg = transcripts[safeActiveIndex]
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

  return (
    <div className="space-y-4">
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
        <p className="mb-1 text-lg font-medium">{lines[safeActiveIndex]}</p>
        {transcripts[safeActiveIndex]?.phonetic && (
          <p className="mb-5 text-sm text-muted-foreground">/{transcripts[safeActiveIndex].phonetic}/</p>
        )}
        {!transcripts[safeActiveIndex]?.phonetic && <div className="mb-6" />}
        <PlaybackControls totalLines={lines.length}>
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
        </PlaybackControls>

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

      <div className="space-y-1">
        {transcripts.map((seg, i) => {
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
          const isExpanded = expandedScoreIndex === i
          return (
            <div key={seg.id}>
              <div
                className={`flex items-center gap-2 rounded px-3 py-1.5 text-sm ${isSaved ? "bg-green-100 text-green-700 line-through" : ""
                  } ${isCompleted && !isSaved ? "bg-amber-50 text-amber-700" : ""} ${i === safeActiveIndex ? "bg-transcript-active font-medium" : ""}`}
              >
                {isSaved && <Check className="size-3.5 shrink-0 text-green-600" />}
                {isCompleted && !isSaved && <span className="size-3.5 shrink-0 rounded-full border border-amber-400" />}
                {segScore && (
                  <button
                    type="button"
                    onClick={() => setExpandedScoreIndex(isExpanded ? null : i)}
                    className="flex items-center gap-1 cursor-pointer"
                  >
                    <Badge className={scoreBadgeClass + " text-white text-xs"}>
                      {segScore.overallScore.toFixed(0)}
                    </Badge>
                    {isExpanded ? (
                      <ChevronUp className="size-3 text-muted-foreground" />
                    ) : (
                      <ChevronDown className="size-3 text-muted-foreground" />
                    )}
                  </button>
                )}
                <div className="min-w-0">
                  <span>{seg.content}</span>
                  {seg.phonetic && (
                    <span className="block text-xs text-muted-foreground">/{seg.phonetic}/</span>
                  )}
                </div>
              </div>
              {isExpanded && segScore && (
                <div className="px-3 py-2">
                  <PronunciationScore result={segScore} />
                </div>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}
