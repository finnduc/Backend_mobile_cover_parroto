import { z } from "zod"

export const transcriptImportEntrySchema = z.object({
  sequence: z.number(),
  content: z.string().min(1, "content is required"),
  phonetic: z.string(),
  vietnamese: z.string(),
  startTimestamp: z.number().min(0),
  endTimestamp: z.number().min(0),
})

export type TranscriptImportEntry = z.infer<typeof transcriptImportEntrySchema>

export const transcriptImportArraySchema = z.array(transcriptImportEntrySchema)
