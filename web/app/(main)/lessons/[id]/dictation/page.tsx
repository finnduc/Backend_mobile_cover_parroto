import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { BookmarkButton } from "@/features/bookmarks/components/BookmarkButton"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getTranscriptBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { Badge } from "@/components/ui/badge"
import { getDictationStatus } from "@/features/lessons/services/dictation-status.get"
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

  const statusRes = await getDictationStatus(lesson.id)
  const completedTranscriptIds = (statusRes.data ?? []).map((s) => s.transcriptId)
  const transcriptIdToIndex = new Map(transcripts.map((t, i) => [t.id, i]))
  const initialCompletedIds = completedTranscriptIds
    .map((tid) => transcriptIdToIndex.get(tid))
    .filter((i): i is number => i !== undefined)

  const bookmarksRes = await getTranscriptBookmarks(lesson.id)
  const bookmarks = bookmarksRes.data ?? []
  const firstTranscript = transcripts[0]
  const firstBookmark = bookmarks.find(
    (b) => firstTranscript && b.transcriptId === firstTranscript.id
  )
  const isBookmarked = bookmarks.length > 0

  const decksRes = await getUserVocabularyDecks()
  const decks = decksRes.data ?? []

  return (
    <PageLayout
      title={lesson.title}
      breadcrumbs={[
        { label: "Chu de", href: ROUTES.USER.CATEGORIES },
        { label: "Bai hoc Nghe chep chinh ta" },
      ]}
      actions={
        firstTranscript ? (
          <BookmarkButton
            transcriptId={firstTranscript.id}
            bookmarkId={firstBookmark?.id ?? null}
            isBookmarked={isBookmarked}
          />
        ) : undefined
      }
    >
      <LessonLayout
        videoUrl={lesson.videoUrl}
        duration={lesson.duration}
        transcripts={transcripts}
        lessonId={lesson.id}
        decks={decks}
        initialCompletedIds={initialCompletedIds}
        exercise={DictationArea}
      />
      <div className="mt-6 rounded-lg border p-4">
        <h3 className="mb-2 text-sm font-medium">Tien do Nghe chep chinh ta</h3>
        <div className="flex items-center gap-3">
          <div className="h-2 flex-1 rounded-full bg-muted">
            <div className="h-2 w-2/5 rounded-full bg-primary" />
          </div>
          <Badge variant="secondary">2/{transcripts.length} cau da hoan thanh</Badge>
        </div>
      </div>
    </PageLayout>
  )
}
