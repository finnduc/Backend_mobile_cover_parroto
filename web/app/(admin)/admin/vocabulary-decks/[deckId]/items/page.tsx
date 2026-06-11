import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { AdminVocabItemsTable } from "@/features/vocabulary/components/admin/AdminVocabItemsTable"
import { getAdminVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"

const DEFAULT_LIMIT = 10

export default async function AdminVocabItemsPage({
  params,
  searchParams,
}: {
  params: Promise<{ deckId: string }>
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { deckId } = await params
  const { page, limit } = await searchParams
  const deckIdNum = Number(deckId)
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)

  const itemsRes = await getAdminVocabularyItems(deckIdNum, pageNum, limitNum)
  const items = itemsRes.data ?? []
  const meta = itemsRes.meta ?? { page: pageNum, limit: limitNum, total: 0, totalPages: 0 }

  return (
    <AdminPageLayout
      title={`Deck #${deckId} Items`}
      backHref={ROUTES.ADMIN.VOCABULARY.DECKS.LIST}
      backLabel="Back to Decks"
    >
      <AdminVocabItemsTable items={items} deckId={deckIdNum} meta={meta} limit={limitNum} />
    </AdminPageLayout>
  )
}
