"use client"

import { useState } from "react"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { PaginationBar } from "@/components/common/PaginationBar"
import { VocabDeckFormDialog } from "@/features/vocabulary/components/admin/VocabDeckFormDialog"
import {
  createAdminVocabularyDeck,
  updateAdminVocabularyDeck,
  deleteAdminVocabularyDeck,
} from "@/features/vocabulary/services/vocabulary.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { VocabularyDeck, VocabularyCategory } from "@/types/vocabulary.models"
import type { PaginatedMeta } from "@/types/base-response"
import type {
  CreateVocabularyDeckDto,
  UpdateVocabularyDeckDto,
} from "@/features/vocabulary/dtos/vocabulary.dto"

export function AdminVocabDecksTable({
  decks,
  categories,
  meta,
  limit,
}: {
  decks: VocabularyDeck[]
  categories: VocabularyCategory[]
  meta: PaginatedMeta
  limit: number
}) {
  const [createOpen, setCreateOpen] = useState(false)
  const [editing, setEditing] = useState<VocabularyDeck | null>(null)

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
          <Button
            size="xs"
            variant="outline"
            onClick={() => setEditing(d)}
          >
            Edit
          </Button>
          <Button size="xs" variant="destructive" onClick={async () => {
            const res = await deleteAdminVocabularyDeck(d.id)
            if (res.error) toast.error(res.error.message)
            else toast.success("Deleted")
          }}>Delete</Button>
        </div>
      ),
    },
  ]

  const handleCreate = async (values: CreateVocabularyDeckDto | UpdateVocabularyDeckDto) => {
    const res = await createAdminVocabularyDeck(values as CreateVocabularyDeckDto)
    if (res.error) toast.error(res.error.message)
    else toast.success("Created")
  }

  const handleEdit = async (values: CreateVocabularyDeckDto | UpdateVocabularyDeckDto) => {
    if (!editing) return
    const res = await updateAdminVocabularyDeck(editing.id, values as UpdateVocabularyDeckDto)
    if (res.error) toast.error(res.error.message)
    else {
      toast.success("Updated")
      setEditing(null)
    }
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Vocabulary Decks</h2>
        <Button onClick={() => setCreateOpen(true)}>Create Deck</Button>
      </div>
      <DataTable columns={columns} data={decks} />
      <PaginationBar
        currentPage={meta.page}
        totalPages={meta.totalPages}
        baseUrl={ROUTES.ADMIN.VOCABULARY.DECKS.LIST}
        searchParams={new URLSearchParams({ limit: String(limit) })}
      />

      <VocabDeckFormDialog
        open={createOpen}
        onOpenChange={setCreateOpen}
        mode="create"
        categories={categories}
        onSubmit={handleCreate}
      />

      {editing && (
        <VocabDeckFormDialog
          open={!!editing}
          onOpenChange={(open) => { if (!open) setEditing(null) }}
          mode="edit"
          categories={categories}
          initialValues={{
            name: editing.name,
            description: editing.description ?? "",
            level: editing.level ?? "",
            categoryId: editing.categoryId ?? null,
            thumbnailUrl: editing.thumbnailUrl ?? "",
          }}
          onSubmit={handleEdit}
        />
      )}
    </div>
  )
}
