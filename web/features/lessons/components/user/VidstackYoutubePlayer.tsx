"use client"

import "@vidstack/react/player/styles/base.css"
import "@vidstack/react/player/styles/default/theme.css"
import "@vidstack/react/player/styles/default/layouts/video.css"

import { MediaPlayer, MediaProvider } from "@vidstack/react"
import {
  DefaultVideoLayout,
  defaultLayoutIcons,
} from "@vidstack/react/player/layouts/default"

function extractYoutubeId(url: string): string | null {
  if (!url) return null
  try {
    const u = new URL(url)
    if (u.hostname.includes("youtube.com") || u.hostname.includes("youtu.be")) {
      if (u.hostname.includes("youtu.be")) {
        return u.pathname.slice(1)
      }
      if (u.pathname === "/watch") return u.searchParams.get("v")
      if (u.pathname.startsWith("/embed/")) return u.pathname.slice(7)
      if (u.pathname.startsWith("/v/")) return u.pathname.slice(3)
    }
  } catch {
    return null
  }
  return null
}

export function VidstackYoutubePlayer({ videoUrl }: { videoUrl: string }) {
  const youtubeId = extractYoutubeId(videoUrl) ?? videoUrl
  const src = youtubeId ? `youtube/${youtubeId}` : null

  if (!src) return null

  return (
    <div className="aspect-video w-full overflow-hidden rounded-xl">
      <MediaPlayer src={src} className="size-full">
        <MediaProvider />
        <DefaultVideoLayout icons={defaultLayoutIcons} />
      </MediaPlayer>
    </div>
  )
}
