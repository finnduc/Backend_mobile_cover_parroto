import { notFound } from "next/navigation"
import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { ShadowingArea } from "@/features/lessons/components/user/ShadowingArea"
import { BookmarkButton } from "@/features/bookmarks/components/BookmarkButton"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons.get"
import { getBookmarks } from "@/features/bookmarks/services/bookmarks.get"
import { getShadowingStatus } from "@/features/lessons/services/shadowing-status.get"
import { ROUTES } from "@/lib/routes"

export default async function ShadowingPage({
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

  const statusRes = await getShadowingStatus(lesson.id)
  const completedTranscriptIds = (statusRes.data ?? []).map((s) => s.transcriptId)
  const transcriptIdToIndex = new Map(transcripts.map((t, i) => [t.id, i]))
  const initialCompletedIds = completedTranscriptIds
    .map((tid) => transcriptIdToIndex.get(tid))
    .filter((i): i is number => i !== undefined)

  const bookmarksRes = await getBookmarks()
  const isBookmarked = (bookmarksRes.data ?? []).some((b) => b.lessionId === lesson.id)

  return (
    <PageLayout
      title={lesson.title}
      breadcrumbs={[
        { label: "Chủ đề", href: ROUTES.USER.CATEGORIES },
        { label: "Bài học Shadowing" },
      ]}
      actions={<BookmarkButton lessonId={lesson.id} isBookmarked={isBookmarked} />}
    >
      <LessonLayout
        videoUrl={lesson.videoUrl}
        duration={lesson.duration}
        transcripts={transcripts}
        initialCompletedIds={initialCompletedIds}
        lessonId={lesson.id}
        exercise={ShadowingArea}
      />
    </PageLayout>
  )
}
