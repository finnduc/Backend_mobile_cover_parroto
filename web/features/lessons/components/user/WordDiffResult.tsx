import { cn } from "@/lib/utils"

type WordDiff = {
  word: string
  status: "correct" | "wrong"
  className?: string
}

type WordDiffResultProps = {
  diff: WordDiff[]
  className?: string
}

export function WordDiffResult({ diff, className }: WordDiffResultProps) {
  return (
    <div className={cn("flex flex-wrap items-center gap-1.5 rounded-lg border bg-muted/30 p-3  text-sm font-medium", className)}>
      {diff.map((d, i) => (
        <span
          key={i}
          className={cn(
            "inline-block rounded px-1.5 py-0.5",
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
