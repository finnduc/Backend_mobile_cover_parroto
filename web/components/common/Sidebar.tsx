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
  Bookmark,
  BookOpen,
  BookText,
} from "lucide-react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { useAuth } from "@clerk/nextjs"

interface NavItem {
  label: string
  href: string
  icon: LucideIcon
  requireAuth?: boolean
}

const mainNav: NavItem[] = [
  {
    label: "Bài học",
    href: ROUTES.USER.LESSONS.LIST,
    icon: BookText,
  },
  {
    label: "Chủ đề",
    href: ROUTES.USER.CATEGORIES,
    icon: BookOpen,
  },
  {
    label: "Bài học đã lưu",
    href: ROUTES.USER.BOOKMARKS,
    icon: Bookmark,
    requireAuth: true,
  },
]

const vocabulary: NavItem[] = [
  {
    label: "Từ điển",
    href: ROUTES.USER.VOCABULARY.LIST,
    icon: BookText,
  },
  {
    label: "Cá nhân",
    href: ROUTES.USER.VOCABULARY.MY_DECKS,
    icon: BookText,
    requireAuth: true,
  },
]

function NavItems({ items }: { items: NavItem[] }) {
  const pathname = usePathname()
  const { isSignedIn } = useAuth()

  const visibleItems = items.filter(
    (item) => !item.requireAuth || isSignedIn
  )

  const activeHref = visibleItems
    .filter(
      (item) =>
        pathname === item.href ||
        pathname.startsWith(item.href + "/")
    )
    .sort((a, b) => b.href.length - a.href.length)[0]?.href

  return (
    <>
      {visibleItems.map((item) => {
        const active = item.href === activeHref

        return (
          <SidebarMenuItem key={item.href}>
            <SidebarMenuButton
              asChild
              isActive={active}
              tooltip={item.label}
            >
              <Link href={item.href}>
                <item.icon />
                <span>{item.label}</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        )
      })}
    </>
  )
}

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader className="px-4 py-3">
        <Link
          href={ROUTES.USER.LESSONS.LIST}
          className="flex items-center gap-2 text-lg font-bold"
        >
          Engflix
        </Link>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Bài học</SidebarGroupLabel>

          <SidebarMenu>
            <NavItems items={mainNav} />
          </SidebarMenu>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupLabel>Từ vựng</SidebarGroupLabel>

          <SidebarMenu>
            <NavItems items={vocabulary} />
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
