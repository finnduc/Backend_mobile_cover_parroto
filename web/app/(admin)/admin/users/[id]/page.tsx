import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { mockUsers } from "@/features/users/mock-data"
import { ROUTES } from "@/lib/routes"

export default async function UserDetailPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const user = mockUsers.find((u) => u.id === Number(id))

  if (!user) {
    return <div className="py-12 text-center text-muted-foreground">User not found</div>
  }

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
