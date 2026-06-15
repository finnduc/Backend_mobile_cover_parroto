import { notFound } from "next/navigation"
import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { getDictationStatus } from "@/features/lessons/services/dictation-status.get"
import type { VocabularyDeck } from "@/types/vocabulary.models"
// import { getTranscriptBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { ROUTES } from "@/lib/routes"

export default async function DictationPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const { userId } = await auth()
  const lessonRes = await getLesson(Number(id))

  if (!lessonRes.data) {
    notFound()
  }

  const lesson = lessonRes.data
  const transcriptsRes = await getTranscripts(lesson.id)
  const transcripts = (transcriptsRes.data ?? [])
    .sort((a, b) => a.sequence - b.sequence)

  const transcriptIdToIndex = new Map(transcripts.map((t, i) => [t.id, i]))
  let initialCompletedIds: number[] = []
  let decks: VocabularyDeck[] = []

  if (userId) {
    const statusRes = await getDictationStatus(lesson.id)
    const completedTranscriptIds = (statusRes.data ?? []).map((s) => s.transcriptId)
    initialCompletedIds = completedTranscriptIds
      .map((tid) => transcriptIdToIndex.get(tid))
      .filter((i): i is number => i !== undefined)

    const decksRes = await getUserVocabularyDecks()
    decks = decksRes.data ?? []
  }

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
          isAuthenticated={!!userId}
        />
      </LessonLayout>
    </PageLayout>
  )
}
