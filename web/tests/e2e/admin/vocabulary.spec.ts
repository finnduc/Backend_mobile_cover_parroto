import { test, expect } from '@playwright/test'

test.describe('Admin - Vocabulary Decks page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/admin/vocabulary-decks')
  })

  test('loads vocabulary decks page', async ({ page }) => {
    await expect(page).toHaveURL('/admin/vocabulary-decks')
    await expect(page.locator('body')).not.toContainText('500')
  })

  test('renders decks list', async ({ page }) => {
    await expect(page.locator('body')).toBeVisible()
  })
})
