"use client"

import { useRouter } from "next/navigation"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { LessonForm } from "@/features/lessons/components/admin/LessonForm"
import { ROUTES } from "@/lib/routes"

export function LessonCreateContent() {
  const router = useRouter()

  return (
    <Card>
      <CardHeader>
        <CardTitle>Create Lesson</CardTitle>
      </CardHeader>
      <CardContent>
        <LessonForm onSubmit={() => router.push(ROUTES.ADMIN.LESSONS.LIST)} />
      </CardContent>
    </Card>
  )
}
