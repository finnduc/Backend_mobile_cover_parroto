import { notFound } from "next/navigation"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonEditContent } from "@/features/lessons/components/admin/LessonEditContent"
import { getAdminLesson } from "@/features/lessons/services/lessons.get"
import { ROUTES } from "@/lib/routes"

export default async function LessonEditPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const res = await getAdminLesson(Number(id))

  if (!res.data) {
    notFound()
  }

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonEditContent lesson={res.data} />
    </AdminPageLayout>
  )
}
