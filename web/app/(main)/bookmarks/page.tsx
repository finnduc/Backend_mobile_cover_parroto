import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"

export default async function BookmarksPage() {
  const { userId } = await auth()
  if (!userId) return null

  return (
    <PageLayout
      title="Bai hoc da luu"
      breadcrumbs={[
        { label: "Bai hoc", href: "/lessons" },
        { label: "Da luu" },
      ]}
    >
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <p className="text-lg font-medium">Chon mot bai hoc de xem bookmark</p>
        <p className="text-sm">
          Tim transcript bookmark trong trang chi tiet bai hoc
        </p>
      </div>
    </PageLayout>
  )
}
