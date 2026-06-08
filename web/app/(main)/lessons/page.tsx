import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonCard } from "@/features/lessons/components/user/LessonCard"
import { getLessons } from "@/features/lessons/services/lessons.get"
import { getCategories } from "@/features/categories/services/categories.get"
import { getLessonBookmarks } from "@/features/lesson-bookmarks/services/lesson-bookmarks.get"
import { ROUTES } from "@/lib/routes"

export default async function LessonsPage({
  searchParams,
}: {
  searchParams: Promise<{ categoryId?: string }>
}) {
  const { userId } = await auth()
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

  const bookmarkedLessonIds = new Set<number>()
  if (userId) {
    const bookmarksRes = await getLessonBookmarks()
    const bookmarks = bookmarksRes.data ?? []
    bookmarks.forEach((b) => bookmarkedLessonIds.add(b.lessonId))
  }

  return (
    <PageLayout
      title={category ? category.name : "Tat ca bai hoc"}
      breadcrumbs={[
        { label: "Chu de", href: ROUTES.USER.CATEGORIES },
        { label: category ? category.name : "Tat ca bai hoc" },
      ]}
    >
      <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4">
        {lessons.map((lesson) => (
          <LessonCard
            key={lesson.id}
            lesson={lesson}
            isBookmarked={bookmarkedLessonIds.has(lesson.id)}
          />
        ))}
      </div>
    </PageLayout>
  )
}
