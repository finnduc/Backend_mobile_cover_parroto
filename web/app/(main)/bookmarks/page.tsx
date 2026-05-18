import { PageLayout } from "@/components/layouts/PageLayout"
import { BookmarksPageContent } from "@/features/bookmarks/components/BookmarksPageContent"
import { getBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { ROUTES } from "@/lib/routes"

export default async function BookmarksPage() {
  const res = await getBookmarks()
  const bookmarks = res.data ?? []
  console.log(res)
  return (
    <PageLayout
      title="Bài học đã lưu"
      breadcrumbs={[
        { label: "Bài học", href: ROUTES.USER.LESSONS.LIST },
        { label: "Đã lưu" },
      ]}
    >
      <BookmarksPageContent bookmarks={bookmarks} />
    </PageLayout>
  )
}
