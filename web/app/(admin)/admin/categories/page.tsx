import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { CategoriesPageContent } from "@/features/categories/components/admin/CategoriesPageContent"
import { getAdminCategories } from "@/features/categories/services/categories.get"

const DEFAULT_LIMIT = 10

export default async function CategoriesAdminPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)

  const res = await getAdminCategories(pageNum, limitNum)
  const categories = res.data ?? []
  const meta = res.meta ?? { page: pageNum, limit: limitNum, total: 0, totalPages: 0 }

  return (
    <AdminPageLayout>
      <CategoriesPageContent categories={categories} meta={meta} limit={limitNum} />
    </AdminPageLayout>
  )
}
