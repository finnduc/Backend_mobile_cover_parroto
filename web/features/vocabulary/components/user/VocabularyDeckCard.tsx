"use client"

import Link from "next/link"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { BookOpen } from "lucide-react"
import type { VocabularyDeck } from "@/types/vocabulary.models"

const levelStyles: Record<string, string> = {
  beginner: "bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200",
  intermediate: "bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200",
  advanced: "bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200",
}

export function VocabularyDeckCard({ deck, href }: { deck: VocabularyDeck; href: string }) {
  return (
    <Link href={href}>
      <Card className="cursor-pointer transition-shadow hover:shadow-md" size="sm">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <BookOpen className="size-4 text-muted-foreground" />
              <CardTitle>{deck.name}</CardTitle>
            </div>
            <Badge variant="outline" className={levelStyles[deck.level] ?? ""}>
              {deck.level}
            </Badge>
          </div>
          {deck.description && (
            <CardDescription>{deck.description}</CardDescription>
          )}
        </CardHeader>
        {deck.isDefault !== undefined && (
          <CardContent>
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              {deck.isDefault && <Badge variant="secondary">Hệ thống</Badge>}
              {!deck.isDefault && deck.userId && (
                <Badge variant="outline">Của bạn</Badge>
              )}
            </div>
          </CardContent>
        )}
      </Card>
    </Link>
  )
}
