"use client"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { BookOpenText, Mic } from "lucide-react"
import Link from "next/link"
import type { Lesson } from "@/types/lessons.models"
import { ROUTES } from "@/lib/routes"

export function LessonActionModal({
  lesson,
  open,
  onOpenChange,
}: {
  lesson: Lesson
  open: boolean
  onOpenChange: (open: boolean) => void
}) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle className="text-center text-base leading-snug">
            {lesson.title}
          </DialogTitle>
        </DialogHeader>
        <div className="flex flex-col gap-3 pt-2">
          <Button asChild size="lg" className="w-full justify-start gap-3">
            <Link href={ROUTES.USER.LESSONS.DICTATION(lesson.id)} onClick={() => onOpenChange(false)}>
              <BookOpenText className="size-5" />
              Dictation
            </Link>
          </Button>
          <Button asChild size="lg" variant="outline" className="w-full justify-start gap-3">
            <Link href={ROUTES.USER.LESSONS.SHADOWING(lesson.id)} onClick={() => onOpenChange(false)}>
              <Mic className="size-5" />
              Shadowing
            </Link>
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}
