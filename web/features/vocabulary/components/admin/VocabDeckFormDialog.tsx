"use client"

import { useEffect, useState } from "react"
import { useForm, Controller } from "react-hook-form"
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import type {
  CreateVocabularyDeckDto,
  UpdateVocabularyDeckDto,
} from "@/features/vocabulary/dtos/vocabulary.dto"
import type { VocabularyCategory } from "@/types/vocabulary.models"

export interface VocabDeckFormValues {
  name: string
  description: string
  level: string
  categoryId: number | null
  thumbnailUrl: string
}

const defaultValues: VocabDeckFormValues = {
  name: "",
  description: "",
  level: "",
  categoryId: null,
  thumbnailUrl: "",
}

export function VocabDeckFormDialog({
  open,
  onOpenChange,
  mode,
  initialValues,
  categories,
  onSubmit,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  mode: "create" | "edit"
  initialValues?: Partial<VocabDeckFormValues>
  categories: VocabularyCategory[]
  onSubmit: (
    values: CreateVocabularyDeckDto | UpdateVocabularyDeckDto
  ) => Promise<void>
}) {
  const form = useForm<VocabDeckFormValues>({
    defaultValues: { ...defaultValues, ...initialValues },
  })

  useEffect(() => {
    if (open) {
      form.reset({ ...defaultValues, ...initialValues })
    }
  }, [open, initialValues, form])

  const handleSubmit = form.handleSubmit(async (values) => {
    await onSubmit(values)
    onOpenChange(false)
  })

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>
            {mode === "create" ? "Tạo bộ từ vựng" : "Sửa bộ từ vựng"}
          </DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <FieldGroup>
            <Controller
              name="name"
              control={form.control}
              rules={{ required: "Name is required" }}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="deck-name">Name</FieldLabel>
                  <Input id="deck-name" {...field} aria-invalid={fieldState.invalid} />
                  {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
                </Field>
              )}
            />
            <Controller
              name="description"
              control={form.control}
              render={({ field }) => (
                <Field>
                  <FieldLabel htmlFor="deck-description">Description</FieldLabel>
                  <Input id="deck-description" {...field} />
                </Field>
              )}
            />
            <Controller
              name="level"
              control={form.control}
              render={({ field }) => (
                <Field>
                  <FieldLabel htmlFor="deck-level">Level</FieldLabel>
                  <Input id="deck-level" placeholder="e.g. beginner, intermediate" {...field} />
                </Field>
              )}
            />
            <Controller
              name="categoryId"
              control={form.control}
              render={({ field }) => (
                <Field>
                  <FieldLabel htmlFor="deck-category">Category</FieldLabel>
                  <select
                    id="deck-category"
                    className="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm"
                    value={field.value ?? ""}
                    onChange={(e) => field.onChange(e.target.value ? Number(e.target.value) : null)}
                  >
                    <option value="">No category</option>
                    {categories.map((cat) => (
                      <option key={cat.id} value={cat.id}>{cat.name}</option>
                    ))}
                  </select>
                </Field>
              )}
            />
            <Controller
              name="thumbnailUrl"
              control={form.control}
              render={({ field }) => (
                <Field>
                  <FieldLabel htmlFor="deck-thumbnail">Thumbnail URL</FieldLabel>
                  <Input id="deck-thumbnail" {...field} />
                </Field>
              )}
            />
          </FieldGroup>
          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)}>
              Hủy
            </Button>
            <Button type="submit" size="sm">
              {mode === "create" ? "Tạo" : "Lưu"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
