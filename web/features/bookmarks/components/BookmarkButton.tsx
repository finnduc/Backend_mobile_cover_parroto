"use client"

import { useState, useTransition } from "react"
import { Button } from "@/components/ui/button"
import { Bookmark } from "lucide-react"
import { cn } from "@/lib/utils"
import {
  createTranscriptBookmark,
  deleteTranscriptBookmark,
  updateTranscriptBookmarkNote,
} from "@/features/bookmarks/services/bookmarks.action"
import { BookmarkNoteDialog } from "@/features/bookmarks/components/BookmarkNoteDialog"
import { toast } from "sonner"

export function BookmarkButton({
  transcriptId,
  bookmarkId: initialBookmarkId,
  isBookmarked,
  className,
}: {
  transcriptId: number
  bookmarkId?: number | null
  isBookmarked: boolean
  className?: string
}) {
  const [bookmarked, setBookmarked] = useState(isBookmarked)
  const [currentBookmarkId, setCurrentBookmarkId] = useState<number | null>(
    initialBookmarkId ?? null
  )
  const [noteOpen, setNoteOpen] = useState(false)
  const [isPending, startTransition] = useTransition()

  const handleToggle = () => {
    startTransition(async () => {
      const next = !bookmarked
      setBookmarked(next)

      if (next) {
        const res = await createTranscriptBookmark(transcriptId, "")
        if (res.error) {
          setBookmarked(!next)
          toast.error(res.error.message)
        } else {
          if (res.data) setCurrentBookmarkId(res.data.id)
          setNoteOpen(true)
        }
      } else {
        if (!currentBookmarkId) return
        const res = await deleteTranscriptBookmark(currentBookmarkId)
        if (res.error) {
          setBookmarked(!next)
          toast.error(res.error.message)
        } else {
          setCurrentBookmarkId(null)
          toast.success("Da bo luu bai hoc")
        }
      }
    })
  }

  const handleSaveNote = (note: string) => {
    if (!currentBookmarkId) return
    startTransition(async () => {
      await updateTranscriptBookmarkNote(currentBookmarkId, note)
    })
  }

  return (
    <>
      <Button
        variant="ghost"
        size="icon"
        className={cn("size-8", className)}
        disabled={isPending}
        onClick={(e) => {
          e.stopPropagation()
          handleToggle()
        }}
        aria-label={bookmarked ? "Bo luu bai hoc" : "Luu bai hoc"}
      >
        <Bookmark
          className={cn("size-4", bookmarked && "fill-current text-yellow-500")}
        />
      </Button>
      {bookmarked && (
        <BookmarkNoteDialog
          open={noteOpen}
          onOpenChange={setNoteOpen}
          onSave={handleSaveNote}
        />
      )}
    </>
  )
}
