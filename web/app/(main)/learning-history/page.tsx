import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import {
  getFinishedLessons,
  getUnfinishedLessons,
  getLearningHistorySummary,
} from "@/features/learning-history/services/learning-history.get"
import { LearningHistoryPageContent } from "@/features/learning-history/components/LearningHistoryPageContent"

export default async function LearningHistoryPage() {
  const { userId } = await auth()
  if (!userId) return null

  const [finishedRes, unfinishedRes, summaryRes] = await Promise.all([
    getFinishedLessons(),
    getUnfinishedLessons(),
    getLearningHistorySummary(),
  ])

  return (
    <PageLayout
      title="Lich su hoc tap"
      breadcrumbs={[
        { label: "Bai hoc", href: "/lessons" },
        { label: "Lich su" },
      ]}
    >
      <LearningHistoryPageContent
        finished={finishedRes.data ?? []}
        unfinished={unfinishedRes.data ?? []}
        summary={summaryRes.data ?? { completedCount: 0, unfinishedCount: 0 }}
      />
    </PageLayout>
  )
}
