'use client'

import { useRouter } from 'next/navigation';
import { SignInAuthForm } from '@/components/sign-in-auth-form';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { ROUTES } from '@/lib/routes';

export function SignInScreen() {
  const router = useRouter();

  const onSignIn = () => {
    router.push(ROUTES.USER.HOME);
  }

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
            onSignUpClick={() => router.push(ROUTES.AUTH.REGISTER)}
          />
        </CardContent>
      </Card>
    </div>
  )
}