"use client"

import { useUser } from "@/lib/firebase/hooks"
import { PublicNavbar } from "./PublicNavbar"
import { AppSidebar } from "./Sidebar"

export function AuthAwareNav() {
  const user = useUser()

  if (user) {
    return <AppSidebar />
  }

  return <PublicNavbar />
}
