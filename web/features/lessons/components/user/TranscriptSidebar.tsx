"use client"

import { useState } from "react"
import { TranscriptLine } from "@/components/common/TranscriptLine"
// import { BookmarkNoteDialog } from "@/features/bookmarks/components/BookmarkNoteDialog"
import { AddToDeckDialog } from "@/features/lessons/components/user/AddToDeckDialog"
import { Progress } from "@/components/ui/progress"
import { usePlayerContext } from "@/features/lessons/context/player-context"
// import { createTranscriptBookmark, deleteTranscriptBookmark } from "@/features/bookmarks/services/bookmarks.action"
import type { Transcript } from "@/types/lessons.models"
// import type { TranscriptBookmark } from "@/types/book-mark.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function TranscriptSidebar({
  transcripts,
  completedIds,
  // bookmarks = [],
  lessonId,
  decks = [],
}: {
  transcripts: Transcript[]
  completedIds: number[]
  // bookmarks?: TranscriptBookmark[]
  lessonId: number
  decks?: VocabularyDeck[]
}) {
  const { highlightedIndex, onTranscriptClick } = usePlayerContext()

  const [addToDeckTarget, setAddToDeckTarget] = useState<{
    transcriptId: number
    text: string
  } | null>(null)

  // const [noteDialogOpen, setNoteDialogOpen] = useState(false)
  // const [activeBookmarkId, setActiveBookmarkId] = useState<number | null>(null)

  // const handleBookmarkToggle = async (transcriptId: number) => {
  //   const existing = bookmarks.find((b) => b.transcriptId === transcriptId)
  //   if (existing) {
  //     await deleteTranscriptBookmark(existing.id)
  //   } else {
  //     const res = await createTranscriptBookmark(transcriptId, "")
  //     if (res.data) {
  //       setActiveBookmarkId(res.data.id)
  //       setNoteDialogOpen(true)
  //     }
  //   }
  // }

  return (
    <aside className="hidden w-80 shrink-0 lg:block">
      <div className="sticky top-6 space-y-3">
        <h2 className="text-sm font-semibold text-muted-foreground">Transcript</h2>
        <div className="max-h-[calc(100vh-12rem)] space-y-1 overflow-y-auto rounded-xl border p-2">
          {transcripts.map((seg, i) => (
            <TranscriptLine
              key={seg.id}
              text={seg.content}
              phonetic={seg.phonetic}
              active={i === highlightedIndex}
              completed={completedIds.includes(seg.id)}
              // bookmarked={bookmarks.some((b) => b.transcriptId === seg.id)}
              transcriptId={seg.id}
              timestamp={`${Math.floor(seg.startTimestamp / 60)}:${(seg.startTimestamp % 60).toString().padStart(2, "0")}`}
              onClick={() => onTranscriptClick(i)}
              onAddToDeck={(transcriptId, text) => {
                setAddToDeckTarget({ transcriptId, text })
              }}
              // onBookmark={() => handleBookmarkToggle(seg.id)}
            />
          ))}
        </div>
        <div className="space-y-2 pt-2">
          <div className="flex items-center justify-between text-xs text-muted-foreground">
            <span>Tiến độ</span>
            <span>{completedIds.length}/{transcripts.length} câu</span>
          </div>
          <Progress
            value={transcripts.length > 0 ? (completedIds.length / transcripts.length) * 100 : 0}
          />
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

      {/* {activeBookmarkId && (
        <BookmarkNoteDialog
          open={noteDialogOpen}
          onOpenChange={setNoteDialogOpen}
          onSave={async (note: string) => {
            const { updateTranscriptBookmarkNote } = await import("@/features/bookmarks/services/bookmarks.action")
            await updateTranscriptBookmarkNote(activeBookmarkId, note)
          }}
        />
      )} */}
    </aside>
  )
}
