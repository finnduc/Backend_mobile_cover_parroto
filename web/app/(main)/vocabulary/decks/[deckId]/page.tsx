import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { DeckItemList } from "@/features/vocabulary/components/user/DeckItemList"
import { getVocabularyDeck, getVocabularyItems } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"
import { Button } from "@/components/ui/button"
import { Layers } from "lucide-react"
import Link from "next/link"

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
        { label: "Tu vung", href: ROUTES.USER.VOCABULARY.LIST },
        { label: deck.name },
      ]}
      actions={
        items.length > 0 ? (
          <Button asChild variant="outline" size="sm">
            <Link href={`/vocabulary/decks/${deck.id}/flashcard`}>
              <Layers className="mr-1 size-4" />
              Flashcard
            </Link>
          </Button>
        ) : undefined
      }
    >
      {deck.description && (
        <p className="text-sm text-muted-foreground">{deck.description}</p>
      )}
      <DeckItemList items={items} />
    </PageLayout>
  )
}
