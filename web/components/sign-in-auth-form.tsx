"use client";

import type { SignInAuthFormSchema } from "@firebase-oss/ui-core";
import {
  useSignInAuthFormAction,
  useSignInAuthFormSchema,
  useUI,
  type SignInAuthFormProps,
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

export type { SignInAuthFormProps };

export function SignInAuthForm(props: SignInAuthFormProps) {
  const ui = useUI();
  const schema = useSignInAuthFormSchema();
  const action = useSignInAuthFormAction();

  const form = useForm<SignInAuthFormSchema>({
    resolver: standardSchemaResolver(schema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: SignInAuthFormSchema) {
    try {
      const credential = await action(values);
      props.onSignIn?.(credential);
    } catch (error) {
      const message = error instanceof FirebaseUIError ? error.message : String(error);
      form.setError("root", { message });
    }
  }

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Controller
          name="email"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="sign-in-email">
                {getTranslation(ui, "labels", "emailAddress")}
              </FieldLabel>
              <Input
                id="sign-in-email"
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
          name="password"
          control={form.control}
          render={({ field, fieldState }) => (
            <Field data-invalid={fieldState.invalid}>
              <FieldLabel htmlFor="sign-in-password" className="flex items-center gap-2">
                <span className="grow">{getTranslation(ui, "labels", "password")}</span>
                {props.onForgotPasswordClick ? (
                  <Button type="button" variant="link" onClick={props.onForgotPasswordClick} size="sm">
                    <span className="text-xs">{getTranslation(ui, "labels", "forgotPassword")}</span>
                  </Button>
                ) : null}
              </FieldLabel>
              <Input
                id="sign-in-password"
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
        {getTranslation(ui, "labels", "signIn")}
      </Button>

      {form.formState.errors.root && (
         <FieldError errors={[form.formState.errors.root]} />
      )}

      {props.onSignUpClick ? (
        <Button type="button" variant="link" size="sm" onClick={props.onSignUpClick}>
          <span className="text-xs">
            {getTranslation(ui, "prompts", "noAccount")} {getTranslation(ui, "labels", "signUp")}
          </span>
        </Button>
      ) : null}
    </form>
  );
}