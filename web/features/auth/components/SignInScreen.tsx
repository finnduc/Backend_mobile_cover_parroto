'use client'

import { SignInAuthScreen } from '@/components/sign-in-auth-screen';
import { auth } from '@/lib/firebase/client-app';
import { useEffect } from 'react';

export function SignInScreen() {
  useEffect(() => {
    // Monitor ALL auth state changes
    const unsubscribe = auth.onAuthStateChanged((user) => {
      console.log('Auth state changed:', user);
      if (user) {
        console.log('User is authenticated:', {
          uid: user.uid,
          isAnonymous: user.isAnonymous,
          email: user.email,
          provider: user.providerData[0]?.providerId
        });
      } else {
        console.log('No user authenticated');
      }
    });

    return () => unsubscribe();
  }, []);


  const onSignIn = () => {
    console.log('onSignIn');
  }

  return (
    <>
      <SignInAuthScreen onSignIn={onSignIn} />
    </>
  )
}