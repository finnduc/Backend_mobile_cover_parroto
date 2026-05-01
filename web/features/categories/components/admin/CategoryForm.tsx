"use client"

import { useForm, Controller } from "react-hook-form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"

export interface CategoryFormValues {
  name: string
}

export function CategoryForm({ defaultValues, onSubmit }: {
  defaultValues?: CategoryFormValues
  onSubmit: (values: CategoryFormValues) => void
}) {
  const form = useForm<CategoryFormValues>({
    defaultValues: defaultValues ?? { name: "" },
  })

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="name"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="category-name">Name</FieldLabel>
              <Input
                id="category-name"
                {...field}
                placeholder="Enter category name"
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
          )}
        />
      </FieldGroup>
      <Button type="submit" size="sm">{defaultValues ? "Save" : "Create"}</Button>
    </form>
  )
}
