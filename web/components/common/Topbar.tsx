"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { SidebarTrigger } from "@/components/ui/sidebar"
import { Show, SignInButton, UserButton, useAuth } from "@clerk/nextjs"
import { LayoutDashboard, Moon, Search, Sun } from "lucide-react"
import { useTheme } from "next-themes"
import { UserRole } from "@/lib/enums/user-role.enum"
import { ROUTES } from "@/lib/routes"

export function Topbar() {
  const { theme, setTheme } = useTheme()
  const { isLoaded, sessionClaims } = useAuth()
  const isAdmin = sessionClaims?.metadata?.role === UserRole.Admin

  return (
    <header className="flex h-14 items-center gap-4 border-b bg-background px-4">
      <SidebarTrigger />
      <div className="relative flex-1 max-w-md">
        <Search className="absolute left-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground" />
        <Input placeholder="Tra từ điển..." className="pl-9 h-9 rounded-lg bg-muted/50" />
      </div>
      <div className="flex items-center gap-1 ml-auto">
        <Button
          variant="ghost"
          size="icon"
          onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
        >
          {theme === "dark" ? <Sun className="size-4" /> : <Moon className="size-4" />}
        </Button>
        <Show when="signed-in">
          <UserButton>
            {isLoaded && isAdmin && (
              <UserButton.MenuItems>
                <UserButton.Link
                  label="Admin Dashboard"
                  labelIcon={<LayoutDashboard className="size-4" />}
                  href={ROUTES.ADMIN.USERS.LIST}
                />
              </UserButton.MenuItems>
            )}
          </UserButton>
        </Show>
        <Show when="signed-out">
          <SignInButton />
        </Show>
      </div>
    </header>
  )
}
