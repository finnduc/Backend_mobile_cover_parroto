'use client'

import { SignUpAuthScreen } from '@/components/sign-up-auth-screen';

export function SignUpScreen() {

  const onSignUp = () => {}

  return (
    <>
      <SignUpAuthScreen onSignUp={onSignUp} />
    </>
  )
}