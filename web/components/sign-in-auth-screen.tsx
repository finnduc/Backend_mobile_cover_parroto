"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { SignInAuthForm } from "@/components/sign-in-auth-form";

interface SignInAuthScreenProps {
  children?: React.ReactNode;
  onSignIn?: () => void;
  onSignUpClick?: () => void;
  onForgotPasswordClick?: () => void;
}

export function SignInAuthScreen({ children, onSignIn, onSignUpClick, onForgotPasswordClick }: SignInAuthScreenProps) {
  return (
    <div className="max-w-sm mx-auto">
      <Card>
        <CardHeader>
          <CardTitle>Sign in</CardTitle>
          <CardDescription>Sign in to your account</CardDescription>
        </CardHeader>
        <CardContent>
          <SignInAuthForm
            onSignIn={onSignIn}
            onSignUpClick={onSignUpClick}
            onForgotPasswordClick={onForgotPasswordClick}
          />
          {children ? (
            <>
              <Separator className="my-4" />
              <div className="space-y-2">{children}</div>
            </>
          ) : null}
        </CardContent>
      </Card>
    </div>
  );
}
