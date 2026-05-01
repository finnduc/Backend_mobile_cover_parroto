import type { ReactNode } from "react"
import { VideoPlayerPlaceholder } from "@/components/common/VideoPlayerPlaceholder"
import { LessonToolbar } from "@/features/lessons/components/user/LessonToolbar"

export function LessonLayout({
  transcript,
  exercise,
  duration,
}: {
  title?: string
  transcript: ReactNode
  exercise: ReactNode
  duration?: number
}) {
  return (
    <div className="flex h-full gap-6">
      <div className="flex flex-1 flex-col gap-4">
        <VideoPlayerPlaceholder duration={duration} />
        <LessonToolbar />
        <div className="flex-1">{exercise}</div>
      </div>
      <aside className="hidden w-80 shrink-0 lg:block">
        <div className="sticky top-6 space-y-3">
          <h2 className="text-sm font-semibold text-muted-foreground">Transcript</h2>
          <div className="max-h-[calc(100vh-12rem)] space-y-1 overflow-y-auto rounded-xl border p-2">
            {transcript}
          </div>
        </div>
      </aside>
    </div>
  )
}
