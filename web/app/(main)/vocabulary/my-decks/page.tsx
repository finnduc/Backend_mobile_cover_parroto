import { PageLayout } from "@/components/layouts/PageLayout"
import { MyDecksContent } from "@/features/vocabulary/components/user/MyDecksContent"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"

export default async function MyDecksPage() {
  const decksRes = await getUserVocabularyDecks()
  const decks = decksRes.data ?? []

  return (
    <PageLayout
      title="Bộ từ vựng của tôi"
      breadcrumbs={[
        { label: "Từ vựng", href: "/vocabulary" },
        { label: "Bộ từ vựng của tôi" },
      ]}
    >
      <MyDecksContent decks={decks} />
    </PageLayout>
  )
}
