"use client"

import { Button } from "@/components/ui/button"
import { Eye, EyeOff } from "lucide-react"
import { cn } from "@/lib/utils"

type Props = {
  text: string
  activeKey: string | number
  revealed: boolean[]
  onToggleWord: (index: number) => void
  onRevealAll: () => void
}

function blockWidth(word: string): number {
  return Math.max(word.length * 8, 32)
}

export function WordBlockInput({
  text,
  activeKey,
  revealed,
  onToggleWord,
  onRevealAll,
}: Props) {
  const words = text.split(/\s+/).filter((w) => w.length > 0)

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
      <Button
        variant="outline"
        className="self-start bg-orange-500 text-white hover:bg-orange-600 hover:text-white dark:bg-orange-500 dark:hover:bg-orange-600"
        onClick={onRevealAll}
      >
        Hiện tất cả từ
      </Button>
    </div>
  )
}
