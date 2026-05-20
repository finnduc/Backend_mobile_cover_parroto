"use client"

import { useState } from "react"
import { notFound } from "next/navigation"
import { VidstackYoutubePlayer } from "@/features/lessons/components/user/VidstackYoutubePlayer"
import { ModeToggle } from "@/features/lessons/components/user/ModeToggle"
import { TranscriptLine } from "@/components/common/TranscriptLine"
import { useLessonPlayer } from "@/features/lessons/hooks/use-lesson-player"
import { AddToDeckDialog } from "@/features/lessons/components/user/AddToDeckDialog"
import type { VocabularyDeck } from "@/types/vocabulary.models"
import type { ExerciseControlProps } from "@/features/lessons/components/user/ShadowingArea"
import type { Transcript } from "@/types/lessons.models"
import type { ComponentType } from "react"

export function LessonLayout({
  exercise: ExerciseComponent,
  duration,
  videoUrl,
  transcripts,
  lessonId,
  decks = [],
}: {
  exercise?: ComponentType<Partial<ExerciseControlProps>>
  duration?: number
  videoUrl?: string
  transcripts?: Transcript[]
  lessonId?: number
  decks?: VocabularyDeck[]
}) {
  const segments = transcripts ?? []
  const {
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
  } = useLessonPlayer(segments)

  const [addToDeckTarget, setAddToDeckTarget] = useState<{
    transcriptId: number
    text: string
  } | null>(null)

  return (
    <div className="flex h-full gap-6">
      <div className="flex flex-1 flex-col gap-4">
        {videoUrl ? (
          <VidstackYoutubePlayer videoUrl={videoUrl} ref={playerRef} />
        ) : (
          notFound()
        )}
        <ModeToggle mode={playerMode} onChange={setPlayerMode} />
        <div className="flex-1">
          {ExerciseComponent != null ? (
            <ExerciseComponent
              transcripts={segments}
              activeIndex={activeIndex}
              highlightedIndex={highlightedIndex}
              paused={paused}
              onPlay={handlePlay}
              onPause={handlePause}
              onNext={handleNext}
              onPrev={handlePrev}
              onReplay={handleReplay}
              onTranscriptClick={handleTranscriptClick}
            />
          ) : null}
        </div>
      </div>
      <aside className="hidden w-80 shrink-0 lg:block">
        <div className="sticky top-6 space-y-3">
          <h2 className="text-sm font-semibold text-muted-foreground">Transcript</h2>
          <div className="max-h-[calc(100vh-12rem)] space-y-1 overflow-y-auto rounded-xl border p-2">
            {segments.map((seg, i) => (
              <TranscriptLine
                key={seg.id}
                text={seg.content}
                active={i === highlightedIndex}
                transcriptId={seg.id}
                timestamp={`${Math.floor(seg.startTimestamp / 60)}:${(seg.startTimestamp % 60).toString().padStart(2, "0")}`}
                onClick={() => handleTranscriptClick(i)}
                onAddToDeck={(transcriptId, text) => {
                  setAddToDeckTarget({ transcriptId, text })
                }}
              />
            ))}
          </div>
        </div>
      </aside>

      {addToDeckTarget && lessonId != null && (
        <AddToDeckDialog
          open={!!addToDeckTarget}
          onOpenChange={(open) => {
            if (!open) setAddToDeckTarget(null)

          }}
          phrase={addToDeckTarget.text}
          lessonId={lessonId}
          transcriptId={addToDeckTarget.transcriptId}
          decks={decks}
        />
      )}
    </div>
  )
}
