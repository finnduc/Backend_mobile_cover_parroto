import { auth } from "@clerk/nextjs/server"
import { PageLayout } from "@/components/layouts/PageLayout"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Badge } from "@/components/ui/badge"
import { Card, CardContent } from "@/components/ui/card"
import { Skeleton } from "@/components/ui/skeleton"
import { getUserProfile } from "@/features/profile/services/profile.get"

export default async function ProfilePage() {
  const { userId } = await auth()
  if (!userId) return null

  const res = await getUserProfile()
  const profile = res.data

  return (
    <PageLayout
      title="Profile"
      breadcrumbs={[{ label: "Profile" }]}
    >
      {profile ? (
        <Card>
          <CardContent className="flex items-center gap-4 pt-6">
            <Avatar className="size-16">
              <AvatarImage src={profile.avatarUrl} />
              <AvatarFallback>{profile.name?.charAt(0) ?? "U"}</AvatarFallback>
            </Avatar>
            <div className="space-y-1">
              <h2 className="text-lg font-semibold">{profile.name}</h2>
              <p className="text-sm text-muted-foreground">{profile.email}</p>
              <div className="flex items-center gap-2">
                <Badge variant={profile.userRole === "admin" ? "default" : "secondary"}>
                  {profile.userRole}
                </Badge>
                {profile.createdAt && (
                  <span className="text-xs text-muted-foreground">
                    Member since {new Date(profile.createdAt).toLocaleDateString("vi-VN")}
                  </span>
                )}
              </div>
            </div>
          </CardContent>
        </Card>
      ) : res.error ? (
        <div className="rounded-lg border border-dashed p-8 text-center text-sm text-muted-foreground">
          Profile not available. {res.error.message}
        </div>
      ) : (
        <div className="flex items-center gap-4">
          <Skeleton className="size-16 rounded-full" />
          <div className="space-y-2">
            <Skeleton className="h-5 w-32" />
            <Skeleton className="h-4 w-48" />
          </div>
        </div>
      )}
    </PageLayout>
  )
}
