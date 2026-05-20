"use client"

import { useState } from "react"
import { VocabularyDeckCard } from "@/features/vocabulary/components/user/VocabularyDeckCard"
import { Button } from "@/components/ui/button"
import { BookOpen } from "lucide-react"
import type { VocabularyCategory } from "@/types/vocabulary.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function VocabularyBrowser({
  categories,
  decks,
}: {
  categories: VocabularyCategory[]
  decks: VocabularyDeck[]
}) {
  const [activeCategoryId, setActiveCategoryId] = useState<number | null>(null)

  const filteredDecks = activeCategoryId
    ? decks.filter((d) => d.categoryId === activeCategoryId)
    : decks

  return (
    <div className="space-y-6">
      <div className="flex flex-wrap gap-2">
        <Button
          variant={activeCategoryId === null ? "default" : "outline"}
          size="sm"
          onClick={() => setActiveCategoryId(null)}
        >
          Tất cả
        </Button>
        {categories.map((cat) => (
          <Button
            key={cat.id}
            variant={activeCategoryId === cat.id ? "default" : "outline"}
            size="sm"
            onClick={() => setActiveCategoryId(cat.id)}
          >
            {cat.name}
          </Button>
        ))}
      </div>

      {filteredDecks.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
          <BookOpen className="mb-4 size-12" />
          <p className="text-lg font-medium">Chưa có bộ từ vựng nào</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {filteredDecks.map((deck) => (
            <VocabularyDeckCard
              key={deck.id}
              deck={deck}
              href={`/vocabulary/decks/${deck.id}`}
            />
          ))}
        </div>
      )}
    </div>
  )
}
