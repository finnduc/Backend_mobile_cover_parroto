export const CACHE_TAGS = {
  categories: "categories",
  category: (id: number) => `category-${id}`,
  lessons: "lessons",
  lesson: (id: number) => `lesson-${id}`,
  transcripts: "transcripts",
  transcript: (id: number) => `transcript-${id}`,
  users: "users",
  user: (id: string) => `user-${id}`,
  bookmarks: "bookmarks",
  shadowingStatus: "shadowing-status",
  dictationStatus: "dictation-status",
} as const
