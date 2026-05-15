import Link from "next/link"
import { DataTable } from "@/components/common/DataTable"
import { PaginationBar } from "@/components/common/PaginationBar"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { getAdminUsers } from "@/features/users/services/users.get"
import { ROUTES } from "@/lib/routes"
import type { Column } from "@/components/common/DataTable"
import type { User } from "@/types/users.models"

const DEFAULT_LIMIT = 10

export default async function UsersPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)
  const res = await getAdminUsers()
  const users = res.data ?? []
  const totalPages = Math.ceil(users.length / limitNum)
  const data = users.slice((pageNum - 1) * limitNum, pageNum * limitNum)

  const columns: Column<User>[] = [
    { key: "id", header: "ID" },
    { key: "name", header: "Name" },
    { key: "email", header: "Email" },
    {
      key: "actions",
      header: "",
      render: (user) => (
        <Link href={ROUTES.ADMIN.USERS.DETAIL(String(user.id))} className="text-xs text-primary hover:underline">
          View
        </Link>
      ),
    },
  ]

  return (
    <AdminPageLayout title="Users">
      <DataTable columns={columns} data={data} />
      <PaginationBar
        currentPage={pageNum}
        totalPages={totalPages}
        baseUrl={ROUTES.ADMIN.USERS.LIST}
        searchParams={new URLSearchParams({ limit: String(limitNum) })}
      />
    </AdminPageLayout>
  )
}
