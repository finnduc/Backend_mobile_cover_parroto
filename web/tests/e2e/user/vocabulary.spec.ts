import { test, expect } from '@playwright/test'

test.describe('Vocabulary page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/vocabulary')
  })

  test('loads and shows page title', async ({ page }) => {
    await expect(page.getByRole('heading', { name: 'Từ vựng' })).toBeVisible()
  })

  test('breadcrumb shows Từ vựng', async ({ page }) => {
    await expect(page.getByText('Từ vựng').first()).toBeVisible()
  })

  test('renders vocabulary browser', async ({ page }) => {
    // VocabularyBrowser component should render
    await expect(page.locator('main, [role="main"]')).toBeVisible()
  })
})

test.describe('Vocabulary - My Decks page', () => {
  test('loads my decks page', async ({ page }) => {
    await page.goto('/vocabulary/my-decks')
    await expect(page).toHaveURL('/vocabulary/my-decks')
    // Page should render without error
    await expect(page.locator('body')).not.toContainText('500')
  })
})
