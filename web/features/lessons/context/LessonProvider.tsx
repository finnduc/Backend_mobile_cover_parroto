"use client"

import type { ReactNode } from "react"
import { PlayerContext } from "@/features/lessons/context/player-context"
import { useLessonPlayer } from "@/features/lessons/hooks/use-lesson-player"
import type { Transcript } from "@/types/lessons.models"

export function LessonProvider({
  transcripts,
  initialActiveIndex = -1,
  children,
}: {
  transcripts: Transcript[]
  initialActiveIndex?: number
  children: ReactNode
}) {
  const player = useLessonPlayer(transcripts, initialActiveIndex)
  return (
    <PlayerContext.Provider
      value={{
        playerRef: player.playerRef,
        paused: player.paused,
        activeIndex: player.activeIndex,
        highlightedIndex: player.highlightedIndex,
        autoStop: player.autoStop,
        setAutoStop: player.setAutoStop,
        transcriptMode: player.transcriptMode,
        setTranscriptMode: player.setTranscriptMode,
        sidebarVisible: player.sidebarVisible,
        setSidebarVisible: player.setSidebarVisible,
        onPlay: player.handlePlay,
        onPause: player.handlePause,
        onNext: player.handleNext,
        onPrev: player.handlePrev,
        onReplay: player.handleReplay,
        onTranscriptClick: player.handleTranscriptClick,
      }}
    >
      {children}
    </PlayerContext.Provider>
  )
}
