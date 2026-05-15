"use client"

import type { Lesson } from "@/types/lessons.models"
import type { Category } from "@/types/categories.models"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm, type LessonFormValues } from "@/features/lessons/components/admin/LessonForm"
import { updateAdminLesson } from "@/features/lessons/services/lessons.action"
import { toast } from "sonner"

export function LessonEditContent({ lesson, categories }: { lesson: Lesson; categories: Category[] }) {
  const handleSubmit = async (values: LessonFormValues) => {
    const res = await updateAdminLesson(lesson.id, values)
    if (res.error) {
      toast.error(res.error.message)
    } else {
      toast.success("Lesson updated successfully")
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Edit Lesson #{lesson.id}</CardTitle>
      </CardHeader>
      <CardContent>
        <LessonForm
          defaultValues={{
            title: lesson.title,
            description: lesson.description,
            thumbnailUrl: lesson.thumbnailUrl,
            videoUrl: lesson.videoUrl,
            duration: lesson.duration,
            level: lesson.level,
            categoryId: lesson.categoryId,
          }}
          categories={categories}
          onSubmit={handleSubmit}
        />
      </CardContent>
    </Card>
  )
}
