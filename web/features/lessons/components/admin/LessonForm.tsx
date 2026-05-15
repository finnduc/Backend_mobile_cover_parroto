"use client"

import type { Lesson } from "@/types/lessons.models"
import type { Category } from "@/types/categories.models"
import { useForm, Controller } from "react-hook-form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

export type LessonFormValues = Pick<Lesson, "title" | "description" | "thumbnailUrl" | "videoUrl" | "duration" | "level" | "categoryId">

export function LessonForm({
  defaultValues,
  categories,
  onSubmit,
}: {
  defaultValues?: LessonFormValues
  categories: Category[]
  onSubmit: (values: LessonFormValues) => void
}) {
  const form = useForm<LessonFormValues>({
    defaultValues: defaultValues ?? {
      title: "",
      description: "",
      thumbnailUrl: "",
      videoUrl: "",
      duration: 0,
      level: "beginner",
      categoryId: 1,
    },
  })

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="title"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="lesson-title">Title</FieldLabel>
              <Input id="lesson-title" {...field} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />

        <Controller
          name="description"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="lesson-desc">Description</FieldLabel>
              <Textarea id="lesson-desc" {...field} rows={3} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />

        <div className="grid grid-cols-2 gap-4">
          <Controller
            name="thumbnailUrl"
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="lesson-thumb">Thumbnail URL</FieldLabel>
                <Input id="lesson-thumb" {...field} aria-invalid={fieldState.invalid} />
                {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
              </Field>
            )}
          />
          <Controller
            name="videoUrl"
            control={form.control}
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="lesson-video">Video URL</FieldLabel>
                <Input id="lesson-video" {...field} aria-invalid={fieldState.invalid} />
                {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
              </Field>
            )}
          />
        </div>

        <div className="grid grid-cols-2 gap-4">
          <Controller
            name="duration"
            control={form.control}
            render={({ field: { value, onChange, ...field }, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="lesson-duration">Duration (seconds)</FieldLabel>
                <Input
                  id="lesson-duration"
                  type="number"
                  {...field}
                  value={value}
                  onChange={(e) => onChange(Number(e.target.value))}
                  aria-invalid={fieldState.invalid}
                />
                {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
              </Field>
            )}
          />
          <Controller
            name="level"
            control={form.control}
            render={({ field: { value, onChange, ...field }, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="lesson-level">Level</FieldLabel>
                <Select value={value} onValueChange={onChange}>
                  <SelectTrigger id="lesson-level" aria-invalid={fieldState.invalid}>
                    <SelectValue placeholder="Select level" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="beginner">Beginner</SelectItem>
                    <SelectItem value="intermediate">Intermediate</SelectItem>
                    <SelectItem value="advanced">Advanced</SelectItem>
                  </SelectContent>
                </Select>
                {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
              </Field>
            )}
          />
        </div>

        <Controller
          name="categoryId"
          control={form.control}
          render={({ field: { value, onChange, ...field }, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="lesson-category">Category</FieldLabel>
              <Select value={String(value)} onValueChange={(v) => onChange(Number(v))}>
                <SelectTrigger id="lesson-category" aria-invalid={fieldState.invalid}>
                  <SelectValue placeholder="Select category" />
                </SelectTrigger>
                <SelectContent>
                  {categories.map((c) => (
                    <SelectItem key={c.id} value={String(c.id)}>{c.name}</SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
      </FieldGroup>
      <Button type="submit" className="self-end">{defaultValues ? "Save Changes" : "Create Lesson"}</Button>
    </form>
  )
}
