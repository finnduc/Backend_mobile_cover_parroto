'use client'

'use client'

import { useRouter } from 'next/navigation';
import { signOut } from 'firebase/auth';
import { auth } from '@/lib/firebase/client-app';
import { SignUpAuthForm } from '@/components/sign-up-auth-form';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { ROUTES } from '@/lib/routes';

export function SignUpScreen() {
  const router = useRouter();

  const onSignUp = async () => {
    await signOut(auth);
    window.location.href = ROUTES.AUTH.LOGIN;
  }

  return (
    <div className="max-w-sm mx-auto">
      <Card>
        <CardHeader>
          <CardTitle>Sign up</CardTitle>
          <CardDescription>Enter your details to create an account</CardDescription>
        </CardHeader>
        <CardContent>
          <SignUpAuthForm
            onSignUp={onSignUp}
            onSignInClick={() => router.push(ROUTES.AUTH.LOGIN)}
          />
        </CardContent>
      </Card>
    </div>
  )
}