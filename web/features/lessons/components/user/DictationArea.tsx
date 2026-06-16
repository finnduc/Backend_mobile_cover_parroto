"use client"

import { useState, useTransition } from "react"
import { Button } from "@/components/ui/button"
import { Field, FieldGroup, FieldLabel, FieldDescription } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Mic } from "lucide-react"
import { cn } from "@/lib/utils"
import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import { postDictationStatus } from "@/features/lessons/services/dictation-status.action"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { WordBlockInput } from "@/features/lessons/components/user/WordBlockInput"
import { ActionRow } from "@/features/lessons/components/user/ActionRow"

function computeWordDiff(correct: string[], input: string[]) {
  const result: { word: string; status: "correct" | "wrong" }[] = []
  for (let i = 0; i < input.length; i++) {
    const c = correct[i]
    const u = input[i]
    if (c && c.toLowerCase() === u.toLowerCase()) {
      result.push({ word: u, status: "correct" })
    } else {
      result.push({ word: u, status: "wrong" })
    }
  }
  return result
}

function WordDiffResult({ diff }: { diff: { word: string; status: "correct" | "wrong" }[] }) {
  return (
    <div className="flex flex-wrap items-center gap-1.5 rounded-lg border bg-muted/30 p-3">
      {diff.map((d, i) => (
        <span
          key={i}
          className={cn(
            "inline-block rounded px-1.5 py-0.5 text-sm font-medium",
            d.status === "correct"
              ? "bg-green-100 text-green-700 dark:bg-green-900/40 dark:text-green-300"
              : "bg-red-100 text-red-700 line-through dark:bg-red-900/40 dark:text-red-300",
          )}
        >
          {d.word}
        </span>
      ))}
    </div>
  )
}

function DictationAreaInner({
  seg,
  currentLine,
  maxLine,
  onSetMaxLine,
  onNext,
  isAuthenticated,
  lessonId,
}: {
  seg: NonNullable<ExerciseControlProps["transcripts"][number]>
  currentLine: number
  maxLine: number
  onSetMaxLine: (line: number) => void
  onNext: () => void
  isAuthenticated: boolean
  lessonId: number
}) {
  const [, startTransition] = useTransition()
  const [input, setInput] = useState("")
  const [revealed, setRevealed] = useState<boolean[]>(
    seg.content.split(/\s+/).map(() => false),
  )
  const [attempts, setAttempts] = useState(0)
  const [answerRevealed, setAnswerRevealed] = useState(false)
  const [wordDiff, setWordDiff] = useState<{ word: string; status: "correct" | "wrong" }[] | null>(null)
  const [feedback, setFeedback] = useState<{
    type: "correct" | "wrong" | "reveal"
    message: string
  } | null>(null)

  const words = seg.content.split(/\s+/).filter((w) => w.length > 0)
  const correctAnswer = seg.content.trim().toLowerCase()
  const userInput = input.trim().toLowerCase()

  const handleCheck = () => {
    if (answerRevealed) {
      onNext()
      return
    }
    if (userInput === correctAnswer) {
      onSetMaxLine(currentLine + 1)
      setFeedback({ type: "correct", message: "Chính xác!" })
      setWordDiff(null)
      if (isAuthenticated) {
        startTransition(async () => {
          const res = await postDictationStatus(seg.id, lessonId)
          if (res.error) {
            setFeedback({ type: "wrong", message: res.error.message })
          }
        })
      }
      onNext()
    } else {
      const next = attempts + 1
      setAttempts(next)
      const diff = computeWordDiff(words, input.trim().split(/\s+/).filter((w) => w.length > 0))
      setWordDiff(diff)
      if (next >= 3) {
        setFeedback({ type: "reveal", message: `Đáp án: ${seg.content}` })
        setAnswerRevealed(true)
      } else {
        setFeedback({ type: "wrong", message: "Sai rồi, thử lại!" })
      }
    }
  }

  const handleToggleWord = (i: number) => {
    setRevealed((prev) => prev.map((v, idx) => (idx === i ? !v : v)))
  }

  const handleRevealAll = () => {
    setRevealed(words.map(() => true))
  }

  return (
    <FieldGroup>
      <Field>
        <FieldLabel htmlFor="dictation-input">
          Gõ những gì bạn nghe được:
        </FieldLabel>
        <div className="relative">
          <Input
            id="dictation-input"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Gõ câu trả lời của bạn ở đây..."
            autoFocus
          />
          <Button
            size="icon-xs"
            variant="ghost"
            disabled
            aria-label="Mic (chưa hỗ trợ)"
            className="absolute inset-e-1 top-1/2 -translate-y-1/2"
          >
            <Mic />
          </Button>
        </div>
        {feedback && (
          <FieldDescription
            className={
              feedback.type === "correct"
                ? "text-green-600"
                : feedback.type === "reveal"
                  ? "text-blue-600"
                  : "text-red-500"
            }
          >
            {feedback.message}
          </FieldDescription>
        )}
        {wordDiff && wordDiff.length > 0 && <WordDiffResult diff={wordDiff} />}
      </Field>

      <WordBlockInput
        text={seg.content}
        activeKey={currentLine}
        revealed={revealed}
        onToggleWord={handleToggleWord}
        onRevealAll={handleRevealAll}
      />

      <div className="flex items-center justify-between gap-2">
        <Button
          size="sm"
          variant="ghost"
          onClick={() => {
            setInput("")
            setAttempts(0)
            setAnswerRevealed(false)
            setFeedback(null)
            setWordDiff(null)
          }}
        >
          Làm lại
        </Button>
        <Button size="sm" onClick={handleCheck}>
          {answerRevealed ? "Tiếp tục" : "Kiểm tra"}
        </Button>
      </div>

      <ActionRow
        onNext={() => {
          onSetMaxLine(currentLine + 1)
          onNext()
        }}
        nextDisabled={currentLine >= maxLine}
      />
    </FieldGroup>
  )
}

export function DictationArea({
  transcripts,
  initialCompletedIds = [],
  lessonId,
  isAuthenticated = false,
}: ExerciseControlProps) {
  const { activeIndex, highlightedIndex, onNext } = usePlayerContext()
  const currentLine = activeIndex >= 0 ? activeIndex : highlightedIndex >= 0 ? highlightedIndex : 0
  const seg = transcripts[currentLine]

  const [maxLine, setMaxLine] = useState(
    initialCompletedIds.length > 0 ? Math.max(...initialCompletedIds) + 1 : 0,
  )

  if (!seg) return null

  const handleSetMaxLine = (line: number) => {
    setMaxLine((prev) => Math.max(prev, line))
  }

  return (
    <DictationAreaInner
      key={currentLine}
      seg={seg}
      currentLine={currentLine}
      maxLine={maxLine}
      onSetMaxLine={handleSetMaxLine}
      onNext={() => onNext?.()}
      isAuthenticated={isAuthenticated}
      lessonId={lessonId}
    />
  )
}
