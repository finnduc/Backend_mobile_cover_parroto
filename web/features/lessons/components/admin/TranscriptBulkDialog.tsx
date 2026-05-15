"use client"

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { DataTable, type Column } from "@/components/common/DataTable"
import { Loader2, AlertCircle } from "lucide-react"
import { useTranscriptImport } from "@/features/lessons/hooks/use-transcript-import"

type PreviewRow = {
  index: number
  sequence: number
  content: string
  startTimestamp: number
  endTimestamp: number
  status: string
}

export function TranscriptBulkDialog({
  lessonId,
  open,
  onOpenChange,
}: {
  lessonId: number
  open: boolean
  onOpenChange: (open: boolean) => void
}) {
  const {
    fileInputRef,
    mode,
    setMode,
    parsed,
    parseError,
    importing,
    validEntries,
    invalidCount,
    handleFileChange,
    handleImport,
    reset,
  } = useTranscriptImport(lessonId)

  const previewData: PreviewRow[] =
    parsed?.map((item, i) => ({
      index: i + 1,
      sequence: item.entry.sequence as number,
      content: item.entry.content as string,
      startTimestamp: item.entry.startTimestamp as number,
      endTimestamp: item.entry.endTimestamp as number,
      status: item.valid ? "Valid" : item.errors.join(", "),
    })) ?? []

  const columns: Column<PreviewRow>[] = [
    { key: "index", header: "#" },
    { key: "sequence", header: "Seq" },
    {
      key: "content",
      header: "Content",
      render: (row) => <span className="max-w-60 truncate block">{row.content}</span>,
    },
    { key: "startTimestamp", header: "Start" },
    { key: "endTimestamp", header: "End" },
    {
      key: "status",
      header: "Status",
      render: (row) =>
        row.status === "Valid" ? (
          <span className="text-xs text-green-600">Valid</span>
        ) : (
          <span className="text-xs text-destructive">{row.status}</span>
        ),
    },
  ]

  return (
    <Dialog
      open={open}
      onOpenChange={(open) => {
        if (!open) reset()
        onOpenChange(open)
      }}
    >
      <DialogContent className="sm:max-w-2xl">
        <DialogHeader>
          <DialogTitle>Import Transcripts from JSON</DialogTitle>
          <DialogDescription>
            Upload a .json file with transcript segments to import.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="json-file">Select .json file</Label>
            <Input
              id="json-file"
              ref={fileInputRef}
              type="file"
              accept=".json,application/json"
              onChange={handleFileChange}
            />
          </div>

          <div className="space-y-2">
            <Label>Import mode</Label>
            <RadioGroup
              value={mode}
              onValueChange={(v) => setMode(v as "replace" | "append")}
              className="flex gap-4"
            >
              <div className="flex items-center gap-2">
                <RadioGroupItem value="append" id="mode-append" />
                <Label htmlFor="mode-append" className="cursor-pointer font-normal">
                  Append to existing
                </Label>
              </div>
              <div className="flex items-center gap-2">
                <RadioGroupItem value="replace" id="mode-replace" />
                <Label htmlFor="mode-replace" className="cursor-pointer font-normal">
                  Replace all existing
                </Label>
              </div>
            </RadioGroup>
          </div>

          {parseError && (
            <div className="flex items-start gap-2 rounded-md border border-destructive/50 bg-destructive/10 p-3 text-sm text-destructive">
              <AlertCircle className="mt-0.5 size-4 shrink-0" />
              <span>{parseError}</span>
            </div>
          )}

          {parsed && (
            <div className="space-y-2">
              <p className="text-sm text-muted-foreground">
                {parsed.length} segments found
                {invalidCount > 0 && (
                  <span className="ml-2 text-destructive">
                    ({invalidCount} with errors)
                  </span>
                )}
              </p>

              <div className="max-h-64 overflow-auto">
                <DataTable columns={columns} data={previewData} emptyMessage="No segments" />
              </div>
            </div>
          )}
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>
            Cancel
          </Button>
          <Button onClick={handleImport} disabled={validEntries.length === 0 || importing}>
            {importing && <Loader2 className="mr-1 size-4 animate-spin" />}
            Import {validEntries.length > 0 ? `${validEntries.length} segments` : ""}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
