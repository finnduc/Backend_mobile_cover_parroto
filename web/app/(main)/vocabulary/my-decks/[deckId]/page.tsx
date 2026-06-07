import { PageLayout } from "@/components/layouts/PageLayout"
import { MyDeckDetailContent } from "@/features/vocabulary/components/user/MyDeckDetailContent"
import { getVocabularyDeck, getVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"
import { Button } from "@/components/ui/button"
import { Layers } from "lucide-react"
import Link from "next/link"
import { notFound } from "next/navigation"

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
        { label: "Tu vung", href: "/vocabulary" },
        { label: "Bo tu vung cua toi", href: ROUTES.USER.VOCABULARY.MY_DECKS },
        { label: deck.name },
      ]}
      actions={
        items.length > 0 ? (
          <Button asChild variant="outline" size="sm">
            <Link href={`/vocabulary/my-decks/${deck.id}/flashcard`}>
              <Layers className="mr-1 size-4" />
              Flashcard
            </Link>
          </Button>
        ) : undefined
      }
    >
      <MyDeckDetailContent initialItems={items} deckId={deck.id} />
    </PageLayout>
  )
}
