import { test, expect } from '@playwright/test'

test.describe('Admin - Lessons page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/admin/lessons')
  })

  test('loads admin lessons page', async ({ page }) => {
    await expect(page).toHaveURL('/admin/lessons')
    await expect(page.locator('body')).not.toContainText('500')
  })

  test('shows lessons list', async ({ page }) => {
    await expect(page.locator('body')).toBeVisible()
  })

  test('has a link or button to create a new lesson', async ({ page }) => {
    const createButton = page.getByRole('link', { name: /tạo|create|thêm|new/i })
      .or(page.getByRole('button', { name: /tạo|create|thêm|new/i }))
    await expect(createButton).toBeVisible()
  })
})

test.describe('Admin - Create lesson page', () => {
  test('loads create lesson page', async ({ page }) => {
    await page.goto('/admin/lessons/create')
    await expect(page).toHaveURL('/admin/lessons/create')
    await expect(page.locator('body')).not.toContainText('500')
    await expect(page.locator('form, [role="form"]').or(page.locator('input').first())).toBeVisible()
  })
})
