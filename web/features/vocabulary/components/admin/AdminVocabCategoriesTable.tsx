"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { PaginationBar } from "@/components/common/PaginationBar"
import { CreateModal } from "@/components/common/CreateModal"
import { EditModal } from "@/components/common/EditModal"
import { VocabularyCategoryForm } from "@/features/vocabulary/components/admin/VocabularyCategoryForm"
import { createVocabularyCategory, updateVocabularyCategory, deleteVocabularyCategory } from "@/features/vocabulary/services/vocabulary.action"
import { ROUTES } from "@/lib/routes"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { VocabularyCategory } from "@/types/vocabulary.models"
import type { PaginatedMeta } from "@/types/base-response"

export function AdminVocabCategoriesTable({
  categories,
  meta,
  limit,
}: {
  categories: VocabularyCategory[]
  meta: PaginatedMeta
  limit: number
}) {
  const [createOpen, setCreateOpen] = useState(false)
  const [editOpen, setEditOpen] = useState(false)
  const [editing, setEditing] = useState<VocabularyCategory | null>(null)

  const columns: Column<VocabularyCategory>[] = [
    { key: "id", header: "ID" },
    { key: "name", header: "Name" },
    { key: "description", header: "Description" },
    {
      key: "actions",
      header: "",
      render: (cat) => (
        <div className="flex justify-end gap-2">
          <Button size="xs" variant="outline" onClick={() => { setEditing(cat); setEditOpen(true) }}>Edit</Button>
          <Button size="xs" variant="destructive" onClick={async () => {
            const res = await deleteVocabularyCategory(cat.id)
            if (res.error) toast.error(res.error.message)
            else toast.success("Deleted")
          }}>Delete</Button>
        </div>
      ),
    },
  ]

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Vocabulary Categories</h2>
        <Button onClick={() => setCreateOpen(true)}>Create Category</Button>
      </div>
      <DataTable columns={columns} data={categories} />
      <PaginationBar
        currentPage={meta.page}
        totalPages={meta.totalPages}
        baseUrl={ROUTES.ADMIN.VOCABULARY.CATEGORIES.LIST}
        searchParams={new URLSearchParams({ limit: String(limit) })}
      />
      <CreateModal open={createOpen} onOpenChange={setCreateOpen} title="Create Vocabulary Category" submitLabel="">
        <VocabularyCategoryForm onSubmit={async (values) => {
          const res = await createVocabularyCategory({ name: values.name, description: values.description })
          if (res.error) toast.error(res.error.message)
          else setCreateOpen(false)
        }} />
      </CreateModal>
      <EditModal open={editOpen} onOpenChange={setEditOpen} title="Edit Vocabulary Category" submitLabel="">
        {editing && (
          <VocabularyCategoryForm
            key={editing.id}
            defaultValues={{ name: editing.name, description: editing.description }}
            onSubmit={async (values) => {
              if (!editing) return
              const res = await updateVocabularyCategory(editing.id, { name: values.name, description: values.description })
              if (res.error) toast.error(res.error.message)
              else setEditOpen(false)
            }}
          />
        )}
      </EditModal>
    </div>
  )
}
