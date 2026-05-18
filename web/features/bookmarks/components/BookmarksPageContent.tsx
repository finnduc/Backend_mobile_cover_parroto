"use client"

import { BookmarkButton } from "@/features/bookmarks/components/BookmarkButton"
import { Card, CardContent } from "@/components/ui/card"
import { DifficultyBadge } from "@/components/common/DifficultyBadge"
import type { Bookmark } from "@/types/book-mark.models"
import { Bookmark as BookmarkIcon, Play, Clock } from "lucide-react"
import Link from "next/link"
import { ROUTES } from "@/lib/routes"

function formatDuration(seconds: number): string {
  const m = Math.floor(seconds / 60)
  const s = Math.floor(seconds % 60)
  return `${m.toString().padStart(2, "0")}:${s.toString().padStart(2, "0")}`
}

export function BookmarksPageContent({ bookmarks }: { bookmarks: Bookmark[] }) {
  if (bookmarks.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <BookmarkIcon className="mb-4 size-12" />
        <p className="text-lg font-medium">Chưa có bài học nào được lưu</p>
        <p className="text-sm">Hãy lưu các bài học bạn yêu thích để xem sau</p>
      </div>
    )
  }

  return (
    <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
      {bookmarks.map((bookmark) => (
        <Card key={bookmark.lessionId} className="group cursor-pointer overflow-hidden">
          <div className="relative aspect-video overflow-hidden bg-muted">
            <img
              src={bookmark.lesson.thumbnailUrl}
              alt={bookmark.lesson.title}
              className="size-full object-cover transition-transform duration-300 group-hover:scale-105"
            />
            <Link
              href={ROUTES.USER.LESSONS.DICTATION(bookmark.lessionId)}
              className="absolute inset-0 flex items-center justify-center bg-black/30 opacity-0 transition-opacity group-hover:opacity-100"
            >
              <Play className="size-10 fill-white text-white" />
            </Link>
            <BookmarkButton
              lessonId={bookmark.lessionId}
              isBookmarked
              className="absolute top-2 right-2 bg-black/40 hover:bg-black/60 text-white"
            />
          </div>
          <CardContent className="space-y-2 p-3">
            <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
              <DifficultyBadge level={bookmark.lesson.level} />
              <span className="flex items-center gap-0.5">
                <Clock className="size-3" />
                {formatDuration(bookmark.lesson.duration)}
              </span>
            </div>
            <h3 className="line-clamp-2 text-sm font-medium leading-snug">{bookmark.lesson.title}</h3>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
