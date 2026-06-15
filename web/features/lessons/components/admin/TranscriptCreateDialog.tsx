"use client"

import { useState } from "react"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Loader2 } from "lucide-react"
import { createAdminTranscript } from "@/features/lessons/services/transcripts.action"
import { toast } from "sonner"
import type { Transcript } from "@/types/lessons.models"

export function TranscriptCreateDialog({
  lessonId,
  nextSequence,
  open,
  onOpenChange,
  onCreated,
}: {
  lessonId: number
  nextSequence: number
  open: boolean
  onOpenChange: (open: boolean) => void
  onCreated: (transcript: Transcript) => void
}) {
  const [saving, setSaving] = useState(false)
  const [form, setForm] = useState({
    content: "",
    phonetic: "",
    vietnamese: "",
    startTimestamp: 0,
    endTimestamp: 0,
  })

  const set = (field: string, value: string | number) =>
    setForm((prev) => ({ ...prev, [field]: value }))

  const handleSubmit = async () => {
    if (!form.content.trim()) {
      toast.error("Content is required")
      return
    }
    setSaving(true)
    const res = await createAdminTranscript({
      lessonId,
      sequence: nextSequence,
      content: form.content,
      phonetic: form.phonetic,
      vietnamese: form.vietnamese,
      startTimestamp: form.startTimestamp,
      endTimestamp: form.endTimestamp,
    })
    setSaving(false)
    if (res.error) {
      toast.error(res.error.message)
    } else if (res.data) {
      toast.success("Segment created")
      onCreated(res.data)
      onOpenChange(false)
      setForm({ content: "", phonetic: "", vietnamese: "", startTimestamp: 0, endTimestamp: 0 })
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>Add Transcript Segment</DialogTitle>
          <DialogDescription>
            Fill in the fields below to create a new segment.
          </DialogDescription>
        </DialogHeader>

        <div className="grid gap-4 py-2">
          <div className="grid gap-1.5">
            <Label htmlFor="t-content">Content *</Label>
            <Input
              id="t-content"
              value={form.content}
              onChange={(e) => set("content", e.target.value)}
              placeholder="e.g. Hello, how are you?"
            />
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div className="grid gap-1.5">
              <Label htmlFor="t-phonetic">Phonetic</Label>
              <Input
                id="t-phonetic"
                value={form.phonetic}
                onChange={(e) => set("phonetic", e.target.value)}
                placeholder="e.g. həˈloʊ"
              />
            </div>
            <div className="grid gap-1.5">
              <Label htmlFor="t-vietnamese">Vietnamese</Label>
              <Input
                id="t-vietnamese"
                value={form.vietnamese}
                onChange={(e) => set("vietnamese", e.target.value)}
                placeholder="e.g. Xin chào"
              />
            </div>
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div className="grid gap-1.5">
              <Label htmlFor="t-start">Start Timestamp</Label>
              <Input
                id="t-start"
                type="number"
                value={form.startTimestamp}
                onChange={(e) => set("startTimestamp", Number(e.target.value))}
              />
            </div>
            <div className="grid gap-1.5">
              <Label htmlFor="t-end">End Timestamp</Label>
              <Input
                id="t-end"
                type="number"
                value={form.endTimestamp}
                onChange={(e) => set("endTimestamp", Number(e.target.value))}
              />
            </div>
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>
            Cancel
          </Button>
          <Button onClick={handleSubmit} disabled={saving || !form.content.trim()}>
            {saving && <Loader2 className="mr-1 size-4 animate-spin" />}
            Create
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
