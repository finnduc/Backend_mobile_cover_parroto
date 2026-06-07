"use client"

import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { CheckCircle, Clock } from "lucide-react"
import type { LearningHistory, LearningHistorySummary } from "@/types/learning-history.models"
import { ROUTES } from "@/lib/routes"
import Link from "next/link"

export function LearningHistoryPageContent({
  finished,
  unfinished,
  summary,
}: {
  finished: LearningHistory[]
  unfinished: LearningHistory[]
  summary: LearningHistorySummary
}) {
  return (
    <div className="space-y-6">
      <div className="grid grid-cols-2 gap-4">
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <CheckCircle className="size-6 text-green-500" />
            <div>
              <p className="text-2xl font-bold">{summary.completedCount}</p>
              <p className="text-xs text-muted-foreground">Da hoan thanh</p>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="flex items-center gap-3 pt-6">
            <Clock className="size-6 text-amber-500" />
            <div>
              <p className="text-2xl font-bold">{summary.unfinishedCount}</p>
              <p className="text-xs text-muted-foreground">Dang hoc</p>
            </div>
          </CardContent>
        </Card>
      </div>

      <section>
        <h2 className="mb-3 text-lg font-semibold">Da hoan thanh ({finished.length})</h2>
        {finished.length === 0 ? (
          <p className="text-sm text-muted-foreground">Chua co bai hoc nao hoan thanh</p>
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
        <h2 className="mb-3 text-lg font-semibold">Dang hoc ({unfinished.length})</h2>
        {unfinished.length === 0 ? (
          <p className="text-sm text-muted-foreground">Chua co bai hoc dang hoc</p>
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
