"use client"

import "@vidstack/react/player/styles/base.css" 
import "@vidstack/react/player/styles/default/theme.css"
import "@vidstack/react/player/styles/default/layouts/video.css"

import { forwardRef } from "react"
import {
  MediaPlayer,
  MediaProvider,
  type MediaPlayerInstance,
  type MediaTimeUpdateEventDetail,
} from "@vidstack/react"
import {
  DefaultVideoLayout,
  defaultLayoutIcons,
} from "@vidstack/react/player/layouts/default"
import { extractYoutubeId } from "@/lib/utils"

export const VidstackYoutubePlayer = forwardRef<
  MediaPlayerInstance,
  { videoUrl: string; onTimeUpdate?: (detail: MediaTimeUpdateEventDetail) => void }
>(function VidstackYoutubePlayer({ videoUrl, onTimeUpdate }, ref) {
  const youtubeId = extractYoutubeId(videoUrl) ?? videoUrl
  const src = youtubeId ? `youtube/${youtubeId}` : null

  if (!src) return null

  return (
    <div className="aspect-video w-full rounded-xl">
      <MediaPlayer src={src} ref={ref} onTimeUpdate={onTimeUpdate}>
        <MediaProvider />
        <DefaultVideoLayout icons={defaultLayoutIcons} />
      </MediaPlayer>
    </div>
  )
})
