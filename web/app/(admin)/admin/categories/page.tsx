import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { CategoriesPageContent } from "@/features/categories/components/admin/CategoriesPageContent"
import { mockCategories } from "@/features/lessons/mock-data"

export default function CategoriesAdminPage() {
  return (
    <AdminPageLayout>
      <CategoriesPageContent categories={mockCategories} />
    </AdminPageLayout>
  )
}
