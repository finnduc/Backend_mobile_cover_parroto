import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { TranscriptContent } from "@/features/lessons/components/admin/TranscriptContent"
import { mockTranscripts } from "@/features/lessons/mock-data"
import { ROUTES } from "@/lib/routes"

export default async function TranscriptsPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const data = mockTranscripts.filter((t) => t.lessonId === Number(id))

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <TranscriptContent lessonId={Number(id)} transcripts={data} />
    </AdminPageLayout>
  )
}
