import { notFound } from "next/navigation"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { TranscriptContent } from "@/features/lessons/components/admin/TranscriptContent"
import { getAdminLesson, getAdminTranscripts } from "@/features/lessons/services/lessons.get"
import { ROUTES } from "@/lib/routes"

export default async function TranscriptsPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const lessonRes = await getAdminLesson(Number(id))
  if (!lessonRes.data) {
    notFound()
  }

  const res = await getAdminTranscripts(Number(id))

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <TranscriptContent lessonId={Number(id)} transcripts={res.data ?? []} />
    </AdminPageLayout>
  )
}
