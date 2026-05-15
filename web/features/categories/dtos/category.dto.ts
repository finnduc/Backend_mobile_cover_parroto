import type { Category } from "@/types/categories.models"

export type CreateCategoryDto = Omit<Category, "id">
export type UpdateCategoryDto = Partial<Omit<Category, "id">>
