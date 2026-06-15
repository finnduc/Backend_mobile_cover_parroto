"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { DataTable } from "@/components/common/DataTable"
import { Plus, Trash2, Save, Upload, Download } from "lucide-react"
import {
  updateAdminTranscript,
  deleteAdminTranscript,
} from "@/features/lessons/services/transcripts.action"
import { TranscriptBulkDialog } from "@/features/lessons/components/admin/TranscriptBulkDialog"
import { TranscriptCreateDialog } from "@/features/lessons/components/admin/TranscriptCreateDialog"
import { toast } from "sonner"
import type { Transcript } from "@/types/lessons.models"
import type { Column } from "@/components/common/DataTable"

export function TranscriptContent({
  lessonId,
  transcripts: initialTranscripts,
}: {
  lessonId: number
  transcripts: Transcript[]
}) {
  const [transcripts, setTranscripts] = useState(initialTranscripts)

  const [importOpen, setImportOpen] = useState(false)
  const [createOpen, setCreateOpen] = useState(false)

  const exportJson = () => {
    const data = transcripts.map((t) => ({
      sequence: t.sequence,
      content: t.content,
      phonetic: t.phonetic,
      vietnamese: t.vietnamese,
      startTimestamp: t.startTimestamp,
      endTimestamp: t.endTimestamp,
    }))
    const blob = new Blob([JSON.stringify(data, null, 2)], { type: "application/json" })
    const url = URL.createObjectURL(blob)
    const a = document.createElement("a")
    a.href = url
    a.download = `lesson-${lessonId}-transcripts.json`
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
    toast.success("Transcripts exported as JSON.")
  }

  const nextSequence = Math.max(0, ...transcripts.map((t) => t.sequence)) + 1

  const remove = async (id: number) => {
    const res = await deleteAdminTranscript(id, lessonId)
    if (res.error) {
      toast.error(res.error.message)
    } else {
      setTranscripts((prev) => prev.filter((t) => t.id !== id))
    }
  }

  const update = (id: number, field: keyof Transcript, value: string | number) => {
    setTranscripts((prev) =>
      prev.map((t) => (t.id === id ? { ...t, [field]: value } : t))
    )
  }

  const saveRow = async (t: Transcript) => {
    const res = await updateAdminTranscript(t.id, {
      lessonId: t.lessonId,
      sequence: t.sequence,
      content: t.content,
      phonetic: t.phonetic,
      vietnamese: t.vietnamese,
      startTimestamp: t.startTimestamp,
      endTimestamp: t.endTimestamp,
    })
    if (res.error) {
      toast.error(res.error.message)
    }
  }

  const columns: Column<Transcript>[] = [
    { key: "sequence", header: "Seq" },
    {
      key: "content",
      header: "Content",
      render: (t) => (
        <Input
          value={t.content}
          onChange={(e) => update(t.id, "content", e.target.value)}
          onBlur={() => saveRow(t)}
          className="h-8 text-sm"
        />
      ),
    },
    {
      key: "phonetic",
      header: "Phonetic",
      render: (t) => (
        <Input
          value={t.phonetic}
          onChange={(e) => update(t.id, "phonetic", e.target.value)}
          onBlur={() => saveRow(t)}
          className="h-8 text-sm"
        />
      ),
    },
    {
      key: "vietnamese",
      header: "Vietnamese",
      render: (t) => (
        <Input
          value={t.vietnamese}
          onChange={(e) => update(t.id, "vietnamese", e.target.value)}
          onBlur={() => saveRow(t)}
          className="h-8 text-sm"
        />
      ),
    },
    {
      key: "startTimestamp",
      header: "Start",
      render: (t) => (
        <Input
          type="number"
          value={t.startTimestamp}
          onChange={(e) => update(t.id, "startTimestamp", Number(e.target.value))}
          onBlur={() => saveRow(t)}
          className="h-8 w-full text-sm"
        />
      ),
    },
    {
      key: "endTimestamp",
      header: "End",
      render: (t) => (
        <Input
          type="number"
          value={t.endTimestamp}
          onChange={(e) => update(t.id, "endTimestamp", Number(e.target.value))}
          onBlur={() => saveRow(t)}
          className="h-8 w-full text-sm"
        />
      ),
    },
    {
      key: "actions",
      header: "",
      render: (t) => (
        <div className="flex items-center gap-1">
          <Button size="icon" variant="ghost" onClick={() => saveRow(t)}>
            <Save className="size-4 text-primary" />
          </Button>
          <Button size="icon" variant="ghost" onClick={() => remove(t.id)}>
            <Trash2 className="size-4 text-destructive" />
          </Button>
        </div>
      ),
    },
  ]

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h3 className="text-lg font-semibold">Transcripts for Lesson #{lessonId}</h3>
        <div className="flex items-center gap-2">
          <Button size="sm" variant="outline" onClick={() => setImportOpen(true)}>
            <Upload className="mr-1 size-4" />
            Import JSON
          </Button>
          <Button size="sm" variant="outline" onClick={exportJson}>
            <Download className="mr-1 size-4" />
            Export JSON
          </Button>
          <Button size="sm" onClick={() => setCreateOpen(true)}>
            <Plus className="mr-1 size-4" />
            Add Segment
          </Button>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Transcripts for Lesson #{lessonId}</CardTitle>
        </CardHeader>
        <CardContent>
          <DataTable columns={columns} data={transcripts} emptyMessage={"No transcripts yet. Click \"Add Segment\" to add one."} />
        </CardContent>
      </Card>

      <TranscriptBulkDialog
        lessonId={lessonId}
        open={importOpen}
        onOpenChange={setImportOpen}
      />

      <TranscriptCreateDialog
        lessonId={lessonId}
        nextSequence={nextSequence}
        open={createOpen}
        onOpenChange={setCreateOpen}
        onCreated={(t) => setTranscripts((prev) => [...prev, t])}
      />
    </div>
  )
}
