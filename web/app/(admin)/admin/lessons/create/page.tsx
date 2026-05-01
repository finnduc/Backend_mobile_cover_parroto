import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { LessonCreateContent } from "@/features/lessons/components/admin/LessonCreateContent"
import { ROUTES } from "@/lib/routes"

export default function LessonCreatePage() {
  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <LessonCreateContent />
    </AdminPageLayout>
  )
}
