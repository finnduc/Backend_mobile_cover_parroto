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
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"
import { Badge } from "@/components/ui/badge"
import { cn } from "@/lib/utils"
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
  const tokens = phrase.trim().split(/\s+/).filter(Boolean)
  const [selectedTokens, setSelectedTokens] = useState<Set<number>>(new Set(tokens.map((_, i) => i)))
  const [customPhrase, setCustomPhrase] = useState(phrase)
  const [meaning, setMeaning] = useState("")
  const [exampleSentence, setExampleSentence] = useState(phrase)
  const [submitting, setSubmitting] = useState(false)
  const [selectedDeckId, setSelectedDeckId] = useState<string>("")
  const [showCreateForm, setShowCreateForm] = useState(false)
  const [newDeckName, setNewDeckName] = useState("")
  const [newDeckLevel, setNewDeckLevel] = useState("beginner")

  const toggleToken = (index: number) => {
    setSelectedTokens((prev) => {
      const next = new Set(prev)
      if (next.has(index)) {
        next.delete(index)
      } else {
        next.add(index)
      }
      const selected = tokens.filter((_, i) => next.has(i)).join(" ")
      setCustomPhrase(selected)
      return next
    })
  }

  const resetForm = () => {
    setSelectedTokens(new Set(tokens.map((_, i) => i)))
    setCustomPhrase(phrase)
    setExampleSentence(phrase)
    setSelectedDeckId("")
    setShowCreateForm(false)
    setNewDeckName("")
    setNewDeckLevel("beginner")
  }

  const handleDialogChange = (open: boolean) => {
    if (!open) {
      resetForm()
    }
    onOpenChange(open)
  }

  const handleSubmit = async () => {
    if (!selectedDeckId || !customPhrase.trim()) return
    setSubmitting(true)

    const deckId = Number(selectedDeckId)
    const res = await createVocabularyItem(deckId, {
      phrase: customPhrase.trim(),
      lessonId,
      transcriptId,
      normalizedPhrase: customPhrase.trim().toLowerCase(),
      meaning: meaning.trim(),
      exampleSentence: exampleSentence.trim(),
      note: "",
    })

    setSubmitting(false)

    if (!res.error) {
      const deck = decks.find((d) => d.id === deckId)
      toast.success(`Đã thêm vào "${deck?.name ?? "bộ từ vựng"}"`)
      handleDialogChange(false)
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

    const newDeck = deckRes.data!
    toast.success(`Đã tạo "${newDeck.name}". Vui lòng chọn lại bộ từ vựng.`)
  }

  return (
    <Dialog open={open} onOpenChange={handleDialogChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Thêm vào danh sách từ vựng</DialogTitle>
          <DialogDescription>
            Chọn từ trong transcript hoặc nhập từ tùy chỉnh
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
              <div className="space-y-2">
                <Label>Chọn từ trong câu transcript</Label>
                <div className="flex flex-wrap gap-1.5 rounded-lg border p-3">
                  {tokens.map((token, i) => (
                    <Badge
                      key={i}
                      variant={selectedTokens.has(i) ? "default" : "outline"}
                      className={cn(
                        "cursor-pointer select-none text-sm",
                        !selectedTokens.has(i) && "hover:bg-muted"
                      )}
                      onClick={() => toggleToken(i)}
                    >
                      {token}
                    </Badge>
                  ))}
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="add-custom-phrase">Từ / Cụm từ</Label>
                <Input
                  id="add-custom-phrase"
                  value={customPhrase}
                  onChange={(e) => {
                    setCustomPhrase(e.target.value)
                    setSelectedTokens(new Set())
                  }}
                  placeholder="Hoặc nhập từ tùy chỉnh"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="add-meaning">Ý nghĩa</Label>
                <Textarea
                  id="add-meaning"
                  value={meaning}
                  onChange={(e) => setMeaning(e.target.value)}
                  rows={2}
                  placeholder="Ý nghĩa của từ này"
                />
              </div>

              <div className="space-y-2"></div>

              <div className="space-y-2">
                <Label htmlFor="add-example-sentence">Câu ví dụ</Label>
                <Textarea
                  id="add-example-sentence"
                  value={exampleSentence}
                  onChange={(e) => setExampleSentence(e.target.value)}
                  rows={2}
                  placeholder="Ví dụ câu chứa từ này"
                />
              </div>

              <div className="border-t pt-4">
                {decks.length === 0 ? (
                  <div className="space-y-3 rounded-lg border p-4 text-center">
                    <p className="text-sm text-muted-foreground">Bạn chưa có bộ từ vựng nào</p>
                    <Button onClick={() => setShowCreateForm(true)}>
                      Tạo bộ từ vựng đầu tiên
                    </Button>
                  </div>
                ) : (
                  <>
                    <div className="space-y-1">
                      <Label htmlFor="deck-select">Bộ từ vựng</Label>
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

                    <div className="mt-3 flex items-center justify-between">
                      <Button
                        variant="link"
                        size="sm"
                        className="px-0"
                        onClick={() => setShowCreateForm(true)}
                      >
                        + Tạo bộ từ vựng mới
                      </Button>
                      <Button onClick={handleSubmit} disabled={submitting || !selectedDeckId || !customPhrase.trim()}>
                        {submitting ? "Đang thêm..." : "Thêm"}
                      </Button>
                    </div>
                  </>
                )}
              </div>
            </>
          )}
        </div>
      </DialogContent>
    </Dialog>
  )
}
