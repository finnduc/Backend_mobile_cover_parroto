import { test as setup, expect } from '@playwright/test'
import path from 'path'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const AUTH_FILE = path.join(__dirname, '../../../playwright/.auth/user.json')

setup('authenticate as user', async ({ page }) => {
  const email = process.env.TEST_USER_EMAIL
  const password = process.env.TEST_USER_PASSWORD

  if (!email || !password) {
    throw new Error('TEST_USER_EMAIL and TEST_USER_PASSWORD must be set in .env.test')
  }

  await page.goto('/sign-in')

  // Clerk renders sign-in form with email field
  await page.getByLabel('Email address').fill(email)
  await page.getByRole('button', { name: 'Continue', exact: true }).click()

  await page.getByLabel('Password').fill(password)
  await page.getByRole('button', { name: 'Continue', exact: true }).click()

  // Wait for successful redirect after login
  await page.waitForURL(/\/(lessons|onboarding)/, { timeout: 15_000 })

  await page.context().storageState({ path: AUTH_FILE })
})
