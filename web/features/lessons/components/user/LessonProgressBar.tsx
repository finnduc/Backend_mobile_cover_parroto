import { SeekBar } from "@/components/common/SeekBar"

export function LessonProgressBar({
  completed,
  total,
}: {
  completed: number
  total: number
}) {
  return (
    <div className="space-y-1">
      <div className="flex items-center justify-between text-xs text-muted-foreground">
        <span>Tiến độ</span>
        <span>
          {completed}/{total} dòng
        </span>
      </div>
      <SeekBar current={completed} total={total} />
    </div>
  )
}
