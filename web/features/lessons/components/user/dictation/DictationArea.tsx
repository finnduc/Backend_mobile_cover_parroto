"use client"

import { Button } from "@/components/ui/button"
import { Field, FieldDescription, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { ActionRow } from "@/features/lessons/components/user/ActionRow"
import { WordDiffResult } from "@/features/lessons/components/user/WordDiffResult"
import { WordBlockInput } from "@/features/lessons/components/user/dictation/WordBlockInput"
import { usePlayerContext } from "@/features/lessons/context/player-context"
import { postDictationStatus } from "@/features/lessons/services/dictation-status.action"
import type { ExerciseControlProps } from "@/features/lessons/types/exercise.types"
import { computeWordDiff } from "@/features/lessons/utils/word-diff"
import { cn, normalizeSentence, splitWords } from "@/lib/utils"
import { Transcript } from "@/types/lessons.models"
import { useState, useTransition } from "react"

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
