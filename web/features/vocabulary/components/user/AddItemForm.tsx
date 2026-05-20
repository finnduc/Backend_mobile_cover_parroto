"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"
import { Card, CardContent } from "@/components/ui/card"
import { Plus } from "lucide-react"

export function AddItemForm({
  onSubmit,
}: {
  onSubmit: (values: {
    phrase: string
    normalizedPhrase: string
    meaning: string
    exampleSentence: string
    note: string
  }) => void
}) {
  const [phrase, setPhrase] = useState("")
  const [normalizedPhrase, setNormalizedPhrase] = useState("")
  const [meaning, setMeaning] = useState("")
  const [exampleSentence, setExampleSentence] = useState("")
  const [note, setNote] = useState("")

  const handleSubmit = () => {
    if (!phrase.trim() || !meaning.trim()) return
    onSubmit({
      phrase: phrase.trim(),
      normalizedPhrase: normalizedPhrase.trim() || phrase.trim().toLowerCase(),
      meaning: meaning.trim(),
      exampleSentence: exampleSentence.trim(),
      note: note.trim(),
    })
    setPhrase("")
    setNormalizedPhrase("")
    setMeaning("")
    setExampleSentence("")
    setNote("")
  }

  return (
    <Card size="sm">
      <CardContent>
        <div className="space-y-3">
          <div className="grid grid-cols-2 gap-3">
            <div className="space-y-2">
              <Label htmlFor="item-phrase">Từ / Cụm từ</Label>
              <Input
                id="item-phrase"
                value={phrase}
                onChange={(e) => setPhrase(e.target.value)}
                placeholder="hello"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="item-normalized">Dạng chuẩn</Label>
              <Input
                id="item-normalized"
                value={normalizedPhrase}
                onChange={(e) => setNormalizedPhrase(e.target.value)}
                placeholder="hello"
              />
            </div>
          </div>
          <div className="space-y-2">
            <Label htmlFor="item-meaning">Nghĩa</Label>
            <Input
              id="item-meaning"
              value={meaning}
              onChange={(e) => setMeaning(e.target.value)}
              placeholder="xin chào"
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="item-example">Câu ví dụ</Label>
            <Textarea
              id="item-example"
              value={exampleSentence}
              onChange={(e) => setExampleSentence(e.target.value)}
              placeholder="Hello, how are you?"
              rows={2}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="item-note">Ghi chú</Label>
            <Input
              id="item-note"
              value={note}
              onChange={(e) => setNote(e.target.value)}
              placeholder="Ghi chú thêm (tùy chọn)"
            />
          </div>
          <Button
            onClick={handleSubmit}
            disabled={!phrase.trim() || !meaning.trim()}
            size="sm"
          >
            <Plus className="size-4" />
            Thêm từ vựng
          </Button>
        </div>
      </CardContent>
    </Card>
  )
}
