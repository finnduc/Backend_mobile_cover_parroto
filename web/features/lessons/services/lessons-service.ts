import type { Lesson } from "@/types/lessons.models"
import { mockLessons } from "@/features/lessons/mock-data"
import type { PaginatedMeta } from "@/types/base-response"

export function getLessons(
  page = 1,
  limit = 10,
  filters?: { categoryId?: number; level?: string }
): { data: Lesson[]; meta: PaginatedMeta } {
  let filtered = [...mockLessons]
  if (filters?.categoryId) {
    filtered = filtered.filter((l) => l.categoryId === filters.categoryId)
  }
  if (filters?.level) {
    filtered = filtered.filter((l) => l.level === filters.level)
  }
  const total = filtered.length
  const totalPages = Math.ceil(total / limit)
  const start = (page - 1) * limit
  const data = filtered.slice(start, start + limit)
  return { data, meta: { page, limit, total, totalPages } }
}

export function getLesson(id: number): Lesson | undefined {
  return mockLessons.find((l) => l.id === id)
}
