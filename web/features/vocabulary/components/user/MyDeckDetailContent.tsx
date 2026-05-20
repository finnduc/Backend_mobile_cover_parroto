"use client"

import { useState } from "react"
import { DeckItemList } from "@/features/vocabulary/components/user/DeckItemList"
import { AddItemForm } from "@/features/vocabulary/components/user/AddItemForm"
import type { VocabularyItem } from "@/types/vocabulary.models"

export function MyDeckDetailContent({
  initialItems,
  deckId,
}: {
  initialItems: VocabularyItem[]
  deckId: number
}) {
  const [items, setItems] = useState<VocabularyItem[]>(initialItems)

  const handleAddItem = (values: {
    phrase: string
    normalizedPhrase: string
    meaning: string
    exampleSentence: string
    note: string
  }) => {
    const newItem: VocabularyItem = {
      id: Date.now(),
      deckId,
      lessonId: null,
      transcriptId: null,
      ...values,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    }
    setItems((prev) => [...prev, newItem])
  }

  return (
    <div className="space-y-6">
      <AddItemForm onSubmit={handleAddItem} />
      <DeckItemList items={items} />
    </div>
  )
}
