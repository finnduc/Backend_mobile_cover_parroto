import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { AdminVocabItemsTable } from "@/features/vocabulary/components/admin/AdminVocabItemsTable"
import { getAdminVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"

export default async function AdminVocabItemsPage({
  params,
}: {
  params: Promise<{ deckId: string }>
}) {
  const { deckId } = await params
  const itemsRes = await getAdminVocabularyItems(Number(deckId))
  const items = itemsRes.data ?? []

  return (
    <AdminPageLayout
      title={`Deck #${deckId} Items`}
      backHref={ROUTES.ADMIN.VOCABULARY.DECKS.LIST}
      backLabel="Back to Decks"
    >
      <AdminVocabItemsTable items={items} />
    </AdminPageLayout>
  )
}
