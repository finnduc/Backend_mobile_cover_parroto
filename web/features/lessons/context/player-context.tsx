"use client"

import { createContext, useContext, type RefObject } from "react"
import type { MediaPlayerInstance } from "@vidstack/react"

interface PlayerContextValue {
  playerRef: RefObject<MediaPlayerInstance | null>
  paused: boolean
  activeIndex: number
  highlightedIndex: number
  autoStop: boolean
  setAutoStop: (v: boolean) => void
  transcriptMode: "masked" | "full"
  setTranscriptMode: (v: "masked" | "full") => void
  sidebarVisible: boolean
  setSidebarVisible: (v: boolean) => void
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
