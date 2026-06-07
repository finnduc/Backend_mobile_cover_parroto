import { defineConfig, devices } from '@playwright/test'
import path from 'path'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

const USER_AUTH_FILE = path.join(__dirname, 'playwright/.auth/user.json')
const ADMIN_AUTH_FILE = path.join(__dirname, 'playwright/.auth/admin.json')

export default defineConfig({
  testDir: './tests/e2e',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 1 : undefined,
  reporter: [['html', { open: 'never' }], ['list']],
  use: {
    baseURL: process.env.PLAYWRIGHT_BASE_URL ?? 'http://localhost:3000',
    trace: 'on-first-retry',
    screenshot: 'only-on-failure',
  },
  projects: [
    // Auth setup projects
    {
      name: 'user-setup',
      testMatch: /auth\/user\.setup\.ts/,
    },
    {
      name: 'admin-setup',
      testMatch: /auth\/admin\.setup\.ts/,
    },

    // User tests
    {
      name: 'user-chromium',
      use: {
        ...devices['Desktop Chrome'],
        storageState: USER_AUTH_FILE,
      },
      testMatch: /user\/.+\.spec\.ts/,
      dependencies: ['user-setup'],
    },

    // Admin tests
    {
      name: 'admin-chromium',
      use: {
        ...devices['Desktop Chrome'],
        storageState: ADMIN_AUTH_FILE,
      },
      testMatch: /admin\/.+\.spec\.ts/,
      dependencies: ['admin-setup'],
    },

    // Public pages (no auth required)
    {
      name: 'public',
      use: { ...devices['Desktop Chrome'] },
      testMatch: /public\/.+\.spec\.ts/,
    },
  ],
  webServer: {
    command: 'npm run dev',
    url: 'http://localhost:3000',
    reuseExistingServer: !process.env.CI,
    timeout: 120_000,
  },
})
