"use client"

import { useRouter } from "next/navigation"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm, type LessonFormValues } from "@/features/lessons/components/admin/LessonForm"
import { createAdminLesson } from "@/features/lessons/services/lessons.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"

export function LessonCreateContent() {
  const router = useRouter()

  const handleSubmit = async (values: LessonFormValues) => {
    const res = await createAdminLesson(values)
    if (res.error) {
      toast.error(res.error.message)
    } else {
      router.push(ROUTES.ADMIN.LESSONS.LIST)
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Create Lesson</CardTitle>
      </CardHeader>
      <CardContent>
        <LessonForm onSubmit={handleSubmit} />
      </CardContent>
    </Card>
  )
}
