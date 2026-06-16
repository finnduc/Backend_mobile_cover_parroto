"use client"

import { Button } from "@/components/ui/button"
import { ArrowRight } from "lucide-react"

type Props = {
  onNext: () => void
  nextDisabled?: boolean
  nextLabel?: string
}

export function ActionRow({ onNext, nextDisabled, nextLabel = "Tiếp theo" }: Props) {
  return (
    <div className="flex flex-col gap-2">
      <p className="text-xs text-muted-foreground">
        Các từ được tiết lộ sẽ bị tính là lỗi và ảnh hưởng đến điểm số của bạn.
      </p>
      <Button
        size="lg"
        onClick={onNext}
        disabled={nextDisabled}
        className="w-full"
      >
        {nextLabel}
        <ArrowRight data-icon="inline-end" />
      </Button>
    </div>
  )
}
