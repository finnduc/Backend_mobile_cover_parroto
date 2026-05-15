import { LessonsPageContent } from "@/features/lessons/components/admin/LessonsPageContent"
import { getAdminLessons } from "@/features/lessons/services/lessons.get"

const DEFAULT_LIMIT = 10

export default async function LessonsPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)
  const res = await getAdminLessons(pageNum, limitNum)

  return (
    <LessonsPageContent
      data={res.data ?? []}
      meta={res.meta ?? { page: pageNum, limit: limitNum, total: 0, totalPages: 0 }}
      limit={limitNum}
    />
  )
}
