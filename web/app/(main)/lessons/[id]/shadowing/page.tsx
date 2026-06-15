import { notFound } from "next/navigation"
import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { ShadowingArea } from "@/features/lessons/components/user/ShadowingArea"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getShadowingStatus } from "@/features/lessons/services/shadowing-status.get"
import { getPronunciationProgressDetail } from "@/features/pronunciation/services/pronunciation.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"
import type { PronunciationAttempt } from "@/types/pronunciation.models"
import type { VocabularyDeck } from "@/types/vocabulary.models"
// import { getTranscriptBookmarks } from "@/features/bookmarks/services/bookmarks.get"

export default async function ShadowingPage({
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
  const pronunciationScores = new Map<number, PronunciationAttempt>()
  let decks: VocabularyDeck[] = []

  if (userId) {
    const statusRes = await getShadowingStatus(lesson.id)
    const completedTranscriptIds = (statusRes.data ?? []).map((s) => s.transcriptId)
    initialCompletedIds = completedTranscriptIds
      .map((tid) => transcriptIdToIndex.get(tid))
      .filter((i): i is number => i !== undefined)

    const progressRes = await getPronunciationProgressDetail(lesson.id)
    ;(progressRes.data ?? []).forEach((p) => {
      const idx = transcriptIdToIndex.get(p.transcriptId)
      if (idx !== undefined) {
        pronunciationScores.set(idx, {
          text: transcripts[idx]?.content ?? "",
          overallScore: p.overallScore ?? p.bestScore ?? 0,
          scores: p.scores ?? { accuracy: 0, fluency: 0, completeness: 0, prosody: 0 },
          feedback: p.feedback ?? "",
          words: p.words ?? [],
          attempt: {
            id: p.bestAttemptId ?? 0,
            userId: p.userId,
            lessonId: p.lessonId,
            transcriptId: p.transcriptId,
            createdAt: p.createdAt,
          },
        })
      }
    })

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
        { label: "Bài học Shadowing" },
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
        <ShadowingArea
          transcripts={transcripts}
          initialCompletedIds={initialCompletedIds}
          lessonId={lesson.id}
          isAuthenticated={!!userId}
          pronunciationScores={pronunciationScores}
        />
      </LessonLayout>
    </PageLayout>
  )
}
