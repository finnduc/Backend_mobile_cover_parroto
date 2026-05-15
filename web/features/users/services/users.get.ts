'server-only'

import { clerkClient } from '@clerk/nextjs/server'
import type { User } from '@/types/users.models'

function clerkUserToUser(u: NonNullable<Awaited<ReturnType<Awaited<ReturnType<typeof clerkClient>>['users']['getUser']>>>): User {
  return {
    id: u.id,
    name: [u.firstName, u.lastName].filter(Boolean).join(' ') || u.primaryEmailAddress?.emailAddress || u.id,
    email: u.primaryEmailAddress?.emailAddress ?? '',
    avatarUrl: u.imageUrl,
  }
}

export async function getAdminUsers(
  limit?: number,
  offset?: number
): Promise<{ users: User[]; totalCount: number }> {
  const client = await clerkClient()
  const result = await client.users.getUserList({
    limit,
    offset,
    orderBy: '-created_at',
  })
  return {
    users: result.data.map(clerkUserToUser),
    totalCount: result.totalCount,
  }
}

export async function getAdminUser(id: string): Promise<User | null> {
  try {
    const client = await clerkClient()
    const u = await client.users.getUser(id)
    return clerkUserToUser(u)
  } catch {
    return null
  }
}
