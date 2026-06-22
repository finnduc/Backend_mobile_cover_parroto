# Engflix

Nền tảng học tiếng Anh qua video — bao gồm backend API (Go), web frontend (Next.js), và mobile app (Android/Java).

## Tổng quan

| Thành phần | Thư mục | Công nghệ | Mô tả |
|------------|---------|-----------|-------|
| **API** | [`api/`](./api/) | Go, Gin, GORM, PostgreSQL, Clerk | Backend REST API xử lý bài học, từ vựng, bookmark, shadowing, dictation |
| **Web** | [`web/`](./web/) | Next.js 16, TypeScript, Tailwind CSS, shadcn/ui | Frontend web cho người dùng và admin |
| **Mobile** | [`mobile/`](./mobile/) | Android (Java/Kotlin), Clerk, Retrofit | Ứng dụng di động Android |

---

## Bắt đầu nhanh

### API

```bash
cd api
cp .env.example .env
make up                  # Khởi động PostgreSQL
make migrate-up          # Chạy migrations
go run ./cmd/server/main.go
```

Server chạy tại `http://localhost:3001`. Xem chi tiết tại [`api/README.md`](./api/README.md).

### Web

```bash
cd web
cp .env.example .env
npm install
npm run dev
```

App chạy tại `http://localhost:3000`. Xem chi tiết tại [`web/README.md`](./web/README.md).

### Mobile

Mở thư mục `mobile/` bằng Android Studio, cấu hình `CLERK_PUBLISHABLE_KEY` trong `local.properties` và chạy. Xem chi tiết tại [`mobile/README.md`](./mobile/README.md).

---

## Tính năng chính

- **Học qua video**: Xem bài học từ YouTube với transcript, phonetic, dịch nghĩa
- **Shadowing**: Luyện nói theo từng câu, theo dõi tiến độ hoàn thành
- **Dictation**: Luyện nghe chép chính tả với giao diện điền từ
- **Bookmark**: Lưu bài học yêu thích
- **Từ vựng**: Kho từ vựng theo chủ đề, tạo deck cá nhân
- **Theo dõi tiến độ**: Thống kê % hoàn thành shadowing và dictation
- **Admin panel**: Quản lý bài học, danh mục, từ vựng, người dùng

---

## Kiến trúc hệ thống

```
┌──────────┐     ┌──────────┐     ┌──────────┐
│  Mobile   │     │   Web    │     │  Admin   │
│ (Android) │     │ (Next.js)│     │  Panel   │
└─────┬─────┘     └────┬─────┘     └────┬─────┘
      │                │               │
      └────────────────┼───────────────┘
                       │
                  ┌────▼────┐
                  │   API   │  Go + Gin + GORM
                  │ (REST)  │
                  └────┬────┘
                       │
              ┌────────┼────────┐
              │                 │
         ┌────▼────┐     ┌─────▼─────┐
         │PostgreSQL│     │   Clerk   │
         │          │     │  (Auth)   │
         └──────────┘     └───────────┘
```

---

## Yêu cầu hệ thống

- **Go** 1.25+
- **Node.js** 20+
- **Docker** (để chạy PostgreSQL)
- **Android Studio** (để build mobile app)
- **Tài khoản Clerk** (cho xác thực)

---

## Môi trường

Mỗi thành phần có file `.env.example` riêng. Copy và điền giá trị phù hợp:

```bash
# API
cp api/.env.example api/.env

# Web
cp web/.env.example web/.env
```

---

## License

MIT
