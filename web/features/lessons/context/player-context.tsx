"use client"

import { createContext, useContext } from "react"

interface PlayerContextValue {
  paused: boolean
  activeIndex: number
  highlightedIndex: number
  playerMode: "normal" | "transcript"
  setPlayerMode: (mode: "normal" | "transcript") => void
  onPlay: () => void
  onPause: () => void
  onNext: () => void
  onPrev: () => void
  onReplay: () => void
  onTranscriptClick: (index: number) => void
}

export const PlayerContext = createContext<PlayerContextValue | null>(null)

export function usePlayerContext(): PlayerContextValue {
  const ctx = useContext(PlayerContext)
  if (!ctx) {
    throw new Error("usePlayerContext must be used within PlayerContext provider")
  }
  return ctx
}
