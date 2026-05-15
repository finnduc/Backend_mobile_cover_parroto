'use server'

import { clerkClient } from '@clerk/nextjs/server'
import { updateTag, refresh } from 'next/cache'
import { CACHE_TAGS } from '@/lib/tags'
import type { CreateUserDto, UpdateUserDto } from '@/features/users/dtos/user.dto'

export async function createAdminUser(
  body: CreateUserDto
): Promise<{ error: { message: string } | null }> {
  try {
    const client = await clerkClient()
    await client.users.createUser({
      emailAddress: [body.email],
      password: body.password,
      firstName: body.name,
    })
    updateTag(CACHE_TAGS.users)
    refresh()
    return { error: null }
  } catch (err: any) {
    return { error: { message: err.errors[0].message || 'Failed to create user' } }
  }
}

export async function updateAdminUser(
  id: string,
  body: UpdateUserDto
): Promise<{ error: { message: string } | null }> {
  try {
    const client = await clerkClient()
    await client.users.updateUser(id, {
      ...(body.name && { firstName: body.name }),
    })
    updateTag(CACHE_TAGS.users)
    updateTag(CACHE_TAGS.user(id))
    refresh()
    return { error: null }
  } catch (err: any) {
    return { error: { message: err.errors[0].message || 'Failed to update user' } }
  }
}

export async function deleteAdminUser(
  id: string
): Promise<{ error: { message: string } | null }> {
  try {
    const client = await clerkClient()
    await client.users.deleteUser(id)
    updateTag(CACHE_TAGS.users)
    updateTag(CACHE_TAGS.user(id))
    refresh()
    return { error: null }
  } catch (err: any) {
    return { error: { message: err.errors[0].message || 'Failed to delete user' } }
  }
}
