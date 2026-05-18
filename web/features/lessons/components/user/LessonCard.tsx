"use client"

import { useState } from "react"
import { Card, CardContent } from "@/components/ui/card"
import { DifficultyBadge } from "@/components/common/DifficultyBadge"
import { LessonActionModal } from "@/features/lessons/components/user/LessonActionModal"
import { BookmarkButton } from "@/features/bookmarks/components/BookmarkButton"
import type { Lesson } from "@/types/lessons.models"
import { Play, Clock } from "lucide-react"

function formatDuration(seconds: number): string {
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`
}

export function LessonCard({ lesson, isBookmarked }: { lesson: Lesson; isBookmarked?: boolean }) {
  const [open, setOpen] = useState(false)

  return (
    <>
      <Card className="group cursor-pointer overflow-hidden" onClick={() => setOpen(true)}>
        <div className="relative aspect-video overflow-hidden bg-muted">
          <img
            src={lesson.thumbnailUrl}
            alt={lesson.title}
            className="size-full object-cover transition-transform duration-300 group-hover:scale-105"
          />
          <div className="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 transition-opacity group-hover:opacity-100">
            <Play className="size-10 fill-white text-white" />
          </div>
          {isBookmarked !== undefined && (
            <div className="absolute top-2 right-2" onClick={(e) => e.stopPropagation()}>
              <BookmarkButton lessonId={lesson.id} isBookmarked={isBookmarked} className="bg-black/40 hover:bg-black/60 text-white" />
            </div>
          )}
        </div>
        <CardContent className="space-y-2 p-3">
          <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
            <DifficultyBadge level={lesson.level} />
            <span className="flex items-center gap-0.5">
              <Clock className="size-3" />
              {formatDuration(lesson.duration)}
            </span>
          </div>
          <h3 className="line-clamp-2 text-sm font-medium leading-snug">{lesson.title}</h3>
        </CardContent>
      </Card>
      <LessonActionModal lesson={lesson} open={open} onOpenChange={setOpen} />
    </>
  )
}
