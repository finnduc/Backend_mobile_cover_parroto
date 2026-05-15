import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonCreateContent } from "@/features/lessons/components/admin/LessonCreateContent"
import { getAdminCategories } from "@/features/categories/services/categories.get"
import { ROUTES } from "@/lib/routes"

export default async function LessonCreatePage() {
  const res = await getAdminCategories()
  const categories = res.data ?? []

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonCreateContent categories={categories} />
    </AdminPageLayout>
  )
}
