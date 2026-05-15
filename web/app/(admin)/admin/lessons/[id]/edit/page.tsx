import { notFound } from "next/navigation"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonEditContent } from "@/features/lessons/components/admin/LessonEditContent"
import { getAdminLesson } from "@/features/lessons/services/lessons.get"
import { getAdminCategories } from "@/features/categories/services/categories.get"
import { ROUTES } from "@/lib/routes"

export default async function LessonEditPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const [lessonRes, categoriesRes] = await Promise.all([
    getAdminLesson(Number(id)),
    getAdminCategories(),
  ])

  if (!lessonRes.data) {
    notFound()
  }

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonEditContent lesson={lessonRes.data} categories={categoriesRes.data ?? []} />
    </AdminPageLayout>
  )
}
