import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { BookmarkButton } from "@/features/bookmarks/components/BookmarkButton"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { Badge } from "@/components/ui/badge"
import { ROUTES } from "@/lib/routes"

export default async function DictationPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const lessonRes = await getLesson(Number(id))

  if (!lessonRes.data) {
    notFound()
  }

  const lesson = lessonRes.data
  const transcriptsRes = await getTranscripts(lesson.id)
  const transcripts = (transcriptsRes.data ?? [])
    .sort((a, b) => a.sequence - b.sequence)

  const bookmarksRes = await getBookmarks()
  const isBookmarked = (bookmarksRes.data ?? []).some((b) => b.lessionId === lesson.id)

  const decksRes = await getUserVocabularyDecks()
  const decks = decksRes.data ?? []

  return (
    <PageLayout
      title={lesson.title}
      breadcrumbs={[
        { label: "Chủ đề", href: ROUTES.USER.CATEGORIES },
        { label: "Bài học Nghe chép chính tả" },
      ]}
      actions={<BookmarkButton lessonId={lesson.id} isBookmarked={isBookmarked} />}
    >
      <LessonLayout
        videoUrl={lesson.videoUrl}
        duration={lesson.duration}
        transcripts={transcripts}
        lessonId={lesson.id}
        decks={decks}
        exercise={DictationArea}
      />
      <div className="mt-6 rounded-lg border p-4">
        <h3 className="mb-2 text-sm font-medium">Tiến độ Nghe chép chính tả</h3>
        <div className="flex items-center gap-3">
          <div className="h-2 flex-1 rounded-full bg-muted">
            <div className="h-2 w-2/5 rounded-full bg-primary" />
          </div>
          <Badge variant="secondary">2/{transcripts.length} câu đã hoàn thành</Badge>
        </div>
      </div>
    </PageLayout>
  )
}
