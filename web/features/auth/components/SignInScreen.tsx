'use client'

import { SignInAuthScreen } from '@/components/sign-in-auth-screen';

export function SignInScreen() {

  const onSignIn = () => {}

  return (
    <>
      <SignInAuthScreen onSignIn={onSignIn} />
    </>
  )
}