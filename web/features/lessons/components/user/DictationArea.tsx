"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import { postDictationStatus } from "@/features/lessons/services/dictation-status.action"
import { Check, Lightbulb, Mic, RotateCcw } from "lucide-react"
import { useEffect, useState, useTransition } from "react"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { PlaybackControls } from "@/features/lessons/components/user/PlaybackControls"

export function DictationArea({
  transcripts,
  initialCompletedIds,
  lessonId,
  isAuthenticated,
}: ExerciseControlProps) {
  const { activeIndex, highlightedIndex, onNext, onTranscriptClick } = usePlayerContext()
  const initialMaxLine = (initialCompletedIds ?? []).length > 0
    ? Math.max(...(initialCompletedIds ?? [])) + 1
    : 0
  const [maxLine, setMaxLine] = useState(initialMaxLine)
  const [inputs, setInputs] = useState<string[]>(transcripts.map(() => ""))
  const [showHints, setShowHints] = useState(false)
  const [, startTransition] = useTransition()
  const [attempts, setAttempts] = useState<Record<number, number>>({})
  const [answerRevealed, setAnswerRevealed] = useState<Set<number>>(new Set())
  const [feedback, setFeedback] = useState<{ line: number; type: "correct" | "wrong" | "reveal"; message: string } | null>(null)

  const currentLine = activeIndex >= 0 ? activeIndex : highlightedIndex >= 0 ? highlightedIndex : 0

  useEffect(() => {
    if (activeIndex >= 0) {
      setMaxLine((prev) => Math.max(prev, activeIndex + 1))
    }
  }, [activeIndex])

  return (
    <div className="space-y-4">
      <PlaybackControls totalLines={transcripts.length}>
        <Button size="lg" variant="default" className="size-12 rounded-full" disabled>
          <Mic className="size-5" />
        </Button>
      </PlaybackControls>

      <div className="space-y-2">
        {transcripts.map((seg, i) => {
          const isComplete = i < maxLine
          const isCurrent = i === currentLine
          const isSaved = (initialCompletedIds ?? []).includes(i)
          return (
            <div
              key={seg.id}
              onClick={() => {
                if (isComplete && !isCurrent) {
                  onTranscriptClick?.(i)
                }
              }}
              className={`rounded-lg p-3 transition-colors border ${isCurrent ? "border-2 border-primary/20 bg-transcript-active" : ""
                } ${isSaved ? "bg-green-100" : ""} ${isComplete && !isSaved && !isCurrent ? "bg-amber-50" : ""} ${isComplete && !isCurrent ? "cursor-pointer hover:bg-transcript-active/50" : ""
                }`}
            >
              {isComplete && !isCurrent ? (
                <div className="space-y-1">
                  <div className="flex items-center gap-2">
                    {isSaved ? (
                      <Check className="size-4 shrink-0 text-green-600" />
                    ) : (
                      <span className="size-4 shrink-0 rounded-full border border-amber-400" />
                    )}
                    <p className="text-sm">{seg.content}</p>
                  </div>
                  {seg.phonetic && (
                    <p className="text-xs text-muted-foreground pl-6">/{seg.phonetic}/</p>
                  )}
                </div>
              ) : isCurrent ? (
                <div className="space-y-2">
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">
                      {Math.floor(seg.startTimestamp / 60)}:{(seg.startTimestamp % 60).toString().padStart(2, "0")}
                      {" - "}
                      {Math.floor(seg.endTimestamp / 60)}:{(seg.endTimestamp % 60).toString().padStart(2, "0")}
                    </span>
                  </div>
                  {showHints && (
                    <div className="text-xs text-muted-foreground space-y-0.5">
                      {seg.phonetic && <p>/{seg.phonetic}/</p>}
                      <p>{seg.vietnamese || seg.content.slice(0, 20) + "..."}</p>
                    </div>
                  )}
                  <Input
                    value={inputs[i]}
                    onChange={(e) => {
                      const next = [...inputs]
                      next[i] = e.target.value
                      setInputs(next)
                    }}
                    placeholder="Gõ những gì bạn nghe được..."
                    className="w-full"
                    autoFocus
                  />
                  <div className="flex gap-2">
                    <Button
                      size="xs"
                      onClick={() => {
                        const seg = transcripts[currentLine]
                        const userInput = (inputs[currentLine] ?? "").trim().toLowerCase()
                        const correctAnswer = seg?.content.trim().toLowerCase()

                        if (!seg) return

                        // If answer already revealed, just advance
                        if (answerRevealed.has(currentLine)) {
                          onNext?.()
                          return
                        }

                        if (userInput === correctAnswer) {
                          // Correct answer
                          const newMaxLine = Math.max(maxLine, currentLine + 1)
                          setMaxLine(newMaxLine)
                          setFeedback({ line: currentLine, type: "correct", message: "Chính xác!" })
                          if (isAuthenticated) {
                            startTransition(async () => {
                              const res = await postDictationStatus(seg.id, lessonId!)
                              if (res.error) {
                                setFeedback({ line: currentLine, type: "wrong", message: res.error.message })
                              }
                            })
                          }
                          onNext?.()
                        } else {
                          // Wrong answer
                          const currentAttempts = (attempts[currentLine] ?? 0) + 1
                          setAttempts((prev) => ({ ...prev, [currentLine]: currentAttempts }))
                          if (currentAttempts >= 3) {
                            setFeedback({ line: currentLine, type: "reveal", message: `Đáp án: ${seg.content}` })
                            setAnswerRevealed((prev) => new Set(prev).add(currentLine))
                          } else {
                            setFeedback({ line: currentLine, type: "wrong", message: "Sai rồi, thử lại!" })
                          }
                        }
                      }}
                    >
                      {answerRevealed.has(currentLine) ? "Tiếp tục" : "Kiểm tra"}
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setShowHints(!showHints)}>
                      <Lightbulb className="mr-1 size-3" />
                      Gợi ý
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setInputs(transcripts.map(() => ""))}>
                      <RotateCcw className="mr-1 size-3" />
                      Làm lại
                    </Button>
                  </div>
                  {feedback && feedback.line === currentLine && (
                    <p className={`text-sm ${feedback.type === "correct" ? "text-green-600" :
                        feedback.type === "reveal" ? "text-blue-600" :
                          "text-red-500"
                      }`}>
                      {feedback.message}
                    </p>
                  )}
                </div>
              ) : (
                <p className="text-sm text-muted-foreground line-through select-none">
                  {seg.content.split(" ").map((word, i) => (
                    <span
                      key={i}
                      className="mr-1 inline-block h-3 rounded bg-muted align-middle"
                      style={{ width: `${Math.max(word.length * 7, 16)}px` }}
                    />
                  ))}
                </p>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}
