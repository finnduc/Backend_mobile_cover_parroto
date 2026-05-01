"use client";

import type { SignUpAuthFormSchema } from "@firebase-oss/ui-core";
import {
  useSignUpAuthFormAction,
  useSignUpAuthFormSchema,
  useUI,
  type SignUpAuthFormProps,
  useRequireDisplayName,
} from "@firebase-oss/ui-react";
import { Controller, useForm } from "react-hook-form";
import { standardSchemaResolver } from "@hookform/resolvers/standard-schema";
import { FirebaseUIError, getTranslation } from "@firebase-oss/ui-core";

import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Policies } from "./policies";

export type { SignUpAuthFormProps };

export function SignUpAuthForm(props: SignUpAuthFormProps) {
  const ui = useUI();
  const schema = useSignUpAuthFormSchema();
  const action = useSignUpAuthFormAction();
  const requireDisplayName = useRequireDisplayName();

  const form = useForm<SignUpAuthFormSchema>({
    resolver: standardSchemaResolver(schema),
    defaultValues: {
      email: "",
      password: "",
      displayName: requireDisplayName ? "" : undefined,
    },
  });

  async function onSubmit(values: SignUpAuthFormSchema) {
    try {
      const credential = await action(values);
      props.onSignUp?.(credential);
    } catch (error) {
      const message = error instanceof FirebaseUIError ? error.message : String(error);
      form.setError("root", { message });
    }
  }

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        {requireDisplayName ? (
          <Controller
            control={form.control}
            name="displayName"
            render={({ field, fieldState }) => (
              <Field data-invalid={fieldState.invalid}>
                <FieldLabel htmlFor="sign-up-display-name">
                  {getTranslation(ui, "labels", "displayName")}
                </FieldLabel>
                <Input
                  id="sign-up-display-name"
                  {...field}
                  aria-invalid={fieldState.invalid}
                />
                {fieldState.invalid && (
                  <FieldError errors={[fieldState.error]} />
                )}
              </Field>
            )}
          />
        ) : null}

        <Controller
          control={form.control}
          name="email"
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="sign-up-email">
                {getTranslation(ui, "labels", "emailAddress")}
              </FieldLabel>
              <Input
                id="sign-up-email"
                type="email"
                {...field}
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
          )}
        />

        <Controller
          control={form.control}
          name="password"
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="sign-up-password">
                {getTranslation(ui, "labels", "password")}
              </FieldLabel>
              <Input
                id="sign-up-password"
                type="password"
                {...field}
                aria-invalid={fieldState.invalid}
              />
              {fieldState.invalid && (
                <FieldError errors={[fieldState.error]} />
              )}
            </Field>
          )}
        />
      </FieldGroup>

      <Policies />
      
      <Button type="submit" disabled={ui.state !== "idle"}>
        {getTranslation(ui, "labels", "createAccount")}
      </Button>
      
      {form.formState.errors.root && (
        <FieldError errors={[form.formState.errors.root]} />
      )}
      
      {props.onSignInClick ? (
        <Button type="button" variant="link" size="sm" onClick={props.onSignInClick}>
          <span className="text-xs">
            {getTranslation(ui, "prompts", "haveAccount")} {getTranslation(ui, "labels", "signIn")}
          </span>
        </Button>
      ) : null}
    </form>
  );
}