import Link from "next/link"
import { ROUTES } from "@/lib/routes"

export default function NotFound() {
  return (
    <div className="flex h-full flex-col items-center justify-center gap-4 p-6">
      <h1 className="text-4xl font-bold">404</h1>
      <p className="text-muted-foreground">Page or resource not found.</p>
      <Link href={ROUTES.ADMIN.DASHBOARD} className="text-sm text-primary hover:underline">
        Go to Dashboard
      </Link>
    </div>
  )
}
