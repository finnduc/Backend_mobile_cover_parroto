import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { MyDeckDetailContent } from "@/features/vocabulary/components/user/MyDeckDetailContent"
import { getVocabularyDeck, getVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"

export default async function MyDeckDetailPage({
  params,
}: {
  params: Promise<{ deckId: string }>
}) {
  const { deckId } = await params
  const deckRes = await getVocabularyDeck(Number(deckId))

  if (!deckRes.data) {
    notFound()
  }

  const deck = deckRes.data
  const itemsRes = await getVocabularyItems(deck.id)
  const items = itemsRes.data ?? []

  return (
    <PageLayout
      title={deck.name}
      breadcrumbs={[
        { label: "Từ vựng", href: "/vocabulary" },
        { label: "Bộ từ vựng của tôi", href: ROUTES.USER.VOCABULARY.MY_DECKS },
        { label: deck.name },
      ]}
    >
      <MyDeckDetailContent initialItems={items} deckId={deck.id} />
    </PageLayout>
  )
}
