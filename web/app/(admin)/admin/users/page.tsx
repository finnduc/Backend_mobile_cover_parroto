import { UsersPageContent } from "@/features/users/components/admin/UsersPageContent";
import { getAdminUsers } from "@/features/users/services/users.get";

export default async function UsersPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || 10)
  const offset = (pageNum - 1) * limitNum

  const { users, totalCount } = await getAdminUsers(limitNum, offset)

  return (
    <UsersPageContent
      users={users}
      currentPage={pageNum}
      totalPages={Math.ceil(totalCount / limitNum)}
      limit={limitNum}
    />
  )
}