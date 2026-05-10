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
import { mainNav, communityNav, type NavItem } from "@/lib/nav-items"

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
        <Link href={ROUTES.USER.HOME} className="flex items-center gap-2 text-lg font-bold">
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
