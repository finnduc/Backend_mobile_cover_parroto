"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { PaginationBar } from "@/components/common/PaginationBar"
import { VocabItemFormDialog } from "@/features/vocabulary/components/admin/VocabItemFormDialog"
import {
  createAdminVocabularyItem,
  updateAdminVocabularyItem,
  deleteAdminVocabularyItem,
} from "@/features/vocabulary/services/vocabulary.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { VocabularyItem } from "@/types/vocabulary.models"
import type { PaginatedMeta } from "@/types/base-response"
import type {
  CreateVocabularyItemDto,
  UpdateVocabularyItemDto,
} from "@/features/vocabulary/dtos/vocabulary.dto"

export function AdminVocabItemsTable({
  items,
  deckId,
  meta,
  limit,
}: {
  items: VocabularyItem[]
  deckId: number
  meta: PaginatedMeta
  limit: number
}) {
  const [createOpen, setCreateOpen] = useState(false)
  const [editing, setEditing] = useState<VocabularyItem | null>(null)

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
          <Button
            size="xs"
            variant="outline"
            onClick={() => setEditing(i)}
          >
            Edit
          </Button>
          <Button
            size="xs"
            variant="destructive"
            onClick={async () => {
              const res = await deleteAdminVocabularyItem(i.id)
              if (res.error) toast.error(res.error.message)
              else toast.success("Deleted")
            }}
          >
            Delete
          </Button>
        </div>
      ),
    },
  ]

  const handleCreate = async (values: CreateVocabularyItemDto | UpdateVocabularyItemDto) => {
    const res = await createAdminVocabularyItem(deckId, values as CreateVocabularyItemDto)
    if (res.error) toast.error(res.error.message)
    else toast.success("Created")
  }

  const handleEdit = async (values: CreateVocabularyItemDto | UpdateVocabularyItemDto) => {
    if (!editing) return
    const res = await updateAdminVocabularyItem(editing.id, values as UpdateVocabularyItemDto)
    if (res.error) toast.error(res.error.message)
    else {
      toast.success("Updated")
      setEditing(null)
    }
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Vocabulary Items</h2>
        <Button onClick={() => setCreateOpen(true)}>Add Item</Button>
      </div>
      <DataTable columns={columns} data={items} />
      <PaginationBar
        currentPage={meta.page}
        totalPages={meta.totalPages}
        baseUrl={`${ROUTES.ADMIN.VOCABULARY.DECKS.LIST}/${deckId}/items`}
        searchParams={new URLSearchParams({ limit: String(limit) })}
      />

      <VocabItemFormDialog
        open={createOpen}
        onOpenChange={setCreateOpen}
        mode="create"
        onSubmit={handleCreate}
      />

      {editing && (
        <VocabItemFormDialog
          open={!!editing}
          onOpenChange={(open) => { if (!open) setEditing(null) }}
          mode="edit"
          initialValues={{
            phrase: editing.phrase,
            meaning: editing.meaning ?? "",
            exampleSentence: editing.exampleSentence ?? "",
          }}
          onSubmit={handleEdit}
        />
      )}
    </div>
  )
}
