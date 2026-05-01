import { PageLayout } from "@/components/layouts/PageLayout"
import { LessonLayout } from "@/components/layouts/LessonLayout"
import { ShadowingArea } from "@/features/lessons/components/user/ShadowingArea"
import { TranscriptLine } from "@/components/common/TranscriptLine"
import { mockLessons, mockTranscripts } from "@/features/lessons/mock-data"
import { ROUTES } from "@/lib/routes"

export default async function ShadowingPage({
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
        { label: "Bài học Shadowing" },
      ]}
    >
      <LessonLayout
        title={lesson.title}
        duration={lesson.duration}
        transcript={transcripts.map((seg) => (
          <TranscriptLine
            key={seg.id}
            text={seg.content}
            timestamp={`${Math.floor(seg.startTimestamp / 60)}:${(seg.startTimestamp % 60).toString().padStart(2, "0")}`}
          />
        ))}
        exercise={<ShadowingArea />}
      />
    </PageLayout>
  )
}
