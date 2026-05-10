import {
  BookOpen,
  BookText,
  Users,
  MessageCircle,
  MessageSquare,
} from "lucide-react"
import type { LucideIcon } from "lucide-react"
import { ROUTES } from "./routes"

export interface NavItem {
  label: string
  href: string
  icon: LucideIcon
}

export const mainNav: NavItem[] = [
  { label: "Chủ đề", href: ROUTES.USER.CATEGORIES, icon: BookOpen },
  { label: "Từ vựng", href: "/vocabulary", icon: BookText },
]

export const communityNav: NavItem[] = [
  { label: "Cộng đồng", href: "/community", icon: Users },
  { label: "Trò chuyện", href: "/chat", icon: MessageCircle },
  { label: "Feedback", href: "/feedbacks", icon: MessageSquare },
]
