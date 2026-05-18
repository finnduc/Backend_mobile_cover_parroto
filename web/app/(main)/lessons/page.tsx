import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonCard } from "@/features/lessons/components/user/LessonCard"
import { getLessons } from "@/features/lessons/services/lessons.get"
import { getCategories } from "@/features/categories/services/categories.get"
import { getBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { ROUTES } from "@/lib/routes"

export default async function LessonsPage({
  searchParams,
}: {
  searchParams: Promise<{ categoryId?: string }>
}) {
  const { categoryId } = await searchParams

  const categoriesRes = await getCategories()
  const categories = categoriesRes.data ?? []

  const category = categoryId
    ? categories.find((c) => c.id === Number(categoryId))
    : null

  const lessonsRes = await getLessons(1, 100, {
    categoryId: categoryId ? Number(categoryId) : undefined,
  })
  const lessons = lessonsRes.data ?? []

  const bookmarksRes = await getBookmarks()
  const bookmarkLessonIds = new Set((bookmarksRes.data ?? []).map((b) => b.lessionId))

  return (
    <PageLayout
      title={category ? category.name : "Tất cả bài học"}
      breadcrumbs={[
        { label: "Chủ đề", href: ROUTES.USER.CATEGORIES },
        { label: category ? category.name : "Tất cả bài học" },
      ]}
    >
      <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
        {lessons.map((lesson) => (
          <LessonCard key={lesson.id} lesson={lesson} isBookmarked={bookmarkLessonIds.has(lesson.id)} />
        ))}
      </div>
    </PageLayout>
  )
}
