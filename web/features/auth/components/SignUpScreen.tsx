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
    const unsubscribe = onAuthStateChanged(auth, async (user) => {
      if (!user || user.isAnonymous || processed.current) return

      const created = new Date(user.metadata.creationTime ?? 0).getTime()
      if (Date.now() - created > 10000) return

      processed.current = true

      try {
        const token = await user.getIdToken()

        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/auth/complete-signup`,
          {
            method: 'POST',
            headers: { 'Authorization': `Bearer ${token}` },
          }
        )

        if (!res.ok) {
          const err = await res.json().catch(() => ({}))
          throw new Error(err?.error?.message || 'Failed to complete signup')
        }

        await user.getIdToken(true)

        const newToken = await user.getIdToken()
        const payload = JSON.parse(atob(newToken.split('.')[1]))
        console.log('JWT:', newToken)
        console.log('Decoded payload:', payload)

        router.push('/onboarding')
      } catch (err) {
        console.error('Signup failed:', err instanceof Error ? err.message : String(err))
      }
    })

    return unsubscribe
  }, [router])

  return (
    <>
      <SignUpAuthScreen />
    </>
  )
}