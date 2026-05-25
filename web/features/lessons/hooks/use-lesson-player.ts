"use client"

import { useRef, useState, useEffect, useCallback } from "react"

import { useMediaRemote, useMediaState, type MediaPlayerInstance } from "@vidstack/react"
import type { Transcript } from "@/types/lessons.models"

export function useLessonPlayer(transcripts: Transcript[], initialActiveIndex = -1) {
  const playerRef = useRef<MediaPlayerInstance>(null)
  const remote = useMediaRemote(playerRef)
  const currentTime = useMediaState("currentTime", playerRef)
  const paused = useMediaState("paused", playerRef)

  const [playerMode, setPlayerMode] = useState<"normal" | "transcript">("normal")
  const [activeIndex, setActiveIndex] = useState(initialActiveIndex)
  const prevModeRef = useRef(playerMode)

  const highlightedIndex = transcripts.findIndex(
    (seg) => currentTime >= seg.startTimestamp && currentTime < seg.endTimestamp,
  )

  // Auto-stop at segment end in transcript mode
  useEffect(() => {
    if (playerMode !== "transcript" || activeIndex < 0 || activeIndex >= transcripts.length) return
    if (currentTime >= transcripts[activeIndex].endTimestamp) {
      remote.pause()
    }
  }, [currentTime, playerMode, activeIndex, transcripts, remote])

  // Replay current segment when switching to transcript mode
  useEffect(() => {
    const justSwitched = playerMode === "transcript" && prevModeRef.current !== "transcript"
    prevModeRef.current = playerMode
    if (!justSwitched) return
    const idx = activeIndex >= 0 ? activeIndex : highlightedIndex
    if (idx < 0) return
    const seg = transcripts[idx]
    if (!seg) return
    remote.seek(seg.startTimestamp)
    remote.play()
  }, [playerMode, activeIndex, highlightedIndex, transcripts, remote])

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
    const idx = activeIndex >= 0 ? activeIndex : highlightedIndex
    const seg = transcripts[idx]
    if (!seg) return
    remote.seek(seg.startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, highlightedIndex, remote])

  const handleNext = useCallback(() => {
    const idx = activeIndex >= 0 ? activeIndex : highlightedIndex
    if (idx >= transcripts.length - 1) return
    const nextIndex = idx + 1
    setActiveIndex(nextIndex)
    remote.seek(transcripts[nextIndex].startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, highlightedIndex, remote])

  const handlePrev = useCallback(() => {
    const idx = activeIndex >= 0 ? activeIndex : highlightedIndex
    if (idx <= 0) return
    const prevIndex = idx - 1
    setActiveIndex(prevIndex)
    remote.seek(transcripts[prevIndex].startTimestamp)
    remote.play()
  }, [transcripts, activeIndex, highlightedIndex, remote])

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
