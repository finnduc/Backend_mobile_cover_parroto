"use client"

import { VidstackYoutubePlayer } from "@/features/lessons/components/user/player/VidstackYoutubePlayer"
import { usePlayerContext } from "@/features/lessons/context/player-context"

export function VideoPanel({ videoUrl }: { videoUrl?: string }) {
  const { playerRef } = usePlayerContext()

  return (
    <>
      {videoUrl && (
        <div className="aspect-video w-full overflow-hidden rounded-lg">
          <VidstackYoutubePlayer videoUrl={videoUrl} ref={playerRef} />
        </div>
      )}
    </>)
}
