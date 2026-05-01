"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { CreateModal } from "@/components/common/CreateModal"
import { EditModal } from "@/components/common/EditModal"
import { CategoryForm, type CategoryFormValues } from "@/features/categories/components/admin/CategoryForm"
import type { Column } from "@/components/common/DataTable"
import type { Category } from "@/types/categories.models"

export function CategoriesPageContent({
  categories: initialCategories,
}: {
  categories: Category[]
}) {
  const [categories, setCategories] = useState(initialCategories)
  const [createOpen, setCreateOpen] = useState(false)
  const [editOpen, setEditOpen] = useState(false)
  const [editing, setEditing] = useState<Category | null>(null)

  const columns: Column<Category>[] = [
    { key: "id", header: "ID" },
    { key: "name", header: "Name" },
    {
      key: "actions",
      header: "",
      render: (cat) => (
        <div className="flex justify-end gap-2">
          <Button
            size="xs"
            variant="outline"
            onClick={() => {
              setEditing(cat)
              setEditOpen(true)
            }}
          >
            Edit
          </Button>
          <Button
            size="xs"
            variant="destructive"
            onClick={() =>
              setCategories((prev) => prev.filter((c) => c.id !== cat.id))
            }
          >
            Delete
          </Button>
        </div>
      ),
    },
  ]

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-xl font-bold">Categories</h2>
        <Button onClick={() => setCreateOpen(true)}>
          Create Category
        </Button>
      </div>

      <DataTable columns={columns} data={categories} />

      <CreateModal
        open={createOpen}
        onOpenChange={setCreateOpen}
        title="Create Category"
        submitLabel=""
      >
        <CategoryForm
          onSubmit={(values) => {
            const newId = Math.max(0, ...categories.map((c) => c.id)) + 1
            setCategories((prev) => [...prev, { id: newId, name: values.name }])
            setCreateOpen(false)
          }}
        />
      </CreateModal>

      <EditModal
        open={editOpen}
        onOpenChange={setEditOpen}
        title="Edit Category"
        submitLabel=""
      >
        {editing && (
          <CategoryForm
            key={editing.id}
            defaultValues={{ name: editing.name }}
            onSubmit={(values) => {
              if (!editing) return
              setCategories((prev) =>
                prev.map((c) => (c.id === editing.id ? { ...c, name: values.name } : c))
              )
              setEditOpen(false)
            }}
          />
        )}
      </EditModal>
    </div>
  )
}
