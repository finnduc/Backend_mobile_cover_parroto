"use client"

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

const speeds = ["0.5", "0.75", "1", "1.25", "1.5", "2"] as const

export function PlaybackSpeed({
  value,
  onChange,
}: {
  value: number
  onChange: (speed: number) => void
}) {
  return (
    <Select
      value={String(value)}
      onValueChange={(v) => onChange(Number(v))}
    >
      <SelectTrigger className="w-20">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        {speeds.map((s) => (
          <SelectItem key={s} value={s}>
            {s}x
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  )
}
