"use client"

import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { deleteAdminVocabularyItem } from "@/features/vocabulary/services/vocabulary.action"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { VocabularyItem } from "@/types/vocabulary.models"

export function AdminVocabItemsTable({
  items,
}: {
  items: VocabularyItem[]
}) {
  const columns: Column<VocabularyItem>[] = [
    { key: "id", header: "ID" },
    { key: "phrase", header: "Phrase" },
    { key: "meaning", header: "Meaning" },
    {
      key: "exampleSentence", header: "Example",
      render: (i) => (
        <span className="block max-w-[250px] truncate">
          {i.exampleSentence || "-"}
        </span>
      ),
    },
    {
      key: "actions", header: "",
      render: (i) => (
        <div className="flex justify-end gap-1">
          <Button size="xs" variant="outline">Edit</Button>
          <Button size="xs" variant="destructive" onClick={async () => {
            const res = await deleteAdminVocabularyItem(i.id)
            if (res.error) toast.error(res.error.message)
          }}>Delete</Button>
        </div>
      ),
    },
  ]

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Vocabulary Items</h2>
        <Button>Add Item</Button>
      </div>
      <DataTable columns={columns} data={items} />
    </div>
  )
}
