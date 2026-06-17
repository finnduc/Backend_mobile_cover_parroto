"use client"

import { Button } from "@/components/ui/button"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { ActionRow } from "@/features/lessons/components/user/ActionRow"
import { WordBlockInput } from "@/features/lessons/components/user/dictation/WordBlockInput"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { postDictationStatus } from "@/features/lessons/services/dictation-status.action"
import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import { cn, normalizeSentence, normalizeWord, splitWords } from "@/lib/utils"
import { Transcript } from "@/types/lessons.models"
import { useState, useTransition } from "react"


function computeWordDiff(correct: string[], input: string[]) {
  const result: { word: string; status: "correct" | "wrong" }[] = []

  for (let i = 0; i < input.length; i++) {
    const c = normalizeWord(correct[i] ?? "")
    const u = normalizeWord(input[i] ?? "")

    result.push({
      word: input[i],
      status: c === u ? "correct" : "wrong",
    })
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
  seg: Transcript
  currentLine: number
  maxLine: number
  onSetMaxLine: (line: number) => void
  onNext: () => void
  isAuthenticated: boolean
  lessonId: number
}) {
  const [, startTransition] = useTransition()

  const words = splitWords(seg.content)

  const [input, setInput] = useState("")

  const [revealed, setRevealed] = useState<boolean[]>(
    () => words.map(() => false),
  )

  const answerRevealed =
    revealed.length === words.length &&
    revealed.every(Boolean)

  const [wordDiff, setWordDiff] = useState<
    { word: string; status: "correct" | "wrong" }[] | null
  >(null)

  const [feedback, setFeedback] = useState<{
    type: "correct" | "wrong" | "reveal"
    message: string
  } | null>(null)


  const correctAnswer = normalizeSentence(seg.content)
  const userInput = normalizeSentence(input)

  const handleCheck = () => {
    if (answerRevealed) {
      return
    }

    if (userInput === correctAnswer) {
      onSetMaxLine(currentLine + 1)

      setFeedback({
        type: "correct",
        message: "Chính xác!",
      })

      setWordDiff(null)

      setRevealed(words.map(() => true))

      if (isAuthenticated) {
        startTransition(async () => {
          const res = await postDictationStatus(seg.id, lessonId)

          if (res.error) {
            setFeedback({
              type: "wrong",
              message: res.error.message,
            })
          }
        })
      }
    } else {
      const diff = computeWordDiff(
        words,
        input
          .trim()
          .split(/\s+/)
          .filter(Boolean),
      )

      setWordDiff(diff)

      setFeedback({
        type: "wrong",
        message: "Sai rồi, thử lại!",
      })
    }
  }

  const handleToggleWord = (i: number) => {
    setRevealed((prev) =>
      prev.map((v, idx) => (idx === i ? !v : v)),
    )
  }

  const handleRevealAll = () => {
    setRevealed(words.map(() => true))
  }

  const handleHideAll = () => {
    setRevealed(words.map(() => false))
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

        {wordDiff && wordDiff.length > 0 && (
          <WordDiffResult diff={wordDiff} />
        )}
      </Field>

      <WordBlockInput
        words={words}
        activeKey={currentLine}
        revealed={revealed}
        onToggleWord={handleToggleWord}
        onRevealAll={handleRevealAll}
        onHideAll={handleHideAll}
        answerRevealed={answerRevealed}
      />

      <div className="flex items-center justify-between gap-2">
        <Button
          size="sm"
          variant="ghost"
          onClick={() => {
            setInput("")
            setFeedback(null)
            setWordDiff(null)
            setRevealed(words.map(() => false))
          }}
        >
          Làm lại
        </Button>

        <Button size="sm" onClick={handleCheck}>
          Kiểm tra
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
