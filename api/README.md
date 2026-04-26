# Engflix API

Backend REST API cho nền tảng học tiếng Anh qua video — xây dựng bằng Go, Gin, GORM, Firebase Auth.

## Tech Stack

| | |
|---|---|
| **Language** | Go 1.25.3 |
| **Framework** | Gin |
| **ORM** | GORM (PostgreSQL) |
| **Auth** | Firebase Authentication |
| **Migration** | Goose |
| **Docs** | Swagger / OpenAPI |
| **Hot reload** | Air |

---

## Project Structure

```
api/
├── cmd/server/
│   ├── main.go              # Entry point
│   └── docs/                # Swagger generated files
├── internal/
│   ├── configs/             # Load env config
│   ├── core/
│   │   ├── database/        # Query builder
│   │   ├── errors/          # Custom errors
│   │   ├── enums/           # UserRole enum
│   │   └── response/        # Base response, pagination, AppError
│   ├── database/
│   │   ├── models/          # GORM models
│   │   └── migrations/      # Goose SQL migrations
│   ├── firebase/            # Firebase Auth client
│   ├── middleware/          # FirebaseAuth middleware
│   ├── modules/
│   │   ├── auth/            # POST /auth/token, POST /auth/sync
│   │   ├── user/            # GET /user/profile
│   │   ├── lesson/          # GET /lessons, GET /lessons/:lessonId
│   │   ├── category/        # GET /categories
│   │   ├── bookmark/        # GET/POST/DELETE /bookmarks
│   │   ├── learning_history/# POST/GET /learning-history
│   │   └── transcript/      # GET /lessons/:lessonId/transcripts
│   └── utils/               # DTO mapper
├── docker/
│   └── docker-compose.yaml
├── .env.example
├── Makefile
└── server.air.toml
```

---

## Prerequisites

- [Go 1.25+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [goose](https://github.com/pressly/goose) — `go install github.com/pressly/goose/v3/cmd/goose@latest`
- [swag](https://github.com/swaggo/swag) — `go install github.com/swaggo/swag/cmd/swag@latest`
- Firebase project (xem hướng dẫn bên dưới)

---

## Getting Started

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
POSTGRES_DB=parroto
POSTGRES_SSLMODE=disable

FIREBASE_CREDENTIALS_FILE=/đường/dẫn/tới/serviceAccountKey.json
FIREBASE_PROJECT_ID=your-firebase-project-id
FIREBASE_WEB_API_KEY=your-firebase-web-api-key
```

### 3. Setup Firebase

1. Vào [Firebase Console](https://console.firebase.google.com) → tạo project
2. **Authentication** → **Sign-in method** → bật **Email/Password**
3. **Project Settings → Service accounts** → **Generate new private key** → tải file JSON → đặt đường dẫn vào `FIREBASE_CREDENTIALS_FILE`
4. **Project Settings → General** → copy **Web API Key** → điền vào `FIREBASE_WEB_API_KEY`
5. **Project Settings → General** → copy **Project ID** → điền vào `FIREBASE_PROJECT_ID`

### 4. Khởi động PostgreSQL

```bash
make up
```

> Nếu đã có PostgreSQL đang chạy ở port 5432, tạo database thủ công:
> ```bash
> psql -U postgres -c "CREATE DATABASE parroto;"
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

| Method | Endpoint | Auth | Mô tả |
|--------|----------|------|-------|
| POST | `/auth/token` | ❌ | Lấy Firebase ID token (email + password) |
| POST | `/auth/sync` | ❌ | Sync user Firebase vào DB |
| GET | `/user/profile` | ✅ | Lấy profile user hiện tại |
| GET | `/lessons` | ❌ | Danh sách bài học (filter: `category_id`, `level`) |
| GET | `/lessons/:lessonId` | ❌ | Chi tiết bài học |
| GET | `/categories` | ❌ | Danh sách category |
| GET | `/bookmarks` | ✅ | Danh sách bookmark của user |
| POST | `/bookmarks/:lessonId` | ✅ | Thêm bookmark |
| DELETE | `/bookmarks/:lessonId` | ✅ | Xóa bookmark |
| POST | `/learning-history` | ✅ | Ghi lại tiến độ học (upsert) |
| GET | `/learning-history` | ✅ | Lịch sử học của user |
| GET | `/learning-history/:lessonId` | ✅ | Tiến độ học 1 bài cụ thể |
| GET | `/lessons/:lessonId/transcripts` | ✅ | Transcript của bài học |

---

## Swagger UI

Truy cập: `http://localhost:3001/swagger/index.html`

**Cách lấy token để test trên Swagger:**

1. Gọi `POST /auth/token` với `email` + `password`
2. Copy `id_token` từ response
3. Click **Authorize** 🔒 → nhập `Bearer <id_token>` → Authorize
4. Tất cả endpoint có 🔒 sẽ tự động đính kèm token

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

## Chạy Unit Tests

```bash
go test ./internal/modules/.../services/... -v
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

## Environment Variables

| Biến | Mô tả | Default |
|------|-------|---------|
| `PORT` | Port server | `3001` |
| `GIN_MODE` | `debug` hoặc `release` | `debug` |
| `POSTGRES_HOST` | Host PostgreSQL | `localhost` |
| `POSTGRES_PORT` | Port PostgreSQL | `5432` |
| `POSTGRES_USER` | User PostgreSQL | `postgres` |
| `POSTGRES_PASSWORD` | Password PostgreSQL | |
| `POSTGRES_DB` | Tên database | `parroto` |
| `POSTGRES_SSLMODE` | SSL mode | `disable` |
| `FIREBASE_CREDENTIALS_FILE` | Đường dẫn file service account JSON | |
| `FIREBASE_PROJECT_ID` | Firebase Project ID | |
| `FIREBASE_WEB_API_KEY` | Firebase Web API Key | |

docker compose --env-file .env -f docker/docker-compose.yaml up -d --build