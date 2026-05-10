"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { signInWithEmailAndPassword, type AuthError } from "firebase/auth";
import { auth } from "@/lib/firebase/client-app";

import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Policies } from "./policies";

interface SignInFormValues {
  email: string;
  password: string;
}

interface SignInAuthFormProps {
  onSignIn?: () => void;
  onSignUpClick?: () => void;
  onForgotPasswordClick?: () => void;
}

export function SignInAuthForm({ onSignIn, onSignUpClick, onForgotPasswordClick }: SignInAuthFormProps) {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [rootError, setRootError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignInFormValues>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: SignInFormValues) {
    setIsSubmitting(true);
    setRootError(null);

    try {
      await signInWithEmailAndPassword(auth, values.email, values.password);
      toast.success("Signed in successfully!", { duration: 5000 });
      onSignIn?.();
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
        <Field data-invalid={!!errors.email}>
          <FieldLabel htmlFor="sign-in-email">
            Email address
          </FieldLabel>
          <Input
            id="sign-in-email"
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
          <FieldLabel htmlFor="sign-in-password" className="flex items-center gap-2">
            <span className="grow">Password</span>
            {onForgotPasswordClick ? (
              <Button type="button" variant="link" onClick={onForgotPasswordClick} size="sm">
                <span className="text-xs">Forgot password?</span>
              </Button>
            ) : null}
          </FieldLabel>
          <Input
            id="sign-in-password"
            type="password"
            {...register("password", {
              required: "Password is required",
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
        {isSubmitting ? "Signing in..." : "Sign in"}
      </Button>

      {rootError && (
        <FieldError errors={[{ message: rootError }]} />
      )}

      {onSignUpClick ? (
        <Button type="button" variant="link" size="sm" onClick={onSignUpClick}>
          <span className="text-xs">
            Don&apos;t have an account? Sign up
          </span>
        </Button>
      ) : null}
    </form>
  );
}
