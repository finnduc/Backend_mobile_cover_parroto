import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { SavedLessonsPageContent } from "@/features/lesson-bookmarks/components/SavedLessonsPageContent"
import { getLessonBookmarks } from "@/features/lesson-bookmarks/services/lesson-bookmarks.get"
import { getLessons } from "@/features/lessons/services/lessons.get"

export default async function BookmarksPage() {
  const { userId } = await auth()
  if (!userId) return null

  const bookmarksRes = await getLessonBookmarks()
  const bookmarks = bookmarksRes.data ?? []

  const lessonsRes = await getLessons(1, 100)
  const lessons = lessonsRes.data ?? []

  const bookmarksWithLessons = bookmarks.map((bookmark) => ({
    ...bookmark,
    lesson: lessons.find((l) => l.id === bookmark.lessonId),
  }))

  return (
    <PageLayout
      title="Bài học đã lưu"
      breadcrumbs={[
        { label: "Bài học", href: "/lessons" },
        { label: "Đã lưu" },
      ]}
    >
      <SavedLessonsPageContent bookmarks={bookmarksWithLessons} />
    </PageLayout>
  )
}
