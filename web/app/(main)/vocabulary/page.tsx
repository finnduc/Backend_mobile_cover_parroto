import { PageLayout } from "@/components/layouts/PageLayout"
import { VocabularyBrowser } from "@/features/vocabulary/components/user/VocabularyBrowser"
import { getVocabularyCategories, getSystemVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
export default async function VocabularyPage() {
  const categoriesRes = await getVocabularyCategories()
  const categories = categoriesRes.data ?? []

  const decksRes = await getSystemVocabularyDecks()
  const decks = decksRes.data ?? []

  return (
    <PageLayout
      title="Từ vựng"
      breadcrumbs={[
        { label: "Từ vựng" },
      ]}
    >
      <VocabularyBrowser categories={categories} decks={decks} />
    </PageLayout>
  )
}
