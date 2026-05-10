'use client'

import { useRouter } from 'next/navigation';
import { SignUpAuthScreen } from '@/components/sign-up-auth-screen';
import { ROUTES } from '@/lib/routes';

export function SignUpScreen() {
  const router = useRouter();

  const onSignUp = () => {
    router.push(ROUTES.ONBOARDING);
  }

  return (
    <>
      <SignUpAuthScreen onSignUp={onSignUp} />
    </>
  )
}