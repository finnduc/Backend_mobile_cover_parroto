"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { VidstackYoutubePlayer } from "@/features/lessons/components/user/VidstackYoutubePlayer"
import { Bird, Clock } from "lucide-react"
import { usePlayerContext } from "@/features/lessons/context/player-context"

function Timer() {
  return (
    <div className="flex items-center gap-1 text-xs text-muted-foreground">
      <Clock />
      0:00
    </div>
  )
}

export function VideoPanel({ videoUrl }: { videoUrl?: string }) {
  const { playerRef } = usePlayerContext()

  return (
    <Card className="flex flex-col gap-3">
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle>VIDEO</CardTitle>
          <Timer />
        </div>
      </CardHeader>
      <CardContent className="flex flex-col gap-3">
        {videoUrl && (
          <div className="aspect-video w-full overflow-hidden rounded-lg">
            <VidstackYoutubePlayer videoUrl={videoUrl} ref={playerRef} />
          </div>
        )}
      </CardContent>
    </Card>
  )
}
