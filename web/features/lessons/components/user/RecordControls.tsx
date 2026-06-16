"use client"

import { useEffect, useRef, useState } from "react"
import { Button } from "@/components/ui/button"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { postShadowingStatus } from "@/features/lessons/services/shadowing-status.action"
import { assessPronunciation } from "@/features/pronunciation/services/pronunciation.action"
import { Mic, Square, Volume2 } from "lucide-react"
import { toast } from "sonner"

type Props = {
  lessonId: number
  transcriptId: number
  referenceText: string
  isAuthenticated: boolean
  onScored?: () => void
}

export function RecordControls({
  lessonId,
  transcriptId,
  referenceText,
  isAuthenticated,
  onScored,
}: Props) {
  const { isRecording, startRecording, stopRecording } = useAudioRecorder()
  const [blob, setBlob] = useState<Blob | null>(null)
  const [url, setUrl] = useState<string | null>(null)
  const [replaying, setReplaying] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const audioRef = useRef<HTMLAudioElement | null>(null)

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
    } else {
      if (url) URL.revokeObjectURL(url)
      setBlob(null)
      setUrl(null)
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

  const handleSubmit = async () => {
    if (!blob || !isAuthenticated) return
    setSubmitting(true)
    try {
      const fd = new FormData()
      fd.append("audio", blob, "recording.webm")
      fd.append("referenceText", referenceText)
      fd.append("lessonId", String(lessonId))
      fd.append("transcriptId", String(transcriptId))
      const res = await assessPronunciation(fd)
      if (res.error) {
        toast.error(res.error.message)
      } else if (res.data) {
        if (res.data.overallScore >= 80) {
          toast.success(
            `Pronunciation score: ${res.data.overallScore.toFixed(0)} - Grade A`,
          )
        }
        await postShadowingStatus(transcriptId, lessonId)
        onScored?.()
      }
    } finally {
      setSubmitting(false)
      setBlob(null)
      if (url) {
        URL.revokeObjectURL(url)
        setUrl(null)
      }
    }
  }

  return (
    <div className="flex items-center justify-center gap-3">
      <Button
        variant="outline"
        onClick={handleReplay}
        disabled={!blob || isRecording || submitting}
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
      <Button
        onClick={blob && !isRecording ? handleSubmit : handleToggleRecord}
        disabled={submitting}
      >
        {submitting ? (
          <span className="text-xs">...</span>
        ) : isRecording ? (
          <>
            <Square data-icon="inline-start" />
            Dừng ghi
          </>
        ) : blob ? (
          "Gửi"
        ) : (
          <>
            <Mic data-icon="inline-start" />
            Ghi âm
          </>
        )}
      </Button>
    </div>
  )
}
