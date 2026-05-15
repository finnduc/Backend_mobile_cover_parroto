import { PaginationBar } from "@/components/common/PaginationBar"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { UsersPageContent } from "@/features/users/components/admin/UsersPageContent"
import { getAdminUsers } from "@/features/users/services/users.get"
import { ROUTES } from "@/lib/routes"

const DEFAULT_LIMIT = 10

export default async function UsersPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)
  const offset = (pageNum - 1) * limitNum
  const { users, totalCount } = await getAdminUsers(limitNum, offset)

  return (
    <AdminPageLayout title="Users">
      <UsersPageContent users={users} />
      <PaginationBar
        currentPage={pageNum}
        totalPages={Math.ceil(totalCount / limitNum)}
        baseUrl={ROUTES.ADMIN.USERS.LIST}
        searchParams={new URLSearchParams({ limit: String(limitNum) })}
      />
    </AdminPageLayout>
  )
}
