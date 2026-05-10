"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { SignUpAuthForm } from "@/components/sign-up-auth-form";

interface SignUpAuthScreenProps {
  children?: React.ReactNode;
  onSignUp?: () => void;
  onSignInClick?: () => void;
}

export function SignUpAuthScreen({ children, onSignUp, onSignInClick }: SignUpAuthScreenProps) {
  return (
    <div className="max-w-sm mx-auto">
      <Card>
        <CardHeader>
          <CardTitle>Sign up</CardTitle>
          <CardDescription>Enter your details to create an account</CardDescription>
        </CardHeader>
        <CardContent>
          <SignUpAuthForm onSignUp={onSignUp} onSignInClick={onSignInClick} />
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
