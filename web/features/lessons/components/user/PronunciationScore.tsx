"use client"

import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Progress } from "@/components/ui/progress"
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import type { PronunciationAttempt } from "@/types/pronunciation.models"

function scoreGrade(score: number): { label: string; className: string } {
  if (score >= 80) return { label: "A", className: "bg-green-500" }
  if (score >= 60) return { label: "B", className: "bg-amber-500" }
  return { label: "C", className: "bg-red-500" }
}

export function PronunciationScore({ result }: { result: PronunciationAttempt }) {
  const grade = scoreGrade(result.overallScore)

  return (
    <Card>
      <CardContent className="space-y-4 pt-6">
        <div className="flex items-center justify-between">
          <span className="text-sm font-medium">Pronunciation Score</span>
          <Badge className={grade.className + " text-white"}>{grade.label} ({result.overallScore.toFixed(0)})</Badge>
        </div>

        <div className="space-y-2">
          {(["accuracy", "fluency", "completeness", "prosody"] as const).map((key) => {
            const val = result.scores[key]
            const barColor = val >= 80 ? "bg-green-500" : val >= 60 ? "bg-amber-500" : "bg-red-500"
            return (
              <div key={key} className="flex items-center gap-3">
                <span className="w-24 text-xs capitalize text-muted-foreground">{key}</span>
                <div className="flex-1">
                  <Progress value={val} />
                </div>
                <span className="w-10 text-right text-xs">{val.toFixed(0)}</span>
              </div>
            )
          })}
        </div>

        {result.feedback && (
          <p className="text-xs text-muted-foreground">{result.feedback}</p>
        )}

        {result.words && result.words.length > 0 && (
          <Accordion type="single" collapsible>
            <AccordionItem value="words" className="border-none">
              <AccordionTrigger className="text-xs py-0">Word Details</AccordionTrigger>
              <AccordionContent>
                <div className="mt-2 space-y-1">
                  {result.words.map((w, i) => (
                    <div key={i} className="flex items-center gap-2 text-xs">
                      <Badge variant="outline" className="text-xs">{w.score.toFixed(0)}</Badge>
                      <span className="font-mono">{w.word}</span>
                      {w.feedback && <span className="text-muted-foreground">- {w.feedback}</span>}
                    </div>
                  ))}
                </div>
              </AccordionContent>
            </AccordionItem>
          </Accordion>
        )}
      </CardContent>
    </Card>
  )
}
