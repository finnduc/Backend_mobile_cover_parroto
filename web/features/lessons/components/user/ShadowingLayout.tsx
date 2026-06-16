"use client"

import { useState, type ReactNode } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { VidstackYoutubePlayer } from "@/features/lessons/components/user/VidstackYoutubePlayer"
import { VideoControlBar } from "@/features/lessons/components/user/VideoControlBar"
import { TranscriptSidebar } from "@/features/lessons/components/user/TranscriptSidebar"
import { LessonProvider } from "@/features/lessons/context/LessonProvider"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import type { Transcript } from "@/types/lessons.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"

function ConnectedVideoPlayer({ videoUrl }: { videoUrl: string }) {
  const { playerRef } = usePlayerContext()
  return <VidstackYoutubePlayer videoUrl={videoUrl} ref={playerRef} />
}

export function ShadowingLayout({
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
  const [videoLarge, setVideoLarge] = useState(false)

  return (
    <LessonProvider transcripts={transcripts} initialActiveIndex={initialActiveIndex}>
      <div className="flex gap-6">
        <div className="flex flex-1 flex-col gap-4">
          {videoUrl && <ConnectedVideoPlayer videoUrl={videoUrl} />}
          <VideoControlBar
            videoLarge={videoLarge}
            onVideoLargeChange={setVideoLarge}
          />
          <Card>
            <CardHeader>
              <CardTitle>Bài tập Shadowing</CardTitle>
            </CardHeader>
            <CardContent>{children}</CardContent>
          </Card>
        </div>
        <TranscriptSidebar
          transcripts={transcripts}
          completedIds={completedTranscriptIds}
          lessonId={lessonId}
          decks={decks}
          mode="full"
        />
      </div>
    </LessonProvider>
  )
}
