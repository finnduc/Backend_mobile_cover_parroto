"use client"

import { useState } from "react"
import { VocabularyDeckCard } from "@/features/vocabulary/components/user/VocabularyDeckCard"
import { CreateDeckModal } from "@/features/vocabulary/components/user/CreateDeckModal"
import { Button } from "@/components/ui/button"
import { BookOpen, Plus } from "lucide-react"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function MyDecksContent({ decks: initialDecks }: { decks: VocabularyDeck[] }) {
  const [decks, setDecks] = useState<VocabularyDeck[]>(initialDecks)
  const [createOpen, setCreateOpen] = useState(false)

  const handleCreateDeck = (values: { name: string; description: string; level: string }) => {
    const newDeck: VocabularyDeck = {
      id: Date.now(),
      userId: "user_001",
      categoryId: null,
      name: values.name,
      description: values.description,
      thumbnailUrl: "",
      level: values.level,
      isDefault: false,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    }
    setDecks((prev) => [newDeck, ...prev])
  }

  if (decks.length === 0) {
    return (
      <>
        <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
          <BookOpen className="mb-4 size-12" />
          <p className="text-lg font-medium">Chưa có bộ từ vựng nào</p>
          <p className="text-sm">Tạo bộ từ vựng đầu tiên của bạn</p>
          <Button className="mt-4" onClick={() => setCreateOpen(true)}>
            <Plus className="size-4" />
            Tạo bộ từ vựng
          </Button>
        </div>
        <CreateDeckModal open={createOpen} onOpenChange={setCreateOpen} onSubmit={handleCreateDeck} />
      </>
    )
  }

  return (
    <>
      <div className="mb-4 flex justify-end">
        <Button size="sm" onClick={() => setCreateOpen(true)}>
          <Plus className="size-4" />
          Tạo bộ từ vựng
        </Button>
      </div>
      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {decks.map((deck) => (
          <VocabularyDeckCard
            key={deck.id}
            deck={deck}
            href={`/vocabulary/my-decks/${deck.id}`}
          />
        ))}
      </div>
      <CreateDeckModal open={createOpen} onOpenChange={setCreateOpen} onSubmit={handleCreateDeck} />
    </>
  )
}
