'use client'

import { useEffect, useRef } from 'react'
import { useRouter } from 'next/navigation'
import { onAuthStateChanged } from 'firebase/auth'
import { SignUpAuthScreen } from '@/components/sign-up-auth-screen';
import { auth } from '@/lib/firebase/client-app';

export function SignUpScreen() {
  const router = useRouter()
  const processed = useRef(false)

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      if (!user || user.isAnonymous || processed.current) return

      const created = new Date(user.metadata.creationTime ?? 0).getTime()
      if (Date.now() - created > 10000) return

      processed.current = true
      router.push('/onboarding')
    })

    return unsubscribe
  }, [router])

  return (
    <>
      <SignUpAuthScreen />
    </>
  )
}