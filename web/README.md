# Engflix Web

Frontend cho nền tảng học tiếng Anh qua video — xây dựng bằng Next.js 16, Clerk Authentication, shadcn/ui.

## Tech Stack

| | |
|---|---|
| **Framework** | Next.js 16 (App Router) |
| **Ngôn ngữ** | TypeScript |
| **Auth** | Clerk |
| **UI** | shadcn/ui + Radix UI |
| **Styling** | Tailwind CSS 4 |
| **Forms** | React Hook Form + Zod |
| **Video** | Vidstack |
| **Theming** | next-themes |
| **Toast** | Sonner |
| **Icons** | Lucide React |

---

## Cấu trúc dự án

```
web/
├── app/                       # Next.js App Router
│   ├── (auth)/                # Trang đăng nhập / đăng ký
│   ├── (main)/                # Các trang chính (cần đăng nhập)
│   │   ├── bookmarks/         # Bài học đã lưu
│   │   ├── categories/        # Danh mục bài học
│   │   ├── learning-history/  # Tiến độ học tập
│   │   ├── lessons/           # Chi tiết bài học (shadowing/dictation)
│   │   ├── profile/           # Hồ sơ người dùng
│   │   ├── settings/          # Cài đặt
│   │   ├── vocabulary/        # Từ vựng + deck cá nhân
│   │   └── page.tsx           # Trang chủ
│   ├── (admin)/               # Trang quản trị
│   │   └── admin/             # Quản lý bài học, danh mục, từ vựng, người dùng
│   ├── onboarding/            # Chọn vai trò sau khi đăng ký
│   └── layout.tsx             # Root layout
├── components/                # Component dùng chung
│   ├── common/                # Component tái sử dụng
│   ├── layouts/               # Layout components
│   ├── ui/                    # shadcn/ui primitives
│   └── theme-provider.tsx     # Dark/light mode toggle
├── features/                  # Logic theo từng tính năng
│   ├── auth/                  # Auth hooks & components
│   ├── bookmarks/
│   ├── categories/
│   ├── learning-history/
│   ├── lessons/
│   ├── profile/
│   ├── settings/
│   ├── users/                 # Quản lý người dùng (admin)
│   └── vocabulary/
├── hooks/                     # React hooks dùng chung
├── lib/                       # Tiện ích
│   ├── api-fetch.ts           # API client phía server (tự động gắn token Clerk)
│   ├── case.ts                # Chuyển đổi camelCase ↔ snake_case
│   ├── enums/                 # Shared enum types
│   ├── routes.ts              # Định nghĩa route frontend
│   ├── tags.ts                # Next.js cache tags
│   └── utils.ts               # Tiện ích dùng chung
├── types/                     # TypeScript type definitions
├── proxy.ts                   # Clerk middleware (bảo vệ route)
├── .env.example
└── API.md                     # Tài liệu API backend
```

---

## Bắt đầu

### 1. Cài dependencies

```bash
cd web
npm install
```

### 2. Thiết lập biến môi trường

```bash
cp .env.example .env
```

Sửa `.env`:

```env
API_URL=http://localhost:3001/api/v1
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=pk_test_...
CLERK_SECRET_KEY=sk_test_...
```

### 3. Chạy development server

```bash
npm run dev
```

App chạy tại `http://localhost:3000`

---

## Scripts

```bash
npm run dev         # Chạy dev server với Turbopack
npm run build       # Build production
npm run start       # Chạy production server
npm run lint        # Chạy ESLint
npm run format      # Format code với Prettier
npm run typecheck   # Kiểm tra TypeScript
```

---

## Xác thực (Authentication)

Sử dụng Clerk để xác thực, bảo vệ route qua middleware (`proxy.ts`):

- **Public routes**: Trang chủ, bài học, danh mục, từ vựng (chỉ xem)
- **Protected routes**: Bookmarks, lịch sử học, hồ sơ, cài đặt
- **Admin routes**: `/admin/*` — yêu cầu `metadata.role === "admin"` trong Clerk session claims
- **Onboarding**: Người dùng mới chưa có role sẽ được chuyển hướng đến `/onboarding`

---

## Tích hợp API

Tất cả API calls đều đi qua `lib/api-fetch.ts`, một wrapper fetch phía server:

- Tự động gắn Clerk session token vào header `Authorization: Bearer`
- Tự động chuyển đổi camelCase → snake_case cho request body
- Tự động chuyển đổi snake_case → camelCase cho response data
- Hỗ trợ Next.js cache tags cho ISR/revalidation

```ts
import { apiFetch } from "@/lib/api-fetch"

const { data, error } = await apiFetch<Lesson[]>("/lessons", {
  withCredentials: true,
  query: { category_id: "1" },
})
```

---

## Thư viện chính

| Package | Mục đích |
|---------|----------|
| `@clerk/nextjs` | Xác thực người dùng |
| `@vidstack/react` | Video player hỗ trợ phụ đề |
| `radix-ui` | Headless UI primitives |
| `react-hook-form` | Quản lý form state |
| `zod` | Schema validation |
| `sonner` | Toast notifications |
| `next-themes` | Dark/light mode |
| `tailwindcss` | Utility-first CSS |
