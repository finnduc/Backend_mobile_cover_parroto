"use client"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { BookOpen } from "lucide-react"
import type { VocabularyItem } from "@/types/vocabulary.models"

export function DeckItemList({ items }: { items: VocabularyItem[] }) {
  if (items.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <BookOpen className="mb-4 size-12" />
        <p className="text-lg font-medium">Chưa có từ vựng nào</p>
        <p className="text-sm">Hãy thêm từ vựng đầu tiên của bạn</p>
      </div>
    )
  }

  return (
    <div className="space-y-3">
      {items.map((item) => (
        <Card key={item.id} size="sm">
          <CardHeader>
            <CardTitle className="text-base">{item.phrase}</CardTitle>
            <CardDescription>{item.meaning}</CardDescription>
          </CardHeader>
          {item.exampleSentence && (
            <CardContent>
              <p className="text-sm italic text-muted-foreground">
                {item.exampleSentence}
              </p>
              {item.note && (
                <p className="mt-1 text-xs text-muted-foreground">
                  Ghi chú: {item.note}
                </p>
              )}
            </CardContent>
          )}
        </Card>
      ))}
    </div>
  )
}
