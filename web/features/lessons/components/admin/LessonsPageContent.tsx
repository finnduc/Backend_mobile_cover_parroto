"use client"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import { DataTable } from "@/components/common/DataTable"
import { PaginationBar } from "@/components/common/PaginationBar"
import { AdminPageLayout } from "@/components/layouts/AdminPageLayout"
import { ROUTES } from "@/lib/routes"
import type { Lesson } from "@/types/lessons.models"
import type { PaginatedMeta } from "@/types/base-response"
import type { Column } from "@/components/common/DataTable"

function getLessonColumns(): Column<Lesson>[] {
  return [
    { key: "id", header: "ID" },
    {
      key: "title",
      header: "Title",
      render: (l) => <span className="block max-w-[300px] truncate">{l.title}</span>,
    },
    { key: "level", header: "Level" },
    { key: "categoryId", header: "Category" },
    {
      key: "duration",
      header: "Duration",
      render: (l) => {
        const m = Math.floor(l.duration / 60)
        const s = l.duration % 60
        return `${m}:${s.toString().padStart(2, "0")}`
      },
    },
    {
      key: "actions",
      header: "",
      render: (l) => (
        <div className="flex justify-end gap-1">
          <Button size="xs" variant="outline" asChild>
            <Link href={ROUTES.ADMIN.LESSONS.DETAIL(String(l.id))}>View</Link>
          </Button>
          <Button size="xs" variant="outline" asChild>
            <Link href={ROUTES.ADMIN.LESSONS.EDIT(String(l.id))}>Edit</Link>
          </Button>
          <Button size="xs" variant="outline" asChild>
            <Link href={ROUTES.ADMIN.LESSONS.TRANSCRIPTS(String(l.id))}>Transcripts</Link>
          </Button>
        </div>
      ),
    },
  ]
}

export function LessonsPageContent({
  data,
  meta,
  limit,
}: {
  data: Lesson[]
  meta: PaginatedMeta
  limit: number
}) {
  return (
    <AdminPageLayout
      title="Lessons"
      actions={
        <Button asChild>
          <Link href={ROUTES.ADMIN.LESSONS.CREATE}>Create Lesson</Link>
        </Button>
      }
    >
      <DataTable columns={getLessonColumns()} data={data} />
      <PaginationBar
        currentPage={meta.page}
        totalPages={meta.totalPages}
        baseUrl={ROUTES.ADMIN.LESSONS.LIST}
        searchParams={new URLSearchParams({ limit: String(limit) })}
      />
    </AdminPageLayout>
  )
}
