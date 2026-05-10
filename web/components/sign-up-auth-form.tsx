"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { createUserWithEmailAndPassword, updateProfile, type AuthError } from "firebase/auth";
import { auth } from "@/lib/firebase/client-app";
import { completeSignUp } from "@/features/auth/services/auth-service";

import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Policies } from "./policies";

interface SignUpFormValues {
  email: string;
  password: string;
  displayName?: string;
}

interface SignUpAuthFormProps {
  onSignUp?: () => void;
  onSignInClick?: () => void;
}

export function SignUpAuthForm({ onSignUp, onSignInClick }: SignUpAuthFormProps) {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [rootError, setRootError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignUpFormValues>({
    defaultValues: {
      email: "",
      password: "",
      displayName: "",
    },
  });

  async function onSubmit(values: SignUpFormValues) {
    setIsSubmitting(true);
    setRootError(null);

    try {
      const userCredential = await createUserWithEmailAndPassword(
        auth,
        values.email,
        values.password,
      );

      if (values.displayName) {
        await updateProfile(userCredential.user, {
          displayName: values.displayName,
        });
      }

      await completeSignUp();

      toast.success("Account created successfully!", { duration: 5000 });
      onSignUp?.();
    } catch (error: unknown) {
      const authError = error as AuthError;
      const message = authError.message || "An unexpected error occurred";
      toast.error(message, { duration: 5000 });
      setRootError(message);
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <FieldGroup>
        <Field data-invalid={!!errors.displayName}>
          <FieldLabel htmlFor="sign-up-display-name">
            Display name (optional)
          </FieldLabel>
          <Input
            id="sign-up-display-name"
            type="text"
            {...register("displayName")}
            aria-invalid={!!errors.displayName}
          />
          {errors.displayName && (
            <FieldError errors={[errors.displayName]} />
          )}
        </Field>

        <Field data-invalid={!!errors.email}>
          <FieldLabel htmlFor="sign-up-email">
            Email address
          </FieldLabel>
          <Input
            id="sign-up-email"
            type="email"
            {...register("email", {
              required: "Email is required",
              pattern: {
                value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
                message: "Invalid email address",
              },
            })}
            aria-invalid={!!errors.email}
          />
          {errors.email && (
            <FieldError errors={[errors.email]} />
          )}
        </Field>

        <Field data-invalid={!!errors.password}>
          <FieldLabel htmlFor="sign-up-password">
            Password
          </FieldLabel>
          <Input
            id="sign-up-password"
            type="password"
            {...register("password", {
              required: "Password is required",
              minLength: {
                value: 6,
                message: "Password must be at least 6 characters",
              },
            })}
            aria-invalid={!!errors.password}
          />
          {errors.password && (
            <FieldError errors={[errors.password]} />
          )}
        </Field>
      </FieldGroup>

      <Policies />

      <Button type="submit" disabled={isSubmitting}>
        {isSubmitting ? "Creating account..." : "Create account"}
      </Button>

      {rootError && (
        <FieldError errors={[{ message: rootError }]} />
      )}

      {onSignInClick ? (
        <Button type="button" variant="link" size="sm" onClick={onSignInClick}>
          <span className="text-xs">
            Already have an account? Sign in
          </span>
        </Button>
      ) : null}
    </form>
  );
}
