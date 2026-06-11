import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { AdminVocabDecksTable } from "@/features/vocabulary/components/admin/AdminVocabDecksTable"
import { getAdminVocabularyDecks, getAdminVocabularyCategories } from "@/features/vocabulary/services/vocabulary.get"

const DEFAULT_LIMIT = 10

export default async function AdminVocabDecksPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)

  const [res, catRes] = await Promise.all([
    getAdminVocabularyDecks(pageNum, limitNum),
    getAdminVocabularyCategories(),
  ])
  const decks = res.data ?? []
  const categories = catRes.data ?? []
  const meta = res.meta ?? { page: pageNum, limit: limitNum, total: 0, totalPages: 0 }

  return (
    <AdminPageLayout>
      <AdminVocabDecksTable decks={decks} categories={categories} meta={meta} limit={limitNum} />
    </AdminPageLayout>
  )
}
