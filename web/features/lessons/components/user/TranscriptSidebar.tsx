"use client"

import { useState } from "react"
import { TranscriptLine } from "@/components/common/TranscriptLine"
import { AddToDeckDialog } from "@/features/lessons/components/user/AddToDeckDialog"
import { Progress } from "@/components/ui/progress"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import type { Transcript } from "@/types/lessons.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function TranscriptSidebar({
  transcripts,
  completedIds,
  lessonId,
  decks = [],
  mode = "full",
}: {
  transcripts: Transcript[]
  completedIds: number[]
  lessonId: number
  decks?: VocabularyDeck[]
  mode?: "full" | "masked"
}) {
  const { activeIndex, highlightedIndex, onTranscriptClick } = usePlayerContext()
  const [revealedIds, setRevealedIds] = useState<Set<number>>(new Set())

  const [addToDeckTarget, setAddToDeckTarget] = useState<{
    transcriptId: number
    text: string
  } | null>(null)

  const activeHighlight = activeIndex >= 0 ? activeIndex : highlightedIndex

  const percent =
    transcripts.length > 0
      ? Math.round((completedIds.length / transcripts.length) * 100)
      : 0

  return (
    <aside className="hidden w-80 shrink-0 lg:block">
      <div className="sticky top-6 flex flex-col gap-3">
        <div className="flex items-center justify-between">
          <h2 className="text-sm font-semibold">BẢN CHÉP</h2>
          <div className="flex items-center gap-2">
            <span className="text-xs text-muted-foreground">{percent}%</span>
          </div>
        </div>
        <div className="max-h-[calc(100vh-12rem)] flex flex-col gap-1 overflow-y-auto rounded-xl border p-2">
          {transcripts.map((seg, i) => (
            <TranscriptLine
              key={seg.id}
              text={seg.content}
              phonetic={seg.phonetic}
              active={i === activeHighlight}
              completed={completedIds.includes(seg.id)}
              transcriptId={seg.id}
              timestamp={`${Math.floor(seg.startTimestamp / 60)}:${(seg.startTimestamp % 60).toString().padStart(2, "0")}`}
              onClick={() => onTranscriptClick(i)}
              onAddToDeck={(transcriptId, text) => {
                setAddToDeckTarget({ transcriptId, text })
              }}
              masked={mode === "masked"}
              revealed={revealedIds.has(seg.id)}
              onReveal={
                mode === "masked"
                  ? () =>
                      setRevealedIds((prev) => {
                        const next = new Set(prev)
                        if (next.has(seg.id)) {
                          next.delete(seg.id)
                        } else {
                          next.add(seg.id)
                        }
                        return next
                      })
                  : undefined
              }
            />
          ))}
        </div>
        <div className="flex flex-col gap-2 pt-2">
          <div className="flex items-center justify-between text-xs text-muted-foreground">
            <span>Tiến độ</span>
            <span>
              {completedIds.length}/{transcripts.length} câu
            </span>
          </div>
          <Progress value={percent} />
        </div>
      </div>

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
    </aside>
  )
}
