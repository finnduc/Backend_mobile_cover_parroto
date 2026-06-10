"use client"

import { VidstackYoutubePlayer } from "@/features/lessons/components/user/VidstackYoutubePlayer"
import { ModeToggle } from "@/features/lessons/components/user/ModeToggle"
import { TranscriptSidebar } from "@/features/lessons/components/user/TranscriptSidebar"
import { useLessonPlayer } from "@/features/lessons/hooks/use-lesson-player"
import { PlayerContext } from "@/features/lessons/context/player-context"
import type { Transcript } from "@/types/lessons.models"
import type { TranscriptBookmark } from "@/types/book-mark.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"
import type { ReactNode } from "react"

export function LessonLayout({
  children,
  videoUrl,
  transcripts,
  lessonId,
  decks = [],
  initialCompletedIds,
  bookmarks = [],
}: {
  children?: ReactNode
  videoUrl?: string
  transcripts?: Transcript[]
  lessonId?: number
  decks?: VocabularyDeck[]
  initialCompletedIds?: number[]
  bookmarks?: TranscriptBookmark[]
}) {
  const segments = transcripts ?? []
  const completedSet = new Set(initialCompletedIds ?? [])
  const initialActiveIndex = segments.findIndex((_, i) => !completedSet.has(i))
  const playerState = useLessonPlayer(segments, initialActiveIndex)

  return (
    <PlayerContext.Provider value={{
      paused: playerState.paused,
      activeIndex: playerState.activeIndex,
      highlightedIndex: playerState.highlightedIndex,
      playerMode: playerState.playerMode,
      setPlayerMode: playerState.setPlayerMode,
      onPlay: playerState.handlePlay,
      onPause: playerState.handlePause,
      onNext: playerState.handleNext,
      onPrev: playerState.handlePrev,
      onReplay: playerState.handleReplay,
      onTranscriptClick: playerState.handleTranscriptClick,
    }}>
      <div className="flex h-full gap-6">
        <div className="flex flex-1 flex-col gap-4">
          {videoUrl && <VidstackYoutubePlayer videoUrl={videoUrl} ref={playerState.playerRef} />}
          <ModeToggle mode={playerState.playerMode} onChange={playerState.setPlayerMode} />
          {children}
        </div>
        <TranscriptSidebar
          transcripts={segments}
          completedIds={initialCompletedIds ?? []}
          bookmarks={bookmarks}
          lessonId={lessonId!}
          decks={decks}
        />
      </div>
    </PlayerContext.Provider>
  )
}
