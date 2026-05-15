import { useState, useRef, useCallback } from "react"
import { toast } from "sonner"
import { replaceTranscripts, appendTranscripts } from "@/features/lessons/services/transcripts.action"
import { transcriptImportEntrySchema, type TranscriptImportEntry } from "@/features/lessons/dtos/transcript-import.dto"

type ImportMode = "replace" | "append"

interface ValidatedEntry {
  entry: TranscriptImportEntry
  valid: true
}

interface InvalidEntry {
  entry: Record<string, unknown>
  errors: string[]
  valid: false
}

type ParsedEntry = ValidatedEntry | InvalidEntry

export function useTranscriptImport(lessonId: number) {
  const fileInputRef = useRef<HTMLInputElement>(null)

  const [mode, setMode] = useState<ImportMode>("append")
  const [parsed, setParsed] = useState<ParsedEntry[] | null>(null)
  const [parseError, setParseError] = useState<string | null>(null)
  const [importing, setImporting] = useState(false)

  const validEntries = parsed?.filter((e): e is ValidatedEntry => e.valid) ?? []
  const invalidCount = (parsed?.length ?? 0) - validEntries.length

  const handleFileChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files?.[0]
      if (!file) return
      setParseError(null)
      setParsed(null)

      const reader = new FileReader()
      reader.onload = () => {
        try {
          const raw = JSON.parse(reader.result as string)
          if (!Array.isArray(raw)) {
            setParseError("JSON root must be an array of transcript objects.")
            return
          }
          if (raw.length === 0) {
            setParseError("JSON array is empty. Provide at least one transcript segment.")
            return
          }

          const validated: ParsedEntry[] = raw.map((item) => {
            const result = transcriptImportEntrySchema.safeParse(item)
            if (result.success) {
              return { entry: result.data, valid: true as const }
            }
            return {
              entry: item as Record<string, unknown>,
              errors: result.error.issues.map((i) => `${i.path.join(".")}: ${i.message}`),
              valid: false as const,
            }
          })
          setParsed(validated)
        } catch {
          setParseError("Failed to parse JSON. Check that the file contains valid JSON.")
        }
      }
      reader.readAsText(file)
    },
    []
  )

  const handleImport = async () => {
    if (validEntries.length === 0) return
    setImporting(true)
    const entries = validEntries.map((e) => e.entry)
    const res =
      mode === "replace"
        ? await replaceTranscripts(lessonId, entries)
        : await appendTranscripts(lessonId, entries)
    setImporting(false)
    if (res.error) {
      toast.error(res.error.message)
      return
    }
    toast.success(`Imported ${res.data?.length ?? 0} transcript segments.`)
    window.location.reload()
  }

  const reset = () => {
    setParsed(null)
    setParseError(null)
    setMode("append")
    if (fileInputRef.current) fileInputRef.current.value = ""
  }

  return {
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
  }
}
