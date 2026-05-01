import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonEditContent } from "@/features/lessons/components/admin/LessonEditContent"
import { getLesson } from "@/features/lessons/services/lessons-service"
import { ROUTES } from "@/lib/routes"

export default async function LessonEditPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const lesson = getLesson(Number(id))

  if (!lesson) {
    return <div className="py-12 text-center text-muted-foreground">Lesson not found</div>
  }

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonEditContent lesson={lesson} />
    </AdminPageLayout>
  )
}
