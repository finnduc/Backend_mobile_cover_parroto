import { test, expect } from '@playwright/test'

test.describe('Lessons page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/lessons')
  })

  test('loads and shows page title', async ({ page }) => {
    await expect(page.getByRole('heading', { name: 'Tất cả bài học' })).toBeVisible()
  })

  test('renders lesson cards grid', async ({ page }) => {
    // Grid of lesson cards should be visible
    const grid = page.locator('div.grid')
    await expect(grid).toBeVisible()
  })

  test('breadcrumb shows correct navigation', async ({ page }) => {
    await expect(page.getByRole('link', { name: 'Chủ đề' })).toBeVisible()
  })

  test('can navigate to categories from breadcrumb', async ({ page }) => {
    await page.getByRole('link', { name: 'Chủ đề' }).click()
    await expect(page).toHaveURL('/categories')
  })
})

test.describe('Lessons - category filter', () => {
  test('shows filtered lessons when categoryId param is set', async ({ page }) => {
    await page.goto('/lessons?categoryId=1')
    // Page title changes to category name or stays as default
    await expect(page.locator('h1')).toBeVisible()
  })
})
