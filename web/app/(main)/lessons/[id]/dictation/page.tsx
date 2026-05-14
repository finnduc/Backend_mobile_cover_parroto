import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/features/lessons/components/user/LessonLayout"
import { DictationArea } from "@/features/lessons/components/user/DictationArea"
import { getLesson, getTranscripts } from "@/features/lessons/services/lessons-service"
import { ROUTES } from "@/lib/routes"

export default async function DictationPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const lessonRes = await getLesson(Number(id))

  if (lessonRes.error) {
    return <div className="p-6 text-center text-muted-foreground">{lessonRes.error.message}</div>
  }
  if (!lessonRes.data) {
    return <div className="p-6 text-center text-muted-foreground">Lesson not found</div>
  }

  const lesson = lessonRes.data
  const transcriptsRes = await getTranscripts(lesson.id)
  const transcripts = (transcriptsRes.data ?? [])
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
