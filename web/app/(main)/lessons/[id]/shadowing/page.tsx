import { notFound } from "next/navigation"
import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { ShadowingLayout } from "@/features/lessons/components/user/shadowing/ShadowingLayout"
import { ShadowingArea } from "@/features/lessons/components/user/shadowing/ShadowingArea"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getShadowingStatus } from "@/features/lessons/services/shadowing-status.get"
import { getUserVocabularyDecks } from "@/features/vocabulary/services/vocabulary.get"
import { ROUTES } from "@/lib/routes"
import type { VocabularyDeck } from "@/types/vocabulary.models"

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
  let completedTranscriptIds: number[] = []
  let initialCompletedIndices: number[] = []
  let decks: VocabularyDeck[] = []

  if (userId) {
    const statusRes = await getShadowingStatus(lesson.id)
    completedTranscriptIds = (statusRes.data ?? []).map((s) => s.transcriptId)
    initialCompletedIndices = completedTranscriptIds
      .map((tid) => transcriptIdToIndex.get(tid))
      .filter((i): i is number => i !== undefined)

    const decksRes = await getUserVocabularyDecks()
    decks = decksRes.data ?? []
  }

  return (
    <PageLayout
      title={lesson.title}
      breadcrumbs={[
        { label: "Chủ đề", href: ROUTES.USER.CATEGORIES },
        { label: "Bài học Shadowing" },
      ]}
    >
      <ShadowingLayout
        videoUrl={lesson.videoUrl}
        transcripts={transcripts}
        lessonId={lesson.id}
        decks={decks}
        completedTranscriptIds={completedTranscriptIds}
      >
        <ShadowingArea
          transcripts={transcripts}
          initialCompletedIds={initialCompletedIndices}
          lessonId={lesson.id}
          isAuthenticated={!!userId}
        />
      </ShadowingLayout>
    </PageLayout>
  )
}
