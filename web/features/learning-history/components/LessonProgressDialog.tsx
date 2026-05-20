"use client"

import { useEffect, useState, useTransition } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Badge } from "@/components/ui/badge"
import { CheckCircle, Clock, PlayCircle } from "lucide-react"
import type { Lesson } from "@/types/lessons.models"
import type { LearningHistory } from "@/types/learning-history.models"
import { getHistoryByLesson } from "@/features/learning-history/services/learning-history.action"

function formatDuration(seconds: number): string {
  const h = Math.floor(seconds / 3600)
  const m = Math.floor((seconds % 3600) / 60)
  const s = Math.floor(seconds % 60)
  if (h > 0) return `${h}g ${m}p ${s}s`
  if (m > 0) return `${m}p ${s}s`
  return `${s}s`
}

export function LessonProgressDialog({
  lesson,
  open,
  onOpenChange,
}: {
  lesson: Lesson
  open: boolean
  onOpenChange: (open: boolean) => void
}) {
  const [history, setHistory] = useState<LearningHistory | null>(null)
  const [isPending, startTransition] = useTransition()

  useEffect(() => {
    if (!open) return
    startTransition(async () => {
      const res = await getHistoryByLesson(lesson.id)
      setHistory(res.data)
    })
  }, [open, lesson.id])

  const watchedPct = lesson.duration > 0 && history
    ? Math.min(100, (history.durationWatched / lesson.duration) * 100)
    : 0

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="text-base leading-snug">
            Tiến độ bài học
          </DialogTitle>
        </DialogHeader>

        <div className="space-y-4">
          {/* Sample progress image */}
          <div className="overflow-hidden rounded-xl border bg-muted">
            <img
              src="/process.jpg"
              alt="Tiến độ học tập"
              className="w-full object-cover"
              onError={(e) => {
                const el = e.currentTarget
                el.style.display = "none"
                el.nextElementSibling?.classList.remove("hidden")
              }}
            />
            <div className="hidden flex-col items-center justify-center gap-2 py-10 text-muted-foreground">
              <PlayCircle className="size-10 opacity-30" />
              <span className="text-sm">Chưa có ảnh tiến độ</span>
            </div>
          </div>

          {/* Lesson title */}
          <p className="line-clamp-2 text-sm font-medium">{lesson.title}</p>

          {isPending ? (
            <div className="space-y-2">
              <div className="h-3 w-full animate-pulse rounded-full bg-muted" />
              <div className="h-3 w-2/3 animate-pulse rounded-full bg-muted" />
            </div>
          ) : history ? (
            <div className="space-y-3">
              {/* Status badge */}
              <div className="flex items-center gap-2">
                {history.completed ? (
                  <Badge className="gap-1 bg-green-500 text-white hover:bg-green-600">
                    <CheckCircle className="size-3" />
                    Đã hoàn thành
                  </Badge>
                ) : (
                  <Badge variant="secondary" className="gap-1">
                    <PlayCircle className="size-3" />
                    Đang học
                  </Badge>
                )}
              </div>

              {/* Progress bar */}
              <div className="space-y-1.5">
                <div className="flex items-center justify-between text-xs text-muted-foreground">
                  <span>Thời gian đã học</span>
                  <span>
                    {formatDuration(history.durationWatched)} / {formatDuration(lesson.duration)}
                  </span>
                </div>
                <div className="h-2 w-full overflow-hidden rounded-full bg-muted">
                  <div
                    className="h-full rounded-full bg-primary transition-all duration-500"
                    style={{ width: `${watchedPct}%` }}
                  />
                </div>
                <p className="text-right text-xs text-muted-foreground">
                  {watchedPct.toFixed(0)}%
                </p>
              </div>

              {/* Stats */}
              <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
                <Clock className="size-3.5" />
                <span>Bắt đầu lần đầu: {new Date(history.createdAt).toLocaleDateString("vi-VN")}</span>
              </div>
            </div>
          ) : (
            <div className="rounded-lg border border-dashed p-4 text-center text-sm text-muted-foreground">
              Bạn chưa bắt đầu học bài này. Hãy thử <strong>Dictation</strong> hoặc <strong>Shadowing</strong>!
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  )
}
