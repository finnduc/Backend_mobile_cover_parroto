"use client"

import { useForm, Controller } from "react-hook-form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Field, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field"

export interface VocabularyCategoryFormValues {
  name: string
  description: string
}

export function VocabularyCategoryForm({
  defaultValues,
  onSubmit,
}: {
  defaultValues?: VocabularyCategoryFormValues
  onSubmit: (values: VocabularyCategoryFormValues) => void
}) {
  const form = useForm<VocabularyCategoryFormValues>({
    defaultValues: defaultValues ?? { name: "", description: "" },
  })

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="name"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="vc-name">Name</FieldLabel>
              <Input id="vc-name" {...field} placeholder="Enter category name" aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
        <Controller
          name="description"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="vc-desc">Description</FieldLabel>
              <Input id="vc-desc" {...field} placeholder="Enter description" aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
      </FieldGroup>
      <Button type="submit" size="sm">{defaultValues ? "Save" : "Create"}</Button>
    </form>
  )
}
