"use client"

import { useRouter } from "next/navigation"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm, type LessonFormValues } from "@/features/lessons/components/admin/LessonForm"
import { updateAdminLesson } from "@/features/lessons/services/lessons.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"
import type { Lesson } from "@/types/lessons.models"

export function LessonEditContent({ lesson }: { lesson: Lesson }) {
  const router = useRouter()

  const handleSubmit = async (values: LessonFormValues) => {
    const res = await updateAdminLesson(lesson.id, values)
    if (res.error) {
      toast.error(res.error.message)
    } else {
      router.push(ROUTES.ADMIN.LESSONS.LIST)
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
          onSubmit={handleSubmit}
        />
      </CardContent>
    </Card>
  )
}
