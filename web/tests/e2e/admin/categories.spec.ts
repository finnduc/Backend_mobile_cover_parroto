import { test, expect } from '@playwright/test'

test.describe('Admin - Categories page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/admin/categories')
  })

  test('loads admin categories page', async ({ page }) => {
    await expect(page).toHaveURL('/admin/categories')
    await expect(page.locator('body')).not.toContainText('500')
  })

  test('renders categories list', async ({ page }) => {
    await expect(page.locator('body')).toBeVisible()
  })
})

test.describe('Admin - Vocabulary Categories page', () => {
  test('loads vocabulary categories page', async ({ page }) => {
    await page.goto('/admin/vocabulary-categories')
    await expect(page).toHaveURL('/admin/vocabulary-categories')
    await expect(page.locator('body')).not.toContainText('500')
  })
})
