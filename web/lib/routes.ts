export const ROUTES = {
  USER: {
    HOME: "/",
    PROFILE: "/profile",
    SETTINGS: "/settings",
    BOOKMARKS: "/bookmarks",
    LEARNING_HISTORY: "/learning-history",
    LESSONS: {
      LIST: "/lessons",
      DETAIL: (id: string) => `/lessons/${id}`,
      SHADOWING: (slug: string) => `/lessons/shadowing/${slug}`,
    },
    CATEGORIES: "/categories",
  },
  AUTH: {
    LOGIN: "/login",
    REGISTER: "/register",
    FORGOT_PASSWORD: "/forgot-password",
    RESET_PASSWORD: "/reset-password",
  },
  ADMIN: {
    DASHBOARD: "/admin/dashboard",
    USERS: {
      LIST: "/admin/users",
      DETAIL: (id: string) => `/admin/users/${id}`, // Page - complex
      // Create/Edit via modal on LIST page
    },
    CATEGORIES: {
      LIST: "/admin/categories",
      // Create/Edit/Delete ALL via modal - simple resource
    },
    LESSONS: {
      LIST: "/admin/lessons",
      CREATE: "/admin/lessons/create", // Page - many fields
      EDIT: (id: string) => `/admin/lessons/${id}/edit`, // Page - complex form
      DETAIL: (id: string) => `/admin/lessons/${id}`, // Page - preview with transcripts
      TRANSCRIPTS: (id: string) => `/admin/lessons/${id}/transcripts`, // Page - nested
    },
    BOOKMARKS: {
      LIST: "/admin/bookmarks",
      // NO create/edit - users manage their own
      // Delete modal on LIST
    },
    LEARNING_HISTORY: {
      LIST: "/admin/learning-history",
      // NO create/edit - system managed
      // View only, maybe delete modal
    },
  }
}