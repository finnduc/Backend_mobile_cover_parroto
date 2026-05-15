"use client"

import type { Category } from "@/types/categories.models"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm, type LessonFormValues } from "@/features/lessons/components/admin/LessonForm"
import { createAdminLesson } from "@/features/lessons/services/lessons.action"
import { toast } from "sonner"

export function LessonCreateContent({ categories }: { categories: Category[] }) {
  const handleSubmit = async (values: LessonFormValues) => {
    const res = await createAdminLesson(values)
    if (res.error) {
      toast.error(res.error.message)
    } else {
      toast.success("Lesson created successfully")
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Create Lesson</CardTitle>
      </CardHeader>
      <CardContent>
        <LessonForm categories={categories} onSubmit={handleSubmit} />
      </CardContent>
    </Card>
  )
}
