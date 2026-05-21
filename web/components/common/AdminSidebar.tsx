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
import { Users, FolderOpen, BookOpen, LibraryBig } from "lucide-react"

const adminNav = [
  { label: "Người dùng", href: ROUTES.ADMIN.USERS.LIST, icon: Users },
  { label: "Danh mục", href: ROUTES.ADMIN.CATEGORIES.LIST, icon: FolderOpen },
  { label: "Bài học", href: ROUTES.ADMIN.LESSONS.LIST, icon: BookOpen },
  { label: "Vocab - Categories", href: ROUTES.ADMIN.VOCABULARY.CATEGORIES.LIST, icon: LibraryBig },
  { label: "Vocab - Decks", href: ROUTES.ADMIN.VOCABULARY.DECKS.LIST, icon: LibraryBig },
]

export function AdminSidebar() {
  const pathname = usePathname()

  return (
    <Sidebar>
      <SidebarHeader className="px-4 py-3">
        <Link href={ROUTES.ADMIN.USERS.LIST} className="flex items-center gap-2 text-lg font-bold">
          Engflix Admin
        </Link>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Management</SidebarGroupLabel>
          <SidebarMenu>
            {adminNav.map((item) => {
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
            })}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
