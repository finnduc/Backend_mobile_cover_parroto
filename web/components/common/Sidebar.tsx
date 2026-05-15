"use client"

import Link from "next/link"
import { usePathname } from "next/navigation"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
  SidebarHeader,
} from "@/components/ui/sidebar"
import { ROUTES } from "@/lib/routes"
import {
  BookOpen,
  BookText,
  Mic,
  GraduationCap,
  Swords,
  Repeat,
  Users,
  Trophy,
  MessageCircle,
  MessageSquare,
  Gift,
  CreditCard,
} from "lucide-react"
import type { LucideIcon } from "lucide-react"

interface NavItem {
  label: string
  href: string
  icon: LucideIcon
}

const mainNav: NavItem[] = [
  { label: "Bài học", href: ROUTES.USER.LESSONS.LIST, icon: BookText },
  { label: "Chủ đề", href: ROUTES.USER.CATEGORIES, icon: BookOpen },
  { label: "Từ vựng", href: "/vocabulary", icon: BookText },
]

const communityNav: NavItem[] = [
  { label: "Cộng đồng", href: "/community", icon: Users },
  { label: "Trò chuyện", href: "/chat", icon: MessageCircle },
  { label: "Feedback", href: "/feedbacks", icon: MessageSquare },
]

function NavItems({ items }: { items: NavItem[] }) {
  const pathname = usePathname()
  return items.map((item) => {
    const active = pathname.startsWith(item.href)
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
