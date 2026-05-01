"use client";

import { type UserCredential, type MultiFactorInfo } from "firebase/auth";
import { FirebaseUIError, getTranslation } from "@firebase-oss/ui-core";
import {
  useMultiFactorTotpAuthVerifyFormSchema,
  useUI,
  useTotpMultiFactorAssertionFormAction,
} from "@firebase-oss/ui-react";
import { Controller, useForm } from "react-hook-form";
import { standardSchemaResolver } from "@hookform/resolvers/standard-schema";

import {
  Field,
  FieldError,
  FieldLabel,
} from "@/components/ui/field";
import { Button } from "@/components/ui/button";
import { InputOTP, InputOTPGroup, InputOTPSlot } from "@/components/ui/input-otp";

type TotpMultiFactorAssertionFormProps = {
  hint: MultiFactorInfo;
  onSuccess?: (credential: UserCredential) => void;
};

export function TotpMultiFactorAssertionForm(props: TotpMultiFactorAssertionFormProps) {
  const ui = useUI();
  const schema = useMultiFactorTotpAuthVerifyFormSchema();
  const action = useTotpMultiFactorAssertionFormAction();

  const form = useForm<{ verificationCode: string }>({
    resolver: standardSchemaResolver(schema),
    defaultValues: {
      verificationCode: "",
    },
  });

  const onSubmit = async (values: { verificationCode: string }) => {
    try {
      const credential = await action({ verificationCode: values.verificationCode, hint: props.hint });
      props.onSuccess?.(credential);
    } catch (error) {
      const message = error instanceof FirebaseUIError ? error.message : String(error);
      form.setError("root", { message });
    }
  };

  return (
    <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-4">
      <Controller
        name="verificationCode"
        control={form.control}
        render={({ field, fieldState }) => (
          <Field data-invalid={fieldState.invalid}>
            <FieldLabel htmlFor="verify-code">
              {getTranslation(ui, "labels", "verificationCode")}
            </FieldLabel>
            <InputOTP
              id="verify-code"
              maxLength={6}
              {...field}
              aria-invalid={fieldState.invalid}
            >
              <InputOTPGroup>
                <InputOTPSlot index={0} />
                <InputOTPSlot index={1} />
                <InputOTPSlot index={2} />
                <InputOTPSlot index={3} />
                <InputOTPSlot index={4} />
                <InputOTPSlot index={5} />
              </InputOTPGroup>
            </InputOTP>
            {fieldState.invalid && (
              <FieldError errors={[fieldState.error]} />
            )}
          </Field>
        )}
      />
      <Button type="submit" disabled={ui.state !== "idle"}>
        {getTranslation(ui, "labels", "verifyCode")}
      </Button>
      {form.formState.errors.root && (
        <FieldError errors={[form.formState.errors.root]} />
      )}
    </form>
  );
}
