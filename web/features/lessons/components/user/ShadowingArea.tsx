"use client"

import { LessonProgressBar } from "@/features/lessons/components/user/LessonProgressBar"
import { Button } from "@/components/ui/button"
import { Mic, Square, Play, RefreshCw } from "lucide-react"
import { useState } from "react"

const shadowLines = [
  "In a world where the forest meets the iron age,",
  "a young prince embarks on a journey to find a cure for a curse.",
  "He encounters a fierce young woman raised by wolves,",
  "who fights to protect the forest and its spirits.",
]

export function ShadowingArea() {
  const [currentLine, setCurrentLine] = useState(0)
  const [recording, setRecording] = useState(false)
  const [completed, setCompleted] = useState<number[]>([])

  return (
    <div className="space-y-4">
      <LessonProgressBar completed={completed.length} total={shadowLines.length} />
      <div className="rounded-xl bg-gradient-to-b from-muted/50 to-background p-6 text-center">
        <div className="mb-4 flex items-center justify-center gap-2">
          <div className="flex h-16 w-64 items-center justify-center gap-1 rounded-full bg-muted px-4">
            {Array.from({ length: 40 }).map((_, i) => (
              <div
                key={i}
                className="w-0.5 rounded-full bg-primary"
                style={{
                  height: `${Math.random() * 40 + 8}px`,
                  opacity: recording ? 1 : 0.3,
                }}
              />
            ))}
          </div>
        </div>
        <p className="mb-6 text-lg font-medium">{shadowLines[currentLine]}</p>
        <div className="flex items-center justify-center gap-3">
          <Button
            size="lg"
            variant={recording ? "destructive" : "default"}
            className="size-14 rounded-full"
            onClick={() => setRecording(!recording)}
          >
            {recording ? <Square className="size-6" /> : <Mic className="size-6" />}
          </Button>
          <Button
            size="icon"
            variant="outline"
            onClick={() => {
              if (!completed.includes(currentLine)) {
                setCompleted([...completed, currentLine])
              }
              if (currentLine < shadowLines.length - 1) {
                setCurrentLine(currentLine + 1)
              }
            }}
          >
            <Play className="size-4" />
          </Button>
          <Button size="icon" variant="ghost" onClick={() => { setCurrentLine(0); setCompleted([]) }}>
            <RefreshCw className="size-4" />
          </Button>
        </div>
      </div>
      <div className="space-y-1">
        {shadowLines.map((line, i) => (
          <div
            key={i}
            className={`rounded px-3 py-1.5 text-sm ${
              completed.includes(i) ? "bg-transcript-complete text-muted-foreground line-through" : ""
            } ${i === currentLine ? "bg-transcript-active font-medium" : ""}`}
          >
            {line}
          </div>
        ))}
      </div>
    </div>
  )
}
