# Engflix API

Backend REST API cho nền tảng học tiếng Anh qua video — xây dựng bằng Go, Gin, GORM, Clerk Authentication.

## Tech Stack

| | |
|---|---|
| **Ngôn ngữ** | Go 1.25.3 |
| **Framework** | Gin |
| **ORM** | GORM (PostgreSQL) |
| **Auth** | Clerk |
| **Migration** | Goose |
| **Logging** | Zap |
| **Tài liệu** | Swagger / OpenAPI |
| **Hot reload** | Air |

---

## Cấu trúc dự án

```
api/
├── cmd/server/
│   ├── main.go              # Entry point
│   └── docs/                # Swagger generated files
├── internal/
│   ├── configs/             # Load env config
│   ├── core/
│   │   ├── constants.go     # Hằng số toàn app
│   │   ├── enums/           # UserRole enum
│   │   ├── errors/          # Custom AppError types
│   │   ├── logger/          # Zap logger setup
│   │   ├── policy/          # Authorization policies
│   │   └── response/        # Base response, pagination, AppError
│   ├── database/
│   │   ├── database.go      # Kết nối DB
│   │   ├── models/          # GORM models
│   │   ├── migrations/      # Goose SQL migrations
│   │   ├── query.go         # Common query helpers
│   │   ├── repositories/    # Generic repository layer
│   │   └── transaction/     # Transaction support
│   ├── middleware/           # Auth middleware (Clerk)
│   ├── modules/
│   │   ├── auth/            # POST /auth/complete-signup
│   │   ├── bookmark/        # GET/POST/DELETE /bookmarks
│   │   ├── category/        # GET /categories
│   │   ├── dictation_status/# POST/GET /dictation-status
│   │   ├── lesson/          # GET /lessons, GET /lessons/:lessonId
│   │   ├── shadowing_status/# POST/GET /shadowing-status
│   │   ├── transcript/      # GET /lessons/:lessonId/transcripts
│   │   ├── vocabulary_category/ # CRUD /vocabulary-categories
│   │   ├── vocabulary_deck/     # CRUD /vocabulary-decks
│   │   └── vocabulary_item/     # CRUD /vocabulary-decks/:deckId/items
│   └── utils/               # DTO mapper (copier)
├── docker/
│   └── docker-compose.yaml
├── .env.example
├── API.md                   # Tài liệu API cho frontend
├── Makefile
└── server.air.toml
```

---

## Yêu cầu

- [Go 1.25+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [goose](https://github.com/pressly/goose) — `go install github.com/pressly/goose/v3/cmd/goose@latest`
- [swag](https://github.com/swaggo/swag) — `go install github.com/swaggo/swag/cmd/swag@latest`
- [Air](https://github.com/air-verse/air) (tùy chọn, dùng cho hot reload)
- Tài khoản Clerk (xem hướng dẫn bên dưới)

---

## Bắt đầu

### 1. Clone & cài dependencies

```bash
git clone <repo-url>
cd api
go mod tidy
```

### 2. Tạo file `.env`

```bash
cp .env.example .env
```

Điền các giá trị vào `.env`:

```env
PORT=3001
GIN_MODE=debug

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=engflix
POSTGRES_SSLMODE=disable

CLERK_SECRET_KEY=sk_test_...
CLERK_PUBLISHABLE_KEY=pk_test_...
```

### 3. Setup Clerk

1. Vào [Clerk Dashboard](https://dashboard.clerk.com) → tạo application
2. **Configure** → **API Keys** → copy **Secret Key** và **Publishable Key**
3. Điền `CLERK_SECRET_KEY` và `CLERK_PUBLISHABLE_KEY` vào `.env`

### 4. Khởi động PostgreSQL

```bash
make up
```

> Nếu đã có PostgreSQL đang chạy ở port 5432, tạo database thủ công:
> ```bash
> psql -U postgres -c "CREATE DATABASE engflix;"
> ```

### 5. Chạy migrations

```bash
make migrate-up
```

### 6. Chạy server

```bash
# Development
go run ./cmd/server/main.go

# Hot reload (cần Air)
air -c server.air.toml
```

Server chạy tại: `http://localhost:3001`

---

## API Endpoints

**Base URL:** `http://localhost:3001/api/v1`

Xem [API.md](./API.md) để biết chi tiết request/response.

| Method | Endpoint | Auth | Mô tả |
|--------|----------|------|-------|
| GET | `/lessons` | ❌ | Danh sách bài học (filter: `category_id`, `level`, `search`) |
| GET | `/lessons/:lessonId` | ❌ | Chi tiết bài học |
| GET | `/categories` | ❌ | Danh sách category |
| GET | `/lessons/:lessonId/transcripts` | ❌ | Transcript của bài học |
| GET | `/bookmarks` | ✅ | Danh sách bookmark của user |
| POST | `/bookmarks/:lessonId` | ✅ | Thêm bookmark |
| DELETE | `/bookmarks/:lessonId` | ✅ | Xóa bookmark |
| POST | `/shadowing-status/:transcriptId` | ✅ | Đánh dấu đã hoàn thành shadowing |
| GET | `/shadowing-status?lesson_id=` | ✅ | Danh sách transcript đã shadowing |
| POST | `/dictation-status/:transcriptId` | ✅ | Đánh dấu đã hoàn thành dictation |
| GET | `/dictation-status?lesson_id=` | ✅ | Danh sách transcript đã dictation |
| GET | `/vocabulary-categories` | ❌ | Danh sách vocabulary categories |
| GET | `/vocabulary-decks?category_id=` | ❌ | Danh sách vocabulary decks |
| GET | `/vocabulary-decks/:deckId/items` | ❌ | Danh sách từ vựng trong deck |
| POST | `/vocabulary-decks` | ✅ | Tạo deck cá nhân |
| PUT | `/vocabulary-decks/:id` | ✅ | Sửa deck của mình |
| DELETE | `/vocabulary-decks/:id` | ✅ | Xóa deck của mình |
| POST | `/vocabulary-decks/:deckId/items` | ✅ | Thêm từ vào deck của mình |
| PUT | `/vocabulary-items/:id` | ✅ | Sửa từ của mình |
| DELETE | `/vocabulary-items/:id` | ✅ | Xóa từ của mình |
| POST | `/auth/complete-signup` | ✅ | Hoàn tất đăng ký user |

---

## Swagger UI

Truy cập: `http://localhost:3001/swagger/index.html`

**Cách lấy token để test trên Swagger:**

1. Đăng nhập qua Clerk trên frontend, lấy JWT session token
2. Click **Authorize** → nhập `Bearer <session_token>` → Authorize
3. Tất cả endpoint có khóa sẽ tự động đính kèm token

**Regenerate Swagger docs** (sau khi thay đổi annotation):

```bash
make swag-init
```

---

## Database Migrations

```bash
# Tạo migration mới
make migrate-create name=create_example_table

# Chạy tất cả migrations
make migrate-up

# Rollback migration cuối
make migrate-down
```

---

## Chạy test

### Chạy tất cả test

```bash
go test ./... -v
```

### Chạy test trong một module cụ thể

```bash
go test ./internal/modules/<module-name>/services/... -v
```

Ví dụ:

```bash
go test ./internal/modules/bookmark/services/... -v
go test ./internal/modules/category/services/... -v
go test ./internal/modules/lesson/services/... -v
go test ./internal/modules/vocabulary_deck/services/... -v
go test ./internal/modules/vocabulary_item/services/... -v
```

### Chạy test với coverage

```bash
go test ./... -cover
```

### Chạy test và xuất coverage report (HTML)

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Chạy test với flag race condition

```bash
go test ./... -race -v
```

### Chỉ chạy một test function cụ thể

```bash
go test ./internal/modules/<module-name>/services/... -run TestFunctionName -v
```

---

## Docker Commands

```bash
make up        # Khởi động containers
make down      # Dừng containers
make logs      # Xem logs
make restart   # Restart containers
make clean     # Xóa containers + volumes (⚠️ mất data)
```

---

## Biến môi trường

| Biến | Mô tả | Default |
|------|-------|---------|
| `PORT` | Port server | `3001` |
| `GIN_MODE` | `debug` hoặc `release` | `debug` |
| `POSTGRES_HOST` | Host PostgreSQL | `localhost` |
| `POSTGRES_PORT` | Port PostgreSQL | `5432` |
| `POSTGRES_USER` | User PostgreSQL | `postgres` |
| `POSTGRES_PASSWORD` | Password PostgreSQL | |
| `POSTGRES_DB` | Tên database | `engflix` |
| `POSTGRES_SSLMODE` | SSL mode | `disable` |
| `CLERK_SECRET_KEY` | Clerk Secret Key | |
| `CLERK_PUBLISHABLE_KEY` | Clerk Publishable Key | |

---

## Kiến trúc module

Mỗi module theo pattern `controller → service → repository`:

```
internal/modules/<module-name>/
├── <module-name>.module.go        # Route registration + DI wiring
├── <module-name>.controller.go    # HTTP handlers với Swagger docs
├── dtos/
│   ├── req/                       # Request DTOs
│   └── res/                       # Response DTOs
├── services/                      # Business logic + interface
└── repositories/                  # Data access + interface
```

Xem [MODULE_GUIDE.md](./internal/modules/MODULE_GUIDE.md) để biết chi tiết.
