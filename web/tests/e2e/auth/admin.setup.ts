import { test as setup } from '@playwright/test'
import path from 'path'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))
const AUTH_FILE = path.join(__dirname, '../../../playwright/.auth/admin.json')

setup('authenticate as admin', async ({ page }) => {
  const email = process.env.TEST_ADMIN_EMAIL
  const password = process.env.TEST_ADMIN_PASSWORD

  if (!email || !password) {
    throw new Error('TEST_ADMIN_EMAIL and TEST_ADMIN_PASSWORD must be set in .env.test')
  }

  await page.goto('/sign-in')

  await page.getByLabel('Email address').fill(email)
  await page.getByRole('button', { name: 'Continue', exact: true }).click()

  await page.getByLabel('Password').fill(password)
  await page.getByRole('button', { name: 'Continue', exact: true }).click()

  await page.waitForURL(/\/(lessons|admin|onboarding)/, { timeout: 15_000 })

  await page.context().storageState({ path: AUTH_FILE })
})
