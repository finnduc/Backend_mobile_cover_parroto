"use client"

import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { RefreshCw, Lightbulb } from "lucide-react"
import { useState } from "react"

interface DictationLine {
  index: number
  text: string
  hint: string
}

const mockLines: DictationLine[] = [
  { index: 0, text: "In a world where the forest meets the iron age,", hint: "In a world..." },
  { index: 1, text: "a young prince embarks on a journey to find a cure for a curse.", hint: "a young prince..." },
  { index: 2, text: "He encounters a fierce young woman raised by wolves,", hint: "He encounters..." },
  { index: 3, text: "who fights to protect the forest and its spirits.", hint: "who fights..." },
  { index: 4, text: "Together they must stop the forces of destruction", hint: "Together they must..." },
  { index: 5, text: "before the balance of nature is lost forever.", hint: "before the balance..." },
]

export function DictationArea() {
  const [currentLine, setCurrentLine] = useState(0)
  const [inputs, setInputs] = useState<string[]>(mockLines.map(() => ""))
  const [showHints, setShowHints] = useState(false)
  const [showResult, setShowResult] = useState(false)

  return (
    <div className="space-y-4">
      <LessonProgressBar completed={showResult ? currentLine + 1 : currentLine} total={mockLines.length} />
      <div className="space-y-2">
        {mockLines.map((line, i) => {
          const isComplete = i < currentLine
          const isCurrent = i === currentLine
          return (
            <div
              key={i}
              className={`rounded-lg p-3 transition-colors ${
                isCurrent ? "border-2 border-primary/20 bg-transcript-active" : ""
              } ${isComplete ? "bg-transcript-complete" : ""}`}
            >
              {isComplete && showResult ? (
                <p className="text-sm text-foreground">{line.text}</p>
              ) : isCurrent ? (
                <div className="space-y-2">
                  {showHints && (
                    <p className="text-xs text-muted-foreground">{line.hint}</p>
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
                        setShowResult(true)
                        if (currentLine < mockLines.length - 1) {
                          setCurrentLine(currentLine + 1)
                        }
                      }}
                    >
                      Kiểm tra
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setShowHints(!showHints)}>
                      <Lightbulb className="mr-1 size-3" />
                      Gợi ý
                    </Button>
                    <Button size="xs" variant="ghost" onClick={() => setInputs(mockLines.map(() => ""))}>
                      <RefreshCw className="mr-1 size-3" />
                      Làm lại
                    </Button>
                  </div>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground line-through">
                  {line.text.replace(/[a-zA-Z0-9]/g, "▸")}
                </p>
              )}
            </div>
          )
        })}
      </div>
    </div>
  )
}
