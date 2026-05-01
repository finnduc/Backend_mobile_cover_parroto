import { LessonsPageContent } from "@/features/lessons/components/admin/LessonsPageContent"
import { getLessons } from "@/features/lessons/services/lessons-service"

const DEFAULT_LIMIT = 10

export default async function LessonsPage({
  searchParams,
}: {
  searchParams: Promise<{ page?: string; limit?: string }>
}) {
  const { page, limit } = await searchParams
  const pageNum = Math.max(1, Number(page) || 1)
  const limitNum = Math.max(1, Number(limit) || DEFAULT_LIMIT)
  const { data, meta } = getLessons(pageNum, limitNum)

  return <LessonsPageContent data={data} meta={meta} limit={limitNum} />
}
