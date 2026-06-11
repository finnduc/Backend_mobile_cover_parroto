"use client"

import { SidebarProvider } from "@/components/ui/sidebar"
import { AdminSidebar } from "@/components/common/AdminSidebar"
import { UserButton, useAuth } from "@clerk/nextjs"
import { ArrowLeft } from "lucide-react"
import { UserRole } from "@/lib/enums/user-role.enum"
import { ROUTES } from "@/lib/routes"

export default function AdminLayout({ children }: { children: React.ReactNode }) {
  const { isLoaded, sessionClaims } = useAuth()
  const isAdmin = sessionClaims?.metadata?.role === UserRole.Admin

  return (
    <SidebarProvider>
      <AdminSidebar />
      <main className="flex flex-1 flex-col">
        <header className="flex h-14 items-center justify-between border-b bg-background px-6">
          <h1 className="text-sm font-medium text-muted-foreground">Admin Dashboard</h1>
          <UserButton>
            {isLoaded && isAdmin && (
              <UserButton.MenuItems>
                <UserButton.Link
                  label="Back to site"
                  labelIcon={<ArrowLeft className="size-4" />}
                  href={ROUTES.USER.LESSONS.LIST}
                />
              </UserButton.MenuItems>
            )}
          </UserButton>
        </header>
        <div className="flex-1 p-6">
          {children}
        </div>
      </main>
    </SidebarProvider>
  )
}
