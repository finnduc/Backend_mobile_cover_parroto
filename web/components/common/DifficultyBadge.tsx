import { Badge } from "@/components/ui/badge"
import { cn } from "@/lib/utils"

const levelStyles: Record<string, string> = {
  beginner: "bg-level-a1 text-white",
  intermediate: "bg-level-b1 text-white",
  advanced: "bg-level-c1 text-white",
}

export function DifficultyBadge({ level }: { level: string }) {
  return (
    <Badge className={cn("border-0 text-[11px] font-semibold", levelStyles[level] ?? "bg-muted text-muted-foreground")}>
      {level}
    </Badge>
  )
}
