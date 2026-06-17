import { clerkMiddleware, createRouteMatcher } from "@clerk/nextjs/server"
import { NextResponse } from "next/server"
import { UserRole } from "./lib/enums/user-role.enum"

const isProtectedRoute = createRouteMatcher(["/admin(.*)", "/profile(.*)"])
const isAdminRoute = createRouteMatcher(["/admin(.*)"])
const isOnboardingRoute = createRouteMatcher(['/onboarding'])

export default clerkMiddleware(async (auth, req) => {
  const { isAuthenticated, sessionClaims } = await auth()

  if (isAuthenticated && isOnboardingRoute(req)) {
    return NextResponse.next()
  }

  if (isAuthenticated && !sessionClaims?.metadata?.role) {
    const onboardingUrl = new URL('/onboarding', req.url)
    return NextResponse.redirect(onboardingUrl)
  }

  if (!isAuthenticated && isProtectedRoute(req)) {
    return NextResponse.redirect(new URL("/", req.url))
  }

  if (isAuthenticated && isAdminRoute(req) && sessionClaims?.metadata?.role !== UserRole.Admin) {
    return NextResponse.redirect(new URL("/", req.url))
  }

  return NextResponse.next()
})

export const config = {
  matcher: [
    "/((?!_next|[^?]*\\.(?:html?|css|js(?!on)|jpe?g|webp|png|gif|svg|ttf|woff2?|ico|csv|docx?|xlsx?|zip|webmanifest)).*)",
    "/(api|trpc)(.*)",
    "/__clerk/(.*)",
  ],
}