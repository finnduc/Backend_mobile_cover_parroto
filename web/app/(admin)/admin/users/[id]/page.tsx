import { notFound } from "next/navigation"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { getAdminUser } from "@/features/users/services/users.get"
import { ROUTES } from "@/lib/routes"

export default async function UserDetailPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const res = await getAdminUser(Number(id))

  if (!res.data) {
    notFound()
  }

  const user = res.data

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.USERS.LIST} backLabel="Back to Users" maxWidth="narrow">
      <Card>
        <CardHeader>
          <CardTitle>User Detail</CardTitle>
        </CardHeader>
        <CardContent className="space-y-3 text-sm">
          <div className="flex gap-2">
            <span className="w-16 font-medium text-muted-foreground">ID:</span>
            <span>{user.id}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-16 font-medium text-muted-foreground">Name:</span>
            <span>{user.name}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-16 font-medium text-muted-foreground">Email:</span>
            <span>{user.email}</span>
          </div>
        </CardContent>
      </Card>
    </AdminPageLayout>
  )
}
