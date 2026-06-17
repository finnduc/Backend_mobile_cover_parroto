import { test, expect } from '@playwright/test'

test.describe('Profile page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/profile')
  })

  test('loads profile page without errors', async ({ page }) => {
    await expect(page).toHaveURL('/profile')
    await expect(page.locator('body')).not.toContainText('500')
  })

  test('page renders content', async ({ page }) => {
    await expect(page.locator('main, [role="main"], body')).toBeVisible()
  })
})

test.describe('Settings page', () => {
  test('loads settings page without errors', async ({ page }) => {
    await page.goto('/settings')
    await expect(page).toHaveURL('/settings')
    await expect(page.locator('body')).not.toContainText('500')
  })
})

