"use client"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Badge } from "@/components/ui/badge"
import { CheckCircle, PlayCircle } from "lucide-react"
import type { Lesson } from "@/types/lessons.models"
import type { LearningHistory } from "@/types/learning-history.models"

export function LessonProgressDialog({
  lesson,
  history,
  open,
  onOpenChange,
}: {
  lesson: Lesson
  history: LearningHistory | null
  open: boolean
  onOpenChange: (open: boolean) => void
}) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="text-base leading-snug">
            Tiến độ bài học
          </DialogTitle>
        </DialogHeader>

        <div className="space-y-4">
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

          <p className="line-clamp-2 text-sm font-medium">{lesson.title}</p>

          {history ? (
            <div className="space-y-3">
              <div className="flex flex-wrap items-center gap-2">
                {history.completedDictation && (
                  <Badge className="gap-1 bg-green-500 text-white hover:bg-green-600">
                    <CheckCircle className="size-3" />
                    Dictation hoàn thành
                  </Badge>
                )}
                {history.completedPronunciation && (
                  <Badge className="gap-1 bg-blue-500 text-white hover:bg-blue-600">
                    <CheckCircle className="size-3" />
                    Pronunciation hoàn thành
                  </Badge>
                )}
                {!history.completedDictation && !history.completedPronunciation && (
                  <Badge variant="secondary" className="gap-1">
                    <PlayCircle className="size-3" />
                    Đang học
                  </Badge>
                )}
              </div>
            </div>
          ) : (
            <div className="rounded-lg border border-dashed p-4 text-center text-sm text-muted-foreground">
              Bạn chưa bắt đầu học bài này. Hãy thử <strong>Dictation</strong> hoac{" "}
              <strong>Shadowing</strong>!
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  )
}
