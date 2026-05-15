"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { CreateModal } from "@/components/common/CreateModal"
import { EditModal } from "@/components/common/EditModal"
import { UserCreateForm, UserEditForm } from "@/features/users/components/admin/UserForm"
import { createAdminUser, updateAdminUser, deleteAdminUser } from "@/features/users/services/users.action"
import { toast } from "sonner"
import type { Column } from "@/components/common/DataTable"
import type { User } from "@/types/users.models"
import type { CreateUserDto, UpdateUserDto } from "@/features/users/dtos/user.dto"

export function UsersPageContent({ users }: { users: User[] }) {
  const [createOpen, setCreateOpen] = useState(false)
  const [editOpen, setEditOpen] = useState(false)
  const [editing, setEditing] = useState<User | null>(null)

  const columns: Column<User>[] = [
    { key: "id", header: "ID" },
    { key: "name", header: "Name" },
    { key: "email", header: "Email" },
    {
      key: "actions",
      header: "",
      render: (user) => (
        <div className="flex justify-end gap-2">
          <Button
            size="xs"
            variant="outline"
            onClick={() => {
              setEditing(user)
              setEditOpen(true)
            }}
          >
            Edit
          </Button>
          <Button
            size="xs"
            variant="destructive"
            onClick={async () => {
              const res = await deleteAdminUser(user.id)
              if (res.error) {
                toast.error(res.error.message)
              }
            }}
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
        <h2 className="text-xl font-bold">Users</h2>
        <Button onClick={() => setCreateOpen(true)}>Create User</Button>
      </div>

      <DataTable columns={columns} data={users} />

      <CreateModal
        open={createOpen}
        onOpenChange={setCreateOpen}
        title="Create User"
        submitLabel=""
      >
        <UserCreateForm
          onSubmit={async (values: CreateUserDto) => {
            const res = await createAdminUser(values)
            if (res.error) {
              toast.error(res.error.message)
            } else {
              setCreateOpen(false)
              toast.success("User created successfully")
            }
          }}
        />
      </CreateModal>

      <EditModal
        open={editOpen}
        onOpenChange={setEditOpen}
        title="Edit User"
        submitLabel=""
      >
        {editing && (
          <UserEditForm
            key={editing.id}
            defaultValues={{ name: editing.name }}
            onSubmit={async (values: UpdateUserDto) => {
              if (!editing) return
              const res = await updateAdminUser(editing.id, values)
              if (res.error) {
                toast.error(res.error.message)
              } else {
                setEditOpen(false)
                toast.success("User updated successfully")
              }
            }}
          />
        )}
      </EditModal>
    </div>
  )
}
