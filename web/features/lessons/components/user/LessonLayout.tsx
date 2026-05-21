"use client"

import { notFound } from "next/navigation"
import { VidstackYoutubePlayer } from "@/features/lessons/components/user/VidstackYoutubePlayer"
import { ModeToggle } from "@/features/lessons/components/user/ModeToggle"
import { TranscriptLine } from "@/components/common/TranscriptLine"
import { useLessonPlayer } from "@/features/lessons/hooks/use-lesson-player"
import type { ExerciseControlProps } from "@/features/lessons/components/user/ShadowingArea"
import type { Transcript } from "@/types/lessons.models"
import type { ComponentType } from "react"

export function LessonLayout({
  exercise: ExerciseComponent,
  duration,
  videoUrl,
  transcripts,
  initialCompletedIds,
  lessonId,
}: {
  exercise?: ComponentType<Partial<ExerciseControlProps>>
  duration?: number
  videoUrl?: string
  transcripts?: Transcript[]
  initialCompletedIds?: number[]
  lessonId?: number
}) {
  const segments = transcripts ?? []
  const completedSet = new Set(initialCompletedIds ?? [])
  const initialActiveIndex = segments.findIndex((_, i) => !completedSet.has(i))
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
  } = useLessonPlayer(segments, initialActiveIndex)

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
              initialCompletedIds={initialCompletedIds ?? []}
              lessonId={lessonId}
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
                timestamp={`${Math.floor(seg.startTimestamp / 60)}:${(seg.startTimestamp % 60).toString().padStart(2, "0")}`}
                onClick={() => handleTranscriptClick(i)}
              />
            ))}
          </div>
        </div>
      </aside>
    </div>
  )
}
