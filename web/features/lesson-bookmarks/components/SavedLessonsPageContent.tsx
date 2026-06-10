"use client"

import { LessonCard } from "@/features/lessons/components/user/LessonCard"
import type { LessonBookmark } from "@/types/lesson-bookmark.models"
import type { Lesson } from "@/types/lessons.models"

export function SavedLessonsPageContent({
  bookmarks,
}: {
  bookmarks: (LessonBookmark & { lesson?: Lesson })[]
}) {
  if (bookmarks.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <p className="text-lg font-medium">Chưa có bài học nào được lưu</p>
        <p className="text-sm">
          Nhấn vào biểu tượng bookmark để lưu bài học
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
          <LessonCard
            key={bookmark.id}
            lesson={lesson}
            isBookmarked={true}
          />
        )
      })}
    </div>
  )
}
