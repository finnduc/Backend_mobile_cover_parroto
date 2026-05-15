import type { Lesson } from "@/types/lessons.models"

export type CreateLessonDto = Omit<Lesson, "id" | "createdAt">
export type UpdateLessonDto = Partial<Omit<Lesson, "id" | "createdAt">>
