"use client"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { deleteAdminVocabularyDeck } from "@/features/vocabulary/services/vocabulary.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function AdminVocabDecksTable({
  decks,
}: {
  decks: VocabularyDeck[]
}) {
  const columns: Column<VocabularyDeck>[] = [
    { key: "id", header: "ID" },
    { key: "name", header: "Name" },
    { key: "level", header: "Level" },
    {
      key: "isDefault", header: "System",
      render: (d) => (d.isDefault ? "Yes" : "No"),
    },
    {
      key: "actions", header: "",
      render: (d) => (
        <div className="flex justify-end gap-1">
          <Button size="xs" variant="outline" asChild>
            <Link href={ROUTES.ADMIN.VOCABULARY.DECKS.ITEMS(d.id)}>Items</Link>
          </Button>
          <Button size="xs" variant="outline">Edit</Button>
          <Button size="xs" variant="destructive" onClick={async () => {
            const res = await deleteAdminVocabularyDeck(d.id)
            if (res.error) toast.error(res.error.message)
          }}>Delete</Button>
        </div>
      ),
    },
  ]

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Vocabulary Decks</h2>
        <Button>Create Deck</Button>
      </div>
      <DataTable columns={columns} data={decks} />
    </div>
  )
}
