import { test, expect } from '@playwright/test'

test.describe('Categories page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/categories')
  })

  test('loads and shows page title', async ({ page }) => {
    await expect(page.getByRole('heading', { name: 'Chủ đề' })).toBeVisible()
  })

  test('renders category cards grid', async ({ page }) => {
    const grid = page.locator('div.grid')
    await expect(grid).toBeVisible()
  })

  test('clicking a category navigates to lessons filtered by that category', async ({ page }) => {
    const firstCategory = page.locator('div.grid a').first()
    const isVisible = await firstCategory.isVisible()

    if (isVisible) {
      await firstCategory.click()
      await expect(page).toHaveURL(/\/lessons\?categoryId=\d+/)
    }
  })
})
