"use client"

import { Button } from "@/components/ui/button"
import { WordDiffResult } from "@/features/lessons/components/user/WordDiffResult"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { postShadowingTranscribe } from "@/features/lessons/services/shadowing-transcribe.action"
import { computeWordDiff } from "@/features/lessons/utils/word-diff"
import { cn, splitWords } from "@/lib/utils"
import { Mic, Square, Volume2 } from "lucide-react"
import { useEffect, useRef, useState, useTransition } from "react"
import { toast } from "sonner"

type Props = {
  lessonId: number
  transcriptId: number
  referenceText: string
  isAuthenticated: boolean
}

export function RecordControls({
  lessonId,
  transcriptId,
  referenceText,
  isAuthenticated,
}: Props) {
  const { isRecording, startRecording, stopRecording } = useAudioRecorder()
  const [blob, setBlob] = useState<Blob | null>(null)
  const [url, setUrl] = useState<string | null>(null)
  const [replaying, setReplaying] = useState(false)
  const audioRef = useRef<HTMLAudioElement | null>(null)
  const [isPending, startTransition] = useTransition()

  const [wordDiff, setWordDiff] = useState<
    { word: string; status: "correct" | "wrong" }[] | null
  >(null)
  const [score, setScore] = useState<{ correct: number; total: number } | null>(null)
  const [isComplete, setIsComplete] = useState(false)

  useEffect(() => {
    return () => {
      if (url) URL.revokeObjectURL(url)
      if (audioRef.current) audioRef.current.pause()
    }
  }, [url])

  const handleToggleRecord = async () => {
    if (isRecording) {
      const b = await stopRecording()
      if (url) URL.revokeObjectURL(url)
      setBlob(b)
      setUrl(URL.createObjectURL(b))
      setWordDiff(null)
      setScore(null)
      setIsComplete(false)

      // Transcribe and score
      startTransition(async () => {
        const res = await postShadowingTranscribe(b)

        let transcribedText = ""
        if (res.error) {
          toast.error("Service not available, try again later")
        } else {
          transcribedText = res.data?.transcribedText ?? ""
        }

        const correctWords = splitWords(referenceText)
        const inputWords = transcribedText.trim().split(/\s+/).filter(Boolean)

        // Always show diff, even if empty or all wrong
        const diff = inputWords.length > 0
          ? computeWordDiff(correctWords, inputWords)
          : correctWords.map(word => ({ word, status: "wrong" as const }))

        setWordDiff(diff)

        const correctCount = diff.filter((d) => d.status === "correct").length
        const totalCount = correctWords.length
        setScore({ correct: correctCount, total: totalCount })

        // If 100% match, mark as complete
        if (correctCount === totalCount && totalCount > 0 && isAuthenticated) {
          const statusRes = await postShadowingStatus(transcriptId, lessonId)
          if (statusRes.error) {
            toast.error(statusRes.error.message)
          } else {
            setIsComplete(true)
          }
        }
      })
    } else {
      if (url) URL.revokeObjectURL(url)
      setBlob(null)
      setUrl(null)
      setWordDiff(null)
      setScore(null)
      setIsComplete(false)
      await startRecording()
    }
  }

  const handleReplay = () => {
    if (!url) return
    if (replaying && audioRef.current) {
      audioRef.current.pause()
      audioRef.current = null
      setReplaying(false)
      return
    }
    const audio = new Audio(url)
    audio.onended = () => {
      setReplaying(false)
      audioRef.current = null
    }
    audio.onerror = () => {
      toast.error("Failed to play recording")
      setReplaying(false)
      audioRef.current = null
    }
    audioRef.current = audio
    setReplaying(true)
    audio.play()
  }

  return (
    <div className="flex flex-col gap-4">

      {wordDiff && wordDiff.length > 0 && (
        <WordDiffResult diff={wordDiff} className="items-center justify-center border-none bg-transparent text-xl" />
      )}

      {score && (
        <div className="text-center">
          <span className={cn(
            "text-lg font-semibold",
            score.correct === score.total ? "text-green-600" : "text-red-500"
          )}>
            {score.correct}/{score.total}
          </span>
          {isComplete && (
            <span className="ml-2 text-green-600">Chính xác!</span>
          )}
        </div>
      )}

      <div className="flex items-center justify-center gap-3">
        <Button
          variant="outline"
          onClick={handleReplay}
          disabled={!blob || isRecording}
        >
          {replaying ? (
            <>
              <Square data-icon="inline-start" />
              Dừng
            </>
          ) : (
            <>
              <Volume2 data-icon="inline-start" />
              Phát lại ghi âm
            </>
          )}
        </Button>
        <Button onClick={handleToggleRecord} disabled={isPending}>
          {isRecording ? (
            <>
              <Square data-icon="inline-start" />
              Dừng ghi
            </>
          ) : isPending ? (
            <>
              <span className="animate-spin">⏳</span>
              Đang xử lý...
            </>
          ) : (
            <>
              <Mic data-icon="inline-start" />
              Ghi âm
            </>
          )}
        </Button>
      </div>

    </div>
  )
}
