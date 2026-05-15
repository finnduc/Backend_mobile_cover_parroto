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
import type { CreateUserDto, UpdateUserDto } from "@/features/users/dtos/user.dto"

interface UserFormValues {
  email: string
  name: string
  password: string
}

const createDefaultValues: UserFormValues = { email: "", name: "", password: "" }

export function UserCreateForm({
  onSubmit,
}: {
  onSubmit: (values: CreateUserDto) => void
}) {
  const form = useForm<UserFormValues>({ defaultValues: createDefaultValues })

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="email"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="user-email">Email</FieldLabel>
              <Input id="user-email" type="email" {...field} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
        <Controller
          name="name"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="user-name">Name</FieldLabel>
              <Input id="user-name" {...field} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
        <Controller
          name="password"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="user-password">Password</FieldLabel>
              <Input id="user-password" type="password" {...field} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
              <p className="mt-1 text-xs text-muted-foreground">
                At least 8 characters. Accepts all ASCII characters, space, and Unicode. No truncation.
              </p>
            </Field>
          )}
        />
      </FieldGroup>
      <Button type="submit" size="sm">Create</Button>
    </form>
  )
}

export function UserEditForm({
  defaultValues,
  onSubmit,
}: {
  defaultValues: { name: string }
  onSubmit: (values: UpdateUserDto) => void
}) {
  const form = useForm<{ name: string }>({ defaultValues })

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="name"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="edit-user-name">Name</FieldLabel>
              <Input id="edit-user-name" {...field} aria-invalid={fieldState.invalid} />
              {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
            </Field>
          )}
        />
      </FieldGroup>
      <Button type="submit" size="sm">Save</Button>
    </form>
  )
}
