"use client"

import Link from "next/link"
import { usePathname } from "next/navigation"
import { User, Sun, Moon, LogOut } from "lucide-react"
import { useTheme } from "next-themes"
import { useEffect, useState } from "react"
import { toast } from "sonner"
import { signOut } from "firebase/auth"
import { auth } from "@/lib/firebase/client-app"
import { useUser } from "@/lib/firebase/hooks"
import { mainNav, communityNav, type NavItem } from "@/lib/nav-items"
import { ROUTES } from "@/lib/routes"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

function NavLink({ item }: { item: NavItem }) {
  const pathname = usePathname()
  const active = pathname.startsWith(item.href)

  return (
    <Link
      href={item.href}
      className={cn(
        "flex items-center gap-2 px-3 py-2 text-sm rounded-md transition-colors",
        active
          ? "text-foreground font-medium"
          : "text-muted-foreground hover:text-foreground hover:bg-muted"
      )}
    >
      <item.icon className="size-4" />
      <span>{item.label}</span>
    </Link>
  )
}

export function AppSidebar() {
  const { theme, setTheme } = useTheme()
  const [mounted, setMounted] = useState(false)
  const user = useUser()

  useEffect(() => { setMounted(true) }, [])

  const handleSignOut = async () => {
    await signOut(auth)
    toast.success("Signed out successfully", { duration: 5000 })
    window.location.href = ROUTES.USER.HOME
  }

  return (
    <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="flex h-14 items-center px-4 max-w-7xl mx-auto">
        <Link href={ROUTES.USER.HOME} className="flex items-center gap-2 text-lg font-bold mr-8 shrink-0">
          Engflix
        </Link>
        <nav className="hidden md:flex items-center gap-1">
          {mainNav.map((item) => (
            <NavLink key={item.href} item={item} />
          ))}
          <div className="w-px h-5 bg-border mx-2" />
          {communityNav.map((item) => (
            <NavLink key={item.href} item={item} />
          ))}
          {mounted && (
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
            >
              {theme === "dark" ? <Sun className="size-4" /> : <Moon className="size-4" />}
            </Button>
          )}
        </nav>
        <div className="ml-auto flex items-center gap-2">
          {user && (
            <span className="flex items-center gap-2 text-sm text-muted-foreground mr-2">
              <User className="size-4" />
              {user.displayName || user.email}
            </span>
          )}
          <Button variant="ghost" onClick={handleSignOut}>
            <LogOut className="size-4 mr-1" />
            Sign Out
          </Button>
        </div>
      </div>
    </header>
  )
}
