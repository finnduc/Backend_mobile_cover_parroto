"use client"

import { useRef, useState, useEffect, useCallback } from "react"
import { useMediaRemote, useMediaState, type MediaPlayerInstance } from "@vidstack/react"
import type { Transcript } from "@/types/lessons.models"

export function useLessonPlayer(transcripts: Transcript[]) {
  const playerRef = useRef<MediaPlayerInstance>(null)
  const remote = useMediaRemote(playerRef)
  const currentTime = useMediaState("currentTime", playerRef)
  const paused = useMediaState("paused", playerRef)

  const [playerMode, setPlayerMode] = useState<"normal" | "transcript">("normal")
  const [activeIndex, setActiveIndex] = useState(-1)

  // Auto-stop at segment end in transcript mode
  useEffect(() => {
    if (playerMode !== "transcript" || activeIndex < 0 || activeIndex >= transcripts.length) return
    if (currentTime >= transcripts[activeIndex].endTimestamp) {
      remote.pause()
    }
  }, [currentTime, playerMode, activeIndex, transcripts, remote])

  const handleTranscriptClick = useCallback(
    (index: number) => {
      const seg = transcripts[index]
      if (!seg) return
      setActiveIndex(index)
      remote.seek(seg.startTimestamp)
      remote.play()
    },
    [transcripts, remote],
  )

  const handlePlay = useCallback(() => remote.play(), [remote])
  const handlePause = useCallback(() => remote.pause(), [remote])

  const handleReplay = useCallback(() => {
    const seg = transcripts[activeIndex]
    if (!seg) return
    remote.seek(seg.startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, remote])

  const handleNext = useCallback(() => {
    if (activeIndex >= transcripts.length - 1) return
    const nextIndex = activeIndex + 1
    setActiveIndex(nextIndex)
    remote.seek(transcripts[nextIndex].startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, remote])

  const handlePrev = useCallback(() => {
    if (activeIndex <= 0) return
    const prevIndex = activeIndex - 1
    setActiveIndex(prevIndex)
    remote.seek(transcripts[prevIndex].startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, remote])

  const highlightedIndex = transcripts.findIndex(
    (seg) => currentTime >= seg.startTimestamp && currentTime < seg.endTimestamp,
  )

  return {
    playerRef,
    paused,
    playerMode,
    setPlayerMode,
    activeIndex,
    highlightedIndex,
    handleTranscriptClick,
    handlePlay,
    handlePause,
    handleReplay,
    handleNext,
    handlePrev,
  }
}
