import { test, expect } from '@playwright/test'

test.describe('Auth pages', () => {
  test('sign-in page renders Clerk sign-in component', async ({ page }) => {
    await page.goto('/sign-in')
    await expect(page.getByLabel('Email address')).toBeVisible()
  })

  test('sign-up page renders Clerk sign-up component', async ({ page }) => {
    await page.goto('/sign-up')
    await expect(page.getByLabel('Email address')).toBeVisible()
  })

  test('unauthenticated user is redirected to sign-in when accessing lessons', async ({ page }) => {
    await page.goto('/lessons')
    await expect(page).toHaveURL(/sign-in/)
  })

  test('unauthenticated user is redirected to sign-in when accessing vocabulary', async ({ page }) => {
    await page.goto('/vocabulary')
    await expect(page).toHaveURL(/sign-in/)
  })

  test('unauthenticated user is redirected to sign-in when accessing admin', async ({ page }) => {
    await page.goto('/admin/lessons')
    await expect(page).toHaveURL(/sign-in/)
  })
})
