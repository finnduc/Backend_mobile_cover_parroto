import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonEditContent } from "@/features/lessons/components/admin/LessonEditContent"
import { getAdminLesson } from "@/features/lessons/services/lessons-service"
import { ROUTES } from "@/lib/routes"

export default async function LessonEditPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const res = await getAdminLesson(Number(id))

  if (res.error) {
    return <div className="py-12 text-center text-muted-foreground">{res.error.message}</div>
  }
  if (!res.data) {
    return <div className="py-12 text-center text-muted-foreground">Lesson not found</div>
  }

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonEditContent lesson={res.data} />
    </AdminPageLayout>
  )
}
