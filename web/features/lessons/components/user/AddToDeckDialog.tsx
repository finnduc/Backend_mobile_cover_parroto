"use client"

import { useState } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { toast } from "sonner"
import { createVocabularyDeck, createVocabularyItem } from "@/features/vocabulary/services/vocabulary.action"
import type { VocabularyDeck } from "@/types/vocabulary.models"

export function AddToDeckDialog({
  open,
  onOpenChange,
  phrase,
  lessonId,
  transcriptId,
  decks,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  phrase: string
  lessonId: number
  transcriptId: number
  decks: VocabularyDeck[]
}) {
  const [selectedDeckId, setSelectedDeckId] = useState<string>("")
  const [submitting, setSubmitting] = useState(false)
  const [showCreateForm, setShowCreateForm] = useState(false)
  const [newDeckName, setNewDeckName] = useState("")
  const [newDeckLevel, setNewDeckLevel] = useState("beginner")

  const handleSubmit = async () => {
    if (!selectedDeckId) return
    setSubmitting(true)

    const deckId = Number(selectedDeckId)
    const res = await createVocabularyItem(deckId, {
      phrase,
      lessonId,
      transcriptId,
      normalizedPhrase: phrase,
      meaning: "",
      exampleSentence: "",
      note: "",
    })

    setSubmitting(false)

    if (!res.error) {
      const deck = decks.find((d) => d.id === deckId)
      toast.success(`Đã thêm vào "${deck?.name ?? "bộ từ vựng"}"`)
      onOpenChange(false)
    } else {
      toast.error(res.error.message)
    }
  }

  const handleCreateAndSubmit = async () => {
    if (!newDeckName.trim()) return
    setSubmitting(true)

    const deckRes = await createVocabularyDeck({
      name: newDeckName.trim(),
      description: "",
      thumbnailUrl: "",
      level: newDeckLevel,
      categoryId: null,
    })

    if (deckRes.error) {
      toast.error(deckRes.error.message)
      setSubmitting(false)
      return
    }

    setShowCreateForm(false)
    setSubmitting(false)

    // Server action revalidates cache; dialog closes, user reopens to see new deck
    const newDeck = deckRes.data!
    toast.success(`Đã tạo "${newDeck.name}". Vui lòng chọn lại bộ từ vựng.`)
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Thêm vào danh sách từ vựng</DialogTitle>
          <DialogDescription>
            Thêm &ldquo;{phrase}&rdquo; vào bộ từ vựng của bạn
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {showCreateForm ? (
            <div className="space-y-3 rounded-lg border p-4">
              <div className="space-y-1">
                <Label htmlFor="deck-name">Tên bộ từ vựng</Label>
                <Input
                  id="deck-name"
                  value={newDeckName}
                  onChange={(e) => setNewDeckName(e.target.value)}
                  placeholder="VD: Từ vựng bài 1"
                />
              </div>
              <div className="space-y-1">
                <Label htmlFor="deck-level">Cấp độ</Label>
                <Select value={newDeckLevel} onValueChange={setNewDeckLevel}>
                  <SelectTrigger id="deck-level">
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="beginner">Cơ bản</SelectItem>
                    <SelectItem value="intermediate">Trung cấp</SelectItem>
                    <SelectItem value="advanced">Nâng cao</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="flex gap-2">
                <Button onClick={handleCreateAndSubmit} disabled={submitting || !newDeckName.trim()}>
                  {submitting ? "Đang tạo..." : "Tạo"}
                </Button>
                <Button variant="ghost" onClick={() => setShowCreateForm(false)}>
                  Hủy
                </Button>
              </div>
            </div>
          ) : (
            <>
              {decks.length === 0 ? (
                <div className="space-y-3 rounded-lg border p-4 text-center">
                  <p className="text-sm text-muted-foreground">Bạn chưa có bộ từ vựng nào</p>
                  <Button onClick={() => setShowCreateForm(true)}>
                    Tạo bộ từ vựng đầu tiên
                  </Button>
                </div>
              ) : (
                <div className="space-y-1">
                  <Label htmlFor="deck-select">Chọn bộ từ vựng</Label>
                  <Select value={selectedDeckId} onValueChange={setSelectedDeckId}>
                    <SelectTrigger id="deck-select">
                      <SelectValue placeholder="Chọn bộ từ vựng..." />
                    </SelectTrigger>
                    <SelectContent>
                      {decks.map((deck) => (
                        <SelectItem key={deck.id} value={String(deck.id)}>
                          {deck.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              )}

              {decks.length > 0 && (
                <div className="flex items-center justify-between">
                  <Button
                    variant="link"
                    size="sm"
                    className="px-0"
                    onClick={() => setShowCreateForm(true)}
                  >
                    + Tạo bộ từ vựng mới
                  </Button>
                  <Button onClick={handleSubmit} disabled={submitting || !selectedDeckId}>
                    {submitting ? "Đang thêm..." : "Thêm"}
                  </Button>
                </div>
              )}
            </>
          )}
        </div>
      </DialogContent>
    </Dialog>
  )
}
