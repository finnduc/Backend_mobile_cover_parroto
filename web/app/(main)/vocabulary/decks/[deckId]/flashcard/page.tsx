import { auth } from "@clerk/nextjs/server"
import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { FlashcardView } from "@/features/vocabulary/components/user/FlashcardView"
import { getVocabularyDeck, getVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"

export default async function FlashcardPage({
  params,
}: {
  params: Promise<{ deckId: string }>
}) {
  const { userId } = await auth()
  if (!userId) return null

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
        { label: "Tu vung", href: "/vocabulary" },
        { label: deck.name, href: `/vocabulary/decks/${deck.id}` },
        { label: "Flashcard" },
      ]}
    >
      <FlashcardView items={items} />
    </PageLayout>
  )
}
