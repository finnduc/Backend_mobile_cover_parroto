"use client"

import { cn } from "@/lib/utils"

type Word = { word: string; phonetic?: string }

function splitWords(s: string): string[] {
  return s.split(/\s+/).filter((w) => w.length > 0)
}

function alignPhonetic(content: string, phonetic?: string): Word[] {
  const words = splitWords(content)
  if (!phonetic) return words.map((w) => ({ word: w }))
  const tokens = splitWords(phonetic.replace(/^\/|\/$/g, ""))
  if (tokens.length !== words.length) {
    return words.map((w) => ({ word: w, phonetic: tokens[0] }))
  }
  return words.map((w, i) => ({ word: w, phonetic: tokens[i] }))
}

export function SpacedTextDisplay({
  content,
  phonetic,
  className,
}: {
  content: string
  phonetic?: string
  className?: string
}) {
  const words = alignPhonetic(content, phonetic)
  return (
    <div className={cn("flex flex-col items-center gap-3", className)}>
      <div className="flex flex-wrap items-baseline justify-center gap-x-4 gap-y-2">
        {words.map((w, i) => (
          <div key={i} className="flex flex-col items-center">
            <span className="text-lg font-medium underline decoration-muted-foreground/40 underline-offset-4">
              {w.word}
            </span>
            {w.phonetic && (
              <span className="text-xs text-primary">/{w.phonetic}/</span>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}
