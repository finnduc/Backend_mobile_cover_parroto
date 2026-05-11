import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function extractYoutubeId(url: string): string | null {
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
