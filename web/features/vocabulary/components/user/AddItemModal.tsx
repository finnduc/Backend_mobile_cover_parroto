"use client"

import { useState } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"
import { toast } from "sonner"
import { createVocabularyItem } from "@/features/vocabulary/services/vocabulary.action"

export function AddItemModal({
  open,
  onOpenChange,
  deckId,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  deckId: number
}) {
  const [phrase, setPhrase] = useState("")
  const [normalizedPhrase, setNormalizedPhrase] = useState("")
  const [meaning, setMeaning] = useState("")
  const [exampleSentence, setExampleSentence] = useState("")
  const [note, setNote] = useState("")
  const [submitting, setSubmitting] = useState(false)

  const resetForm = () => {
    setPhrase("")
    setNormalizedPhrase("")
    setMeaning("")
    setExampleSentence("")
    setNote("")
  }

  const handleSubmit = async () => {
    if (!phrase.trim() || !meaning.trim()) return
    setSubmitting(true)

    const res = await createVocabularyItem(deckId, {
      phrase: phrase.trim(),
      normalizedPhrase: normalizedPhrase.trim() || phrase.trim().toLowerCase(),
      meaning: meaning.trim(),
      exampleSentence: exampleSentence.trim(),
      note: note.trim(),
      lessonId: null,
      transcriptId: null,
    })

    setSubmitting(false)

    if (!res.error) {
      toast.success("Đã thêm từ vựng")
      resetForm()
      onOpenChange(false)
    } else {
      toast.error(res.error.message)
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Thêm từ vựng</DialogTitle>
          <DialogDescription>
            Thêm một từ hoặc cụm từ mới vào bộ từ vựng
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          <div className="grid grid-cols-2 gap-3">
            <div className="space-y-2">
              <Label htmlFor="modal-item-phrase">Từ / Cụm từ</Label>
              <Input
                id="modal-item-phrase"
                value={phrase}
                onChange={(e) => setPhrase(e.target.value)}
                placeholder="hello"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="modal-item-normalized">Dạng chuẩn</Label>
              <Input
                id="modal-item-normalized"
                value={normalizedPhrase}
                onChange={(e) => setNormalizedPhrase(e.target.value)}
                placeholder="hello"
              />
            </div>
          </div>
          <div className="space-y-2">
            <Label htmlFor="modal-item-meaning">Nghĩa</Label>
            <Input
              id="modal-item-meaning"
              value={meaning}
              onChange={(e) => setMeaning(e.target.value)}
              placeholder="xin chào"
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="modal-item-example">Câu ví dụ</Label>
            <Textarea
              id="modal-item-example"
              value={exampleSentence}
              onChange={(e) => setExampleSentence(e.target.value)}
              placeholder="Hello, how are you?"
              rows={2}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="modal-item-note">Ghi chú</Label>
            <Input
              id="modal-item-note"
              value={note}
              onChange={(e) => setNote(e.target.value)}
              placeholder="Ghi chú thêm (tùy chọn)"
            />
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>
            Hủy
          </Button>
          <Button
            onClick={handleSubmit}
            disabled={submitting || !phrase.trim() || !meaning.trim()}
          >
            {submitting ? "Đang thêm..." : "Thêm"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
