import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { AdminVocabDecksTable } from "@/features/vocabulary/components/admin/AdminVocabDecksTable"
import { getAdminVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"

export default async function AdminVocabDecksPage() {
  const res = await getAdminVocabularyDecks()
  const decks = res.data ?? []

  return (
    <AdminPageLayout>
      <AdminVocabDecksTable decks={decks} />
    </AdminPageLayout>
  )
}
