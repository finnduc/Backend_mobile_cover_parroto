"use client"

import { useEffect } from "react"
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
import { Textarea } from "@/components/ui/textarea"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import type {
  CreateVocabularyItemDto,
  UpdateVocabularyItemDto,
} from "@/features/vocabulary/dtos/vocabulary.dto"

export interface VocabItemFormValues {
  phrase: string
  meaning: string
  exampleSentence: string
}

const defaultValues: VocabItemFormValues = {
  phrase: "",
  meaning: "",
  exampleSentence: "",
}

export function VocabItemFormDialog({
  open,
  onOpenChange,
  mode,
  initialValues,
  onSubmit,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  mode: "create" | "edit"
  initialValues?: Partial<VocabItemFormValues>
  onSubmit: (
    values: CreateVocabularyItemDto | UpdateVocabularyItemDto
  ) => Promise<void>
}) {
  const form = useForm<VocabItemFormValues>({
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
            {mode === "create" ? "Thêm từ vựng" : "Sửa từ vựng"}
          </DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <FieldGroup>
            <Controller
              name="phrase"
              control={form.control}
              rules={{ required: "Phrase is required" }}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="vocab-phrase">Phrase</FieldLabel>
                  <Input id="vocab-phrase" {...field} aria-invalid={fieldState.invalid} />
                  {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
                </Field>
              )}
            />
            <Controller
              name="meaning"
              control={form.control}
              rules={{ required: "Meaning is required" }}
              render={({ field, fieldState }) => (
                <Field data-invalid={fieldState.invalid}>
                  <FieldLabel htmlFor="vocab-meaning">Meaning</FieldLabel>
                  <Input id="vocab-meaning" {...field} aria-invalid={fieldState.invalid} />
                  {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
                </Field>
              )}
            />
            <Controller
              name="exampleSentence"
              control={form.control}
              render={({ field }) => (
                <Field>
                  <FieldLabel htmlFor="vocab-example">Example</FieldLabel>
                  <Textarea id="vocab-example" rows={3} {...field} />
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
