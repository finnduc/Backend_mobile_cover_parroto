"use client"

import { useRouter } from "next/navigation"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm } from "@/features/lessons/components/admin/LessonForm"
import { ROUTES } from "@/lib/routes"
import type { Lesson } from "@/types/lessons.models"

export function LessonEditContent({ lesson }: { lesson: Lesson }) {
  const router = useRouter()

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
          onSubmit={() => router.push(ROUTES.ADMIN.LESSONS.LIST)}
        />
      </CardContent>
    </Card>
  )
}
