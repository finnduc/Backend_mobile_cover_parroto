import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { mockLessons, mockTranscripts } from "@/features/lessons/mock-data"
import { ROUTES } from "@/lib/routes"

export default async function DictationPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const lesson = mockLessons.find((l) => l.id === Number(id))

  if (!lesson) {
    return <div className="p-6 text-center text-muted-foreground">Lesson not found</div>
  }

  const transcripts = mockTranscripts
    .filter((t) => t.lessonId === lesson.id)
    .sort((a, b) => a.sequence - b.sequence)

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
        title={lesson.title}
        duration={lesson.duration}
        transcripts={transcripts}
        exercise={<DictationArea />}
      />
    </PageLayout>
  )
}
