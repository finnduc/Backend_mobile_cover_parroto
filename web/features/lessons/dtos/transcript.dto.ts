import type { Transcript } from "@/types/lessons.models"

export type CreateTranscriptDto = Omit<Transcript, "id">
export type UpdateTranscriptDto = Partial<Omit<Transcript, "id">>
