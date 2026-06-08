"use client"

import { useState, useTransition } from "react"
import { Button } from "@/components/ui/button"
import { Bookmark } from "lucide-react"
import { cn } from "@/lib/utils"
import { toggleLessonBookmark } from "@/features/lesson-bookmarks/services/lesson-bookmarks.action"
import { toast } from "sonner"

export function LessonBookmarkButton({
  lessonId,
  isBookmarked,
  className,
}: {
  lessonId: number
  isBookmarked: boolean
  className?: string
}) {
  const [bookmarked, setBookmarked] = useState(isBookmarked)
  const [isPending, startTransition] = useTransition()

  const handleToggle = () => {
    startTransition(async () => {
      const next = !bookmarked
      setBookmarked(next)

      const res = await toggleLessonBookmark(lessonId)
      if (res.error) {
        setBookmarked(!next)
        toast.error(res.error.message)
      } else {
        toast.success(next ? "Da luu bai hoc" : "Da bo luu bai hoc")
      }
    })
  }

  return (
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
  )
}
