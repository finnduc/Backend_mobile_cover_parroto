import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { getAdminLesson } from "@/features/lessons/services/lessons-service"
import { ROUTES } from "@/lib/routes"

export default async function LessonDetailPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const res = await getAdminLesson(Number(id))

  if (res.error) {
    return <div className="py-12 text-center text-muted-foreground">{res.error.message}</div>
  }
  if (!res.data) {
    return <div className="py-12 text-center text-muted-foreground">Lesson not found</div>
  }

  const lesson = res.data

  return (
    <AdminPageLayout backHref={ROUTES.ADMIN.LESSONS.LIST} backLabel="Back to Lessons" maxWidth="narrow">
      <Card>
        <CardHeader>
          <CardTitle>{lesson.title}</CardTitle>
        </CardHeader>
        <CardContent className="space-y-3 text-sm">
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">ID:</span>
            <span>{lesson.id}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Level:</span>
            <span>{lesson.level}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Duration:</span>
            <span>{Math.floor(lesson.duration / 60)}:{String(lesson.duration % 60).padStart(2, "0")}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Category:</span>
            <span>{lesson.categoryId}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Description:</span>
            <span className="text-muted-foreground">{lesson.description || "\u2014"}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Thumbnail:</span>
            <span className="break-all text-muted-foreground">{lesson.thumbnailUrl || "\u2014"}</span>
          </div>
          <div className="flex gap-2">
            <span className="w-24 font-medium text-muted-foreground">Video:</span>
            <span className="break-all text-muted-foreground">{lesson.videoUrl || "\u2014"}</span>
          </div>
        </CardContent>
      </Card>

      <div className="flex gap-2">
        <Button asChild variant="outline">
          <Link href={ROUTES.ADMIN.LESSONS.EDIT(String(lesson.id))}>Edit</Link>
        </Button>
        <Button asChild variant="outline">
          <Link href={ROUTES.ADMIN.LESSONS.TRANSCRIPTS(String(lesson.id))}>Manage Transcripts</Link>
        </Button>
      </div>
    </AdminPageLayout>
  )
}
