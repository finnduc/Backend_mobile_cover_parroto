"use client"

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import { ROUTES } from "@/lib/routes"
import type { LucideIcon } from "lucide-react"
import {
  BookOpen,
  Bookmark,
  BookText,
  MessageCircle,
  MessageSquare,
  Users,
} from "lucide-react"
import Link from "next/link"
import { usePathname } from "next/navigation"

interface NavItem {
  label: string
  href: string
  icon: LucideIcon
}

const mainNav: NavItem[] = [
  { label: "Bài học", href: ROUTES.USER.LESSONS.LIST, icon: BookText },
  { label: "Chủ đề", href: ROUTES.USER.CATEGORIES, icon: BookOpen },
  { label: "Từ điển", href: "/vocabulary", icon: BookText },
  { label: "Từ vựng", href: "/vocabulary/my-decks", icon: BookText },
  { label: "Đã lưu", href: ROUTES.USER.BOOKMARKS, icon: Bookmark },
]

const communityNav: NavItem[] = [
  { label: "Cộng đồng", href: "/community", icon: Users },
  { label: "Trò chuyện", href: "/chat", icon: MessageCircle },
  { label: "Feedback", href: "/feedbacks", icon: MessageSquare },
]

function NavItems({ items }: { items: NavItem[] }) {
  const pathname = usePathname()

  const activeHref = items
    .filter((item) => pathname === item.href || pathname.startsWith(item.href + "/"))
    .sort((a, b) => b.href.length - a.href.length)[0]?.href

  return items.map((item) => {
    const active = item.href === activeHref

    return (
      <SidebarMenuItem key={item.href}>
        <SidebarMenuButton asChild isActive={active} tooltip={item.label}>
          <Link href={item.href}>
            <item.icon />
            <span>{item.label}</span>
          </Link>
        </SidebarMenuButton>
      </SidebarMenuItem>
    )
  })
}

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader className="px-4 py-3">
        <Link href={ROUTES.USER.LESSONS.LIST} className="flex items-center gap-2 text-lg font-bold">
          Engflix
        </Link>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Menu</SidebarGroupLabel>
          <SidebarMenu>
            <NavItems items={mainNav} />
          </SidebarMenu>
        </SidebarGroup>
        <SidebarGroup>
          <SidebarGroupLabel>Cộng đồng</SidebarGroupLabel>
          <SidebarMenu>
            <NavItems items={communityNav} />
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
