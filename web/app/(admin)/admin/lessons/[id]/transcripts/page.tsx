import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { TranscriptContent } from "@/features/lessons/components/admin/TranscriptContent"
import { getAdminTranscripts } from "@/features/lessons/services/lessons.get"
import { ROUTES } from "@/lib/routes"

export default async function TranscriptsPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const res = await getAdminTranscripts(Number(id))

  if (res.error) {
    return <div className="py-12 text-center text-muted-foreground">{res.error.message}</div>
  }

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons">
      <TranscriptContent lessonId={Number(id)} transcripts={res.data ?? []} />
    </AdminPageLayout>
  )
}
