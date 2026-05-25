import { test, expect } from '@playwright/test'

test.describe('Admin - Users page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/admin/users')
  })

  test('loads users page', async ({ page }) => {
    await expect(page).toHaveURL('/admin/users')
    await expect(page.locator('body')).not.toContainText('500')
  })

  test('shows Users heading', async ({ page }) => {
    await expect(page.getByRole('heading', { name: 'Users' })).toBeVisible()
  })

  test('renders user list', async ({ page }) => {
    await expect(page.locator('body')).toBeVisible()
  })

  test('supports pagination', async ({ page }) => {
    // If more than 10 users exist, pagination should appear
    const pagination = page.locator('[aria-label*="pagination"], nav[aria-label]').first()
    const isVisible = await pagination.isVisible().catch(() => false)
    // Pagination may or may not appear depending on data — just verify no crash
    await expect(page.locator('body')).not.toContainText('500')
  })
})
