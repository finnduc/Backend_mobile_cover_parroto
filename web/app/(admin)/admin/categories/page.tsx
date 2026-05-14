import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { CategoriesPageContent } from "@/features/categories/components/admin/CategoriesPageContent"
import { getAdminCategories } from "@/features/categories/services/categories-service"

export default async function CategoriesAdminPage() {
  const res = await getAdminCategories()
  const categories = res.data ?? []

  return (
    <AdminPageLayout>
      <CategoriesPageContent categories={categories} />
    </AdminPageLayout>
  )
}
