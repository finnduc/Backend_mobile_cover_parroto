import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonCard } from "@/features/lessons/components/user/LessonCard"
import { mockLessons, mockCategories } from "@/features/lessons/mock-data"
import { ROUTES } from "@/lib/routes"

export default async function LessonsPage({
  searchParams,
}: {
  searchParams: Promise<{ categoryId?: string }>
}) {
  const { categoryId } = await searchParams
  const category = categoryId
    ? mockCategories.find((c) => c.id === Number(categoryId))
    : null
  const lessons = categoryId
    ? mockLessons.filter((l) => l.categoryId === Number(categoryId))
    : mockLessons

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
          <LessonCard key={lesson.id} lesson={lesson} />
        ))}
      </div>
    </PageLayout>
  )
}
