"use client"

import { cn } from "@/lib/utils"

export function SeekBar({
  current = 0,
  total = 100,
  className,
}: {
  current?: number
  total?: number
  className?: string
}) {
  const pct = total > 0 ? (current / total) * 100 : 0
  return (
    <div className={cn("relative h-1 w-full rounded-full bg-muted", className)}>
      <div
        className="absolute inset-y-0 left-0 rounded-full bg-primary transition-all duration-200"
        style={{ width: `${pct}%` }}
      />
      <div
        className="absolute top-1/2 size-3 -translate-y-1/2 rounded-full bg-primary shadow-sm transition-all duration-200"
        style={{ left: `${pct}%`, marginLeft: -6 }}
      />
    </div>
  )
}
