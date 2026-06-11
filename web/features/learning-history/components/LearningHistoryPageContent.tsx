"use client"

import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { CheckCircle, Clock, Mic, Headphones } from "lucide-react"
import type { LearningHistory, LearningHistorySummary } from "@/types/learning-history.models"
import { ROUTES } from "@/lib/routes"
import Link from "next/link"

export function LearningHistoryPageContent({
  finished,
  unfinished,
  summary,
  shadowingPercent,
  dictationPercent,
}: {
  finished: LearningHistory[]
  unfinished: LearningHistory[]
  summary: LearningHistorySummary
  shadowingPercent: number
  dictationPercent: number
}) {
  return (
    <div className="space-y-6">
      <div className="grid grid-cols-2 gap-4 md:grid-cols-4">
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <CheckCircle className="size-6 text-green-500" />
            <div>
              <p className="text-2xl font-bold">{summary.completedCount}</p>
              <p className="text-xs text-muted-foreground">Đã hoàn thành</p>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <Clock className="size-6 text-amber-500" />
            <div>
              <p className="text-2xl font-bold">{summary.unfinishedCount}</p>
              <p className="text-xs text-muted-foreground">Đang học</p>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <Mic className="size-6 text-blue-500" />
            <div>
              <p className="text-2xl font-bold">{shadowingPercent}%</p>
              <p className="text-xs text-muted-foreground">Câu đã shadow</p>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <Headphones className="size-6 text-purple-500" />
            <div>
              <p className="text-2xl font-bold">{dictationPercent}%</p>
              <p className="text-xs text-muted-foreground">Câu đã dictate</p>
            </div>
          </CardContent>
        </Card>
      </div>

      <section>
        <h2 className="mb-3 text-lg font-semibold">Đã hoàn thành ({finished.length})</h2>
        {finished.length === 0 ? (
          <p className="text-sm text-muted-foreground">Chưa có bài học nào hoàn thành</p>
        ) : (
          <div className="space-y-2">
            {finished.map((h) => (
              <Card key={h.id}>
                <CardContent className="flex items-center gap-3 py-3">
                  <CheckCircle className="size-4 text-green-500" />
                  <Link
                    href={ROUTES.USER.LESSONS.DICTATION(h.lessonId)}
                    className="text-sm font-medium hover:underline"
                  >
                    Lesson #{h.lessonId}
                  </Link>
                  <div className="ml-auto flex items-center gap-1.5">
                    {h.completedDictation && (
                      <Badge className="gap-1 bg-green-500 text-white hover:bg-green-600 text-xs">
                        Dictation
                      </Badge>
                    )}
                    {h.completedPronunciation && (
                      <Badge className="gap-1 bg-blue-500 text-white hover:bg-blue-600 text-xs">
                        Pronunciation
                      </Badge>
                    )}
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </section>

      <section>
        <h2 className="mb-3 text-lg font-semibold">Đang học ({unfinished.length})</h2>
        {unfinished.length === 0 ? (
          <p className="text-sm text-muted-foreground">Chưa có bài học đang học</p>
        ) : (
          <div className="space-y-2">
            {unfinished.map((h) => (
              <Card key={h.id}>
                <CardContent className="flex items-center gap-3 py-3">
                  <Clock className="size-4 text-amber-500" />
                  <Link
                    href={ROUTES.USER.LESSONS.DICTATION(h.lessonId)}
                    className="text-sm font-medium hover:underline"
                  >
                    Lesson #{h.lessonId}
                  </Link>
                  <div className="ml-auto flex items-center gap-1.5">
                    {h.completedDictation && (
                      <Badge className="gap-1 bg-green-500 text-white hover:bg-green-600 text-xs">
                        Dictation
                      </Badge>
                    )}
                    {h.completedPronunciation && (
                      <Badge className="gap-1 bg-blue-500 text-white hover:bg-blue-600 text-xs">
                        Pronunciation
                      </Badge>
                    )}
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </section>
    </div>
  )
}
