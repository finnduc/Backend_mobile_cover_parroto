"use client"

import { useState } from "react"
import { DeckItemList } from "@/features/vocabulary/components/user/DeckItemList"
import { AddItemModal } from "@/features/vocabulary/components/user/AddItemModal"
import { Button } from "@/components/ui/button"
import { Plus } from "lucide-react"
import type { VocabularyItem } from "@/types/vocabulary.models"

export function MyDeckDetailContent({
  initialItems,
  deckId,
}: {
  initialItems: VocabularyItem[]
  deckId: number
}) {
  const [modalOpen, setModalOpen] = useState(false)

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-sm font-semibold text-muted-foreground">
          {initialItems.length} từ vựng
        </h2>
        <Button onClick={() => setModalOpen(true)} size="sm">
          <Plus className="size-4" />
          Thêm từ vựng
        </Button>
      </div>
      <DeckItemList items={initialItems} />
      <AddItemModal
        open={modalOpen}
        onOpenChange={setModalOpen}
        deckId={deckId}
      />
    </div>
  )
}
