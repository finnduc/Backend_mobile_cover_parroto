'use client'

import { SignInAuthScreen } from '@/components/sign-in-auth-screen';
import { auth } from '@/lib/firebase/client-app';
import { useRouter } from 'next/navigation';
import { useEffect, useRef } from 'react';

const ALLOWED_ROLES = ['user', 'admin'];

export function SignInScreen() {
  const router = useRouter();
  const initialLoad = useRef(true);

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(async (user) => {
      if (initialLoad.current) {
        initialLoad.current = false;
        return;
      }

      if (!user) return;

      const tokenResult = await user.getIdTokenResult();
      const role = tokenResult.claims.role as string | undefined;

      if (!role || !ALLOWED_ROLES.includes(role)) {
        router.push('/onboarding');
      } else {
        router.push('/');
      }
    });

    return () => unsubscribe();
  }, [router]);

  return (
    <>
      <SignInAuthScreen />
    </>
  )
}