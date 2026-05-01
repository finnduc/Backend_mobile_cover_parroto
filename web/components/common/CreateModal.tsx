"use client"

import { type ReactNode } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"

export function CreateModal({
  open,
  onOpenChange,
  title,
  onSubmit,
  submitLabel,
  children,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  title: string
  onSubmit?: () => void
  submitLabel?: string
  children: ReactNode
}) {
  const showSubmit = submitLabel !== undefined && onSubmit !== undefined

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
        </DialogHeader>
        <div className="space-y-4 py-2">
          {children}
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>Cancel</Button>
          {showSubmit && <Button onClick={onSubmit}>{submitLabel}</Button>}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
