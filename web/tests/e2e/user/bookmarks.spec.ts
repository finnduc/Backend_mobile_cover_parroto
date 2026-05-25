import { test, expect } from '@playwright/test'

test.describe('Bookmarks page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/bookmarks')
  })

  test('loads and shows page title', async ({ page }) => {
    await expect(page.getByRole('heading', { name: 'Bài học đã lưu' })).toBeVisible()
  })

  test('breadcrumb shows correct path', async ({ page }) => {
    await expect(page.getByRole('link', { name: 'Bài học' })).toBeVisible()
    await expect(page.getByText('Đã lưu')).toBeVisible()
  })

  test('page renders without error', async ({ page }) => {
    await expect(page.locator('body')).not.toContainText('500')
  })
})
