"use client"

import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { RecordControls } from "@/features/lessons/components/user/RecordControls"
import { SpacedTextDisplay } from "@/features/lessons/components/user/SpacedTextDisplay"

export function ShadowingArea({
  transcripts,
  lessonId,
  isAuthenticated = false,
}: ExerciseControlProps) {
  const { activeIndex } = usePlayerContext()
  const idx = activeIndex >= 0 && activeIndex < transcripts.length ? activeIndex : 0
  const seg = transcripts[idx]
  if (!seg) return null

  return (
    <div className="flex flex-col gap-6">
      <SpacedTextDisplay content={seg.content} phonetic={seg.phonetic} />
      <RecordControls
        lessonId={lessonId}
        transcriptId={seg.id}
        referenceText={seg.content}
        isAuthenticated={isAuthenticated}
      />
    </div>
  )
}
