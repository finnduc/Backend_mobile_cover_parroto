"use client"

import { useState } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"

export function BookmarkNoteDialog({
  open,
  onOpenChange,
  initialNote,
  onSave,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  initialNote?: string
  onSave: (note: string) => void
}) {
  const [note, setNote] = useState(initialNote ?? "")

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-sm">
        <DialogHeader>
          <DialogTitle>{initialNote ? "Edit Note" : "Add Note"}</DialogTitle>
        </DialogHeader>
        <div className="space-y-2">
          <Label htmlFor="bookmark-note">Note</Label>
          <Textarea
            id="bookmark-note"
            value={note}
            onChange={(e) => setNote(e.target.value)}
            placeholder="Add your notes here..."
            rows={3}
          />
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>
            Cancel
          </Button>
          <Button onClick={() => {
            onSave(note)
            onOpenChange(false)
          }}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
