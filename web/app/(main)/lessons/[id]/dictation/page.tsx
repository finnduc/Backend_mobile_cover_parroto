import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { getDictationStatus } from "@/features/lessons/services/dictation-status.get"
// import { getTranscriptBookmarks } from "@/features/bookmarks/services/bookmarks.get"
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

  const decksRes = await getUserVocabularyDecks()
  const decks = decksRes.data ?? []

  // const bookmarksRes = await getTranscriptBookmarks(lesson.id)
  // const bookmarks = bookmarksRes.data ?? []

  return (
    <PageLayout
      title={lesson.title}
      breadcrumbs={[
        { label: "Chủ đề", href: ROUTES.USER.CATEGORIES },
        { label: "Bài học Nghe chép chính tả" },
      ]}
    >
      <LessonLayout
        videoUrl={lesson.videoUrl}
        transcripts={transcripts}
        lessonId={lesson.id}
        decks={decks}
        initialCompletedIds={initialCompletedIds}
        // bookmarks={bookmarks}
      >
        <DictationArea
          transcripts={transcripts}
          initialCompletedIds={initialCompletedIds}
          lessonId={lesson.id}
        />
      </LessonLayout>
    </PageLayout>
  )
}
