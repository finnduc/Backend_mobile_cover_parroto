"use client"

import { type ReactNode } from "react"
import { Card, CardContent } from "@/components/ui/card"
import { PlaybackControls } from "@/features/lessons/components/user/PlaybackControls"
import { DictationControlBar } from "@/features/lessons/components/user/DictationControlBar"
import { TranscriptSidebar } from "@/features/lessons/components/user/TranscriptSidebar"
import { LessonProvider } from "@/features/lessons/context/LessonProvider"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { VideoPanel } from "@/features/lessons/components/user/VideoPanel"
import type { Transcript } from "@/types/lessons.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"

function DictationContent({
  videoUrl,
  transcripts,
  lessonId,
  decks,
  completedTranscriptIds,
  children,
}: {
  videoUrl?: string
  transcripts: Transcript[]
  lessonId: number
  decks: VocabularyDeck[]
  completedTranscriptIds: number[]
  children: ReactNode
}) {
  const { transcriptMode, sidebarVisible } = usePlayerContext()

  return (
    <div className="grid gap-6 lg:grid-cols-[1fr_18rem_20rem]">
      <div className="flex flex-col gap-4">
        <VideoPanel videoUrl={videoUrl} />
        <DictationControlBar />
      </div>

      <Card>
        <CardContent className="flex flex-col gap-4">
          <PlaybackControls totalLines={transcripts.length}>
            <div />
          </PlaybackControls>
          {children}
        </CardContent>
      </Card>

      {sidebarVisible && (
        <TranscriptSidebar
          transcripts={transcripts}
          completedIds={completedTranscriptIds}
          lessonId={lessonId}
          decks={decks}
          mode={transcriptMode}
        />
      )}
    </div>
  )
}

export function DictationLayout({
  videoUrl,
  transcripts = [],
  lessonId,
  decks = [],
  completedTranscriptIds = [],
  children,
}: {
  videoUrl?: string
  transcripts?: Transcript[]
  lessonId: number
  decks?: VocabularyDeck[]
  completedTranscriptIds?: number[]
  children: ReactNode
}) {
  const completedSet = new Set(completedTranscriptIds)
  const initialActiveIndex = transcripts.findIndex((t) => !completedSet.has(t.id))

  return (
    <LessonProvider transcripts={transcripts} initialActiveIndex={initialActiveIndex}>
      <DictationContent
        videoUrl={videoUrl}
        transcripts={transcripts}
        lessonId={lessonId}
        decks={decks}
        completedTranscriptIds={completedTranscriptIds}
      >
        {children}
      </DictationContent>
    </LessonProvider>
  )
}
