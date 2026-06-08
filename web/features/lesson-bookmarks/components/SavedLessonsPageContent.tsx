"use client"

import Link from "next/link"
import { Card, CardContent } from "@/components/ui/card"
import { DifficultyBadge } from "@/components/common/DifficultyBadge"
import { LessonBookmarkButton } from "@/features/lesson-bookmarks/components/LessonBookmarkButton"
import { Play, Clock } from "lucide-react"
import type { LessonBookmark } from "@/types/lesson-bookmark.models"
import type { Lesson } from "@/types/lessons.models"

function formatDuration(seconds: number): string {
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`
}

export function SavedLessonsPageContent({
  bookmarks,
}: {
  bookmarks: (LessonBookmark & { lesson?: Lesson })[]
}) {
  if (bookmarks.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <p className="text-lg font-medium">Chua co bai hoc nao duoc luu</p>
        <p className="text-sm">
          Nhan vao bieu tuong bookmark de luu bai hoc
        </p>
      </div>
    )
  }

  return (
    <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
      {bookmarks.map((bookmark) => {
        const lesson = bookmark.lesson
        if (!lesson) return null

        return (
          <Card key={bookmark.id} className="group relative overflow-hidden">
            <Link href={`/lessons/${lesson.id}`}>
              <div className="relative aspect-video overflow-hidden bg-muted">
                <img
                  src={lesson.thumbnailUrl}
                  alt={lesson.title}
                  className="size-full object-cover transition-transform duration-300 group-hover:scale-105"
                />
                <div className="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 transition-opacity group-hover:opacity-100">
                  <Play className="size-10 fill-white text-white" />
                </div>
              </div>
              <CardContent className="space-y-2 p-3">
                <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
                  <DifficultyBadge level={lesson.level} />
                  <span className="flex items-center gap-0.5">
                    <Clock className="size-3" />
                    {formatDuration(lesson.duration)}
                  </span>
                </div>
                <h3 className="line-clamp-2 text-sm font-medium leading-snug">
                  {lesson.title}
                </h3>
              </CardContent>
            </Link>
            <div className="absolute right-2 top-2">
              <LessonBookmarkButton
                lessonId={lesson.id}
                isBookmarked={true}
                className="bg-white/80 backdrop-blur-sm"
              />
            </div>
          </Card>
        )
      })}
    </div>
  )
}
