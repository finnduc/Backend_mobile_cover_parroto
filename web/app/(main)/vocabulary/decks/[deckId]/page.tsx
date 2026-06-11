import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { DeckItemList } from "@/features/vocabulary/components/user/DeckItemList"
import { getVocabularyDeck, getVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"

export default async function SystemDeckDetailPage({
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
        { label: "Từ vựng"},
        {label: "Từ điển", href: ROUTES.USER.VOCABULARY.LIST},
        { label: deck.name },
      ]}
    >
      {deck.description && (
        <p className="text-sm text-muted-foreground">{deck.description}</p>
      )}
      <DeckItemList items={items} />
    </PageLayout>
  )
}
