import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { AdminVocabCategoriesTable } from "@/features/vocabulary/components/admin/AdminVocabCategoriesTable"
import { getAdminVocabularyCategories } from "@/features/vocabulary/services/vocabulary.get"

export default async function AdminVocabCategoriesPage() {
  const res = await getAdminVocabularyCategories()
  const categories = res.data ?? []

  return (
    <AdminPageLayout>
      <AdminVocabCategoriesTable categories={categories} />
    </AdminPageLayout>
  )
}
