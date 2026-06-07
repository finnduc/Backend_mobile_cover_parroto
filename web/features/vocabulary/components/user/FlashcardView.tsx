"use client"

import { useState } from "react"
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { ChevronLeft, ChevronRight } from "lucide-react"
import type { VocabularyItem } from "@/types/vocabulary.models"

export function FlashcardView({ items }: { items: VocabularyItem[] }) {
  const [index, setIndex] = useState(0)
  const [flipped, setFlipped] = useState(false)

  if (items.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-muted-foreground">
        <p className="text-lg font-medium">No vocabulary items in this deck</p>
      </div>
    )
  }

  const item = items[index]

  return (
    <div className="mx-auto max-w-md space-y-6">
      <div className="flex items-center justify-between">
        <Badge variant="outline">{index + 1} / {items.length}</Badge>
      </div>

      <div
        className="cursor-pointer"
        onClick={() => setFlipped(!flipped)}
      >
        <Card className="min-h-64">
          <CardContent className="flex flex-col items-center justify-center p-8 pt-8 min-h-64">
            {!flipped ? (
              <>
                <p className="text-2xl font-bold">{item.phrase}</p>
                {item.normalizedPhrase && item.normalizedPhrase !== item.phrase && (
                  <p className="mt-2 text-sm text-muted-foreground">{item.normalizedPhrase}</p>
                )}
              </>
            ) : (
              <div className="space-y-3 text-center">
                <p className="text-xl font-semibold">{item.meaning}</p>
                {item.exampleSentence && (
                  <p className="text-sm text-muted-foreground italic">{item.exampleSentence}</p>
                )}
                {item.note && (
                  <Badge variant="outline" className="text-xs">{item.note}</Badge>
                )}
              </div>
            )}
          </CardContent>
        </Card>
      </div>

      <div className="flex items-center justify-center gap-4">
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            setFlipped(false)
            setIndex((prev) => Math.max(0, prev - 1))
          }}
          disabled={index === 0}
        >
          <ChevronLeft className="size-4" />
        </Button>
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            setFlipped(false)
            setIndex((prev) => Math.min(items.length - 1, prev + 1))
          }}
          disabled={index === items.length - 1}
        >
          <ChevronRight className="size-4" />
        </Button>
      </div>

      <p className="text-center text-xs text-muted-foreground">
        Tap card to flip
      </p>
    </div>
  )
}
