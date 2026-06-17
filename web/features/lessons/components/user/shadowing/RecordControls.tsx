"use client"

import { useEffect, useRef, useState } from "react"
import { Button } from "@/components/ui/button"
import { useAudioRecorder } from "@/features/lessons/hooks/use-audio-recorder"
import { Mic, Square, Volume2 } from "lucide-react"
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

  return (
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
      <Button onClick={handleToggleRecord}>
        {isRecording ? (
          <>
            <Square data-icon="inline-start" />
            Dừng ghi
          </>
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
