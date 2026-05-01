import type { ReactNode } from "react"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"

export function AdminPageLayout({
  title,
  actions,
  backHref,
  backLabel,
  children,
  maxWidth,
}: {
  title?: string
  actions?: ReactNode
  backHref?: string
  backLabel?: string
  children: ReactNode
  maxWidth?: "default" | "narrow"
}) {
  return (
    <div className={maxWidth === "narrow" ? "max-w-2xl space-y-6" : "space-y-4"}>
      {backHref && (
        <Button variant="ghost" size="sm" asChild>
          <Link href={backHref}>
            <ArrowLeft className="mr-1 size-4" />
            {backLabel ?? "Back"}
          </Link>
        </Button>
      )}
      {title && (
        <div className="flex items-center justify-between">
          <h2 className="text-xl font-bold">{title}</h2>
          {actions && <div className="flex items-center gap-2">{actions}</div>}
        </div>
      )}
      {children}
    </div>
  )
}
