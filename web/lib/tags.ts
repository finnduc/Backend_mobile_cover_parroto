export const CACHE_TAGS = {
  categories: "categories",
  category: (id: number) => `category-${id}`,
  lessons: "lessons",
  lesson: (id: number) => `lesson-${id}`,
  transcripts: "transcripts",
  transcript: (id: number) => `transcript-${id}`,
  users: "users",
  user: (id: number) => `user-${id}`,
} as const
