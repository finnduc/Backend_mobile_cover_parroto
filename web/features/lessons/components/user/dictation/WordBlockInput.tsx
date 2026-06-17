"use client"

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { Eye, EyeOff } from "lucide-react"

type Props = {
  words: string[]
  activeKey: string | number
  revealed: boolean[]
  onToggleWord: (index: number) => void
  onRevealAll: () => void
  onHideAll: () => void
  answerRevealed: boolean
}

function blockWidth(word: string): number {
  return Math.max(word.length * 8, 32)
}

export function WordBlockInput({
  words,
  activeKey,
  revealed,
  onToggleWord,
  onRevealAll,
  onHideAll,
  answerRevealed,
}: Props) {

  return (
    <div className="flex flex-col gap-3">
      <div className="flex flex-wrap items-center gap-2">
        {words.map((w, i) => {
          const isRevealed = revealed[i]
          return (
            <div
              key={`${activeKey}-${i}`}
              className={cn(
                "inline-flex items-center gap-1 rounded-md border px-2 py-1",
                isRevealed
                  ? "border-green-300 bg-green-100 text-green-800 dark:bg-green-900/30 dark:text-green-200"
                  : "border-border bg-muted text-transparent",
              )}
              style={!isRevealed ? { minWidth: blockWidth(w) } : undefined}
            >
              <span className={cn("text-sm", !isRevealed && "select-none")}>
                {isRevealed ? w : "****"}
              </span>
              <Button
                size="icon-xs"
                variant="ghost"
                aria-label={isRevealed ? `Ẩn từ ${w}` : `Hiện từ ${w}`}
                onClick={() => onToggleWord(i)}
              >
                {isRevealed ? <EyeOff /> : <Eye />}
              </Button>
            </div>
          )
        })}
      </div>
      {answerRevealed ? <Button
        variant="outline"
        className="self-start bg-orange-500 text-white hover:bg-orange-600 hover:text-white dark:bg-orange-500 dark:hover:bg-orange-600"
        onClick={onHideAll}
      >
        Che tất cả từ
      </Button> : <Button
        variant="outline"
        className="self-start bg-orange-500 text-white hover:bg-orange-600 hover:text-white dark:bg-orange-500 dark:hover:bg-orange-600"
        onClick={onRevealAll}
      >
        Hiện tất cả từ
      </Button>}
    </div>
  )
}
