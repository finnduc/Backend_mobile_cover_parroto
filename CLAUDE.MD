# 📚 Dictation Learning System – Backend Spec

A backend system for learning English via dictation (video-based, similar to Parroto).

---

## 🛠️ Tech Stack & Architecture

- **Language:** Go 1.22+
- **Framework:** Gin (`github.com/gin-gonic/gin`)
- **ORM:** GORM (`gorm.io/gorm`)
- **Database:** PostgreSQL
- **Caching:** Redis (`github.com/redis/go-redis/v9`)
- **Configuration:** Viper (YAML configs in `config/`)
- **Logging:** Zap (`go.uber.org/zap`) + Lumberjack
- **Validation:** `github.com/go-playground/validator/v10`
- **Architecture:** Clean Architecture / Layered Architecture (`Router` -> `Controller` -> `Service` -> `Repo` -> `Models`)

---

## 📂 Project Structure

```text
├── cmd/
│   └── server/          # Entry point (main.go) — chỉ khởi động server
├── config/              # Configuration files (local.yaml, production.yaml)
├── docs/                # API Documentation & Specs
├── global/              # Global singleton instances (DB, Logger, Config, Redis)
├── internal/
│   ├── controller/      # HTTP handlers — binding, validate, gọi service, trả response
│   ├── initialize/      # Khởi tạo và wire dependency (DB, Redis, Router, Logger)
│   ├── middlewares/     # Gin middlewares (Auth, RBAC, CORS, RequestID, Logger)
│   ├── models/          # GORM entities (database structs)
│   ├── repo/            # Data Access Layer — chỉ query DB/Redis
│   ├── routers/         # API route definitions
│   └── service/         # Business logic layer
├── log/                 # Log runtime output directory
├── migrations/          # Versioned SQL migration files (001_init.sql, ...)
├── pkg/
│   ├── logger/          # Custom Zap logger wrapper
│   ├── response/        # Standardized API response helpers
│   ├── errors/          # Custom domain error types
│   ├── settings/        # Viper config structs
│   └── utils/           # Helpers (Levenshtein, Hash, Pagination, ...)
├── scripts/             # Build / deployment shell scripts
├── test/                # Unit & integration tests + fixtures
├── third_party/         # Third-party tools or modified vendor packages
├── .gitignore
├── CLAUDE.MD            # This specification document
├── Makefile
├── go.mod
└── go.sum
```

---

# 💡 Coding Conventions & Business Rules

## 1. Layered Architecture (Clean Architecture)

- **Router** (`internal/routers/`): Khai báo API route, gắn middleware, mapping tới Controller.
- **Controller** (`internal/controller/`): Chỉ binding input (JSON/Params), validate cơ bản, gọi Service. **KHÔNG** viết business logic tại đây.
- **Service** (`internal/service/`): Chứa **100% Business Logic** (scoring, Levenshtein, Redis check, progress update).
- **Repo** (`internal/repo/`): Data Access Layer — chỉ query DB/Redis, hoàn toàn không biết về business rule.

---

## 2. Interface-driven Design ⭐

**Mọi layer đều phải định nghĩa interface của chính nó. Controller phụ thuộc vào Service interface, không phụ thuộc vào concrete struct.**

```go
// internal/repo/interface.go
type ILessonRepo interface {
    FindByID(ctx context.Context, id uint) (*models.Lesson, error)
    FindAll(ctx context.Context, filter LessonFilter) ([]models.Lesson, int64, error)
}

// internal/service/interface.go
type ILessonService interface {
    GetLesson(ctx context.Context, id uint) (*models.Lesson, error)
    ListLessons(ctx context.Context, q PaginationQuery) ([]models.Lesson, int64, error)
}
```

> **Lý do**: Dễ mock khi unit test, dễ swap implementation mà không đụng đến layer trên.

---

## 3. Dependency Injection ⭐

**Không gọi `global.DB` hay `global.Redis` trực tiếp trong Service/Repo. Inject qua constructor.**

```go
// Repo constructor
func NewLessonRepo(db *gorm.DB) ILessonRepo {
    return &lessonRepo{db: db}
}

// Service constructor
func NewLessonService(repo ILessonRepo, rdb *redis.Client) ILessonService {
    return &lessonService{repo: repo, rdb: rdb}
}

// Controller constructor
func NewLessonController(svc ILessonService) *LessonController {
    return &LessonController{svc: svc}
}
```

Wire toàn bộ dependency tại `internal/initialize/wire.go` (thủ công hoặc dùng `google/wire`).

---

## 4. Context Propagation ⭐

**Tất cả method từ Controller đến Repo đều nhận `ctx context.Context` làm tham số đầu tiên.**

```go
// Controller
func (ctrl *LessonController) GetLesson(c *gin.Context) {
    lesson, err := ctrl.svc.GetLesson(c.Request.Context(), id)
}

// Service
func (s *lessonService) GetLesson(ctx context.Context, id uint) (*models.Lesson, error) {
    return s.repo.FindByID(ctx, id)
}

// Repo
func (r *lessonRepo) FindByID(ctx context.Context, id uint) (*models.Lesson, error) {
    var lesson models.Lesson
    return &lesson, r.db.WithContext(ctx).First(&lesson, id).Error
}
```

> **Lý do**: Support timeout, cancellation, và distributed tracing về sau.

---

## 5. Error Handling Convention ⭐

**Service layer KHÔNG trả raw GORM error ra ngoài. Phải wrap thành domain error trước.**

```go
// pkg/errors/errors.go — định nghĩa tập trung
var (
    ErrNotFound       = errors.New("not found")
    ErrUnauthorized   = errors.New("unauthorized")
    ErrForbidden      = errors.New("forbidden")
    ErrInvalidInput   = errors.New("invalid input")
    ErrConflict       = errors.New("conflict")
    ErrInternalServer = errors.New("internal server error")
)

// Service layer — wrap error với context
func (s *lessonService) GetLesson(ctx context.Context, id uint) (*models.Lesson, error) {
    lesson, err := s.repo.FindByID(ctx, id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("lessonService.GetLesson: %w", pkgerrors.ErrNotFound)
    }
    if err != nil {
        return nil, fmt.Errorf("lessonService.GetLesson: %w", err)
    }
    return lesson, nil
}

// Controller — map domain error sang response code
func (ctrl *LessonController) GetLesson(c *gin.Context) {
    lesson, err := ctrl.svc.GetLesson(c.Request.Context(), id)
    if err != nil {
        switch {
        case errors.Is(err, pkgerrors.ErrNotFound):
            response.ErrorResponseData(c, response.ErrCodeNotFound, nil)
        default:
            response.ErrorResponseData(c, response.ErrCodeInternal, nil)
        }
        return
    }
    response.SuccessReponseData(c, lesson)
}
```

---

## 6. Chuẩn hoá API Response (`pkg/response`)

- **Tất cả** controller phải gọi `response.SuccessReponseData` hoặc `response.ErrorResponseData`.
- Format chuẩn: `{ "code": <int>, "message": "<string>", "data": <any|null> }`
- Error code quản lý tập trung tại `response.MsgFlags`.
- HTTP status luôn là `200 OK` — frontend đọc trường `code` nội bộ để xử lý lỗi.

```go
// List response phải kèm metadata pagination
{
  "code": 200,
  "message": "success",
  "data": {
    "items": [...],
    "meta": { "total": 100, "page": 1, "limit": 10 }
  }
}
```

---

## 7. Input Validation

- Dùng `go-playground/validator` với struct tags.
- Validate tại Controller trước khi gọi Service.
- Custom validator cho domain rules (`level` chỉ nhận `"easy"/"medium"/"hard"`).

```go
type CreateAttemptRequest struct {
    LessonID uint `json:"lesson_id" binding:"required,gt=0"`
}

type SubmitAnswerRequest struct {
    AttemptID    uint   `json:"attempt_id"   binding:"required,gt=0"`
    TranscriptID uint   `json:"transcript_id" binding:"required,gt=0"`
    AnswerText   string `json:"answer_text"   binding:"required"`
}

type BulkSubmitAnswerRequest struct {
    AttemptID uint                  `json:"attempt_id" binding:"required,gt=0"`
    Answers   []AnswerItemRequest   `json:"answers"    binding:"required,min=1,dive"`
}

type AnswerItemRequest struct {
    TranscriptID uint   `json:"transcript_id" binding:"required,gt=0"`
    AnswerText   string `json:"answer_text"   binding:"required"`
}
```

---

## 8. Pagination chuẩn

```go
// pkg/utils/pagination.go
type PaginationQuery struct {
    Page  int `form:"page"  binding:"omitempty,min=1"`
    Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

func (p *PaginationQuery) Normalize() {
    if p.Page == 0  { p.Page = 1 }
    if p.Limit == 0 { p.Limit = 10 }
}

func (p *PaginationQuery) Offset() int {
    return (p.Page - 1) * p.Limit
}
```

Mọi list endpoint đều phải trả `meta` kèm `total`, `page`, `limit`.

---

## 9. Structured Logging

```go
// Dùng structured fields, KHÔNG dùng string format
logger.Info("attempt created",
    zap.Int("attempt_id", id),
    zap.Int("user_id", userID),
    zap.String("request_id", requestID),
)

// Error phải kèm stack trace
logger.Error("failed to create attempt",
    zap.Error(err),
    zap.String("request_id", requestID),
)
```

- **Request ID middleware**: tạo UUID per-request, gán vào context, tất cả log trong request đó đều kèm `request_id`.
- Log level: `DEBUG` (dev), `INFO` (production).

---

## 10. Database Best Practices

- Dùng `bigserial` (`uint64`) cho primary key, tránh overflow khi scale.
- Mọi bảng có `created_at`, `updated_at`. Bảng quan trọng có `deleted_at` (soft delete).
- **Index**: tất cả FK column phải có index. Composite index cho query pattern thường gặp (e.g., `(user_id, lesson_id)` trên `user_progress`).
- **Migration**: dùng versioned SQL files (`001_init.sql`, `002_add_rbac.sql`). **Không dùng GORM AutoMigrate trên production.**
- **Validation rule**: `transcript_id` trong `user_answers` phải thuộc `attempt.lesson_id` — validate tại Service trước khi insert.

---

## 11. Configuration & Secret Management

- Secret (JWT secret, DB password, Redis password) **KHÔNG** hardcode trong YAML — đọc từ environment variable.
- YAML chỉ chứa non-sensitive config (port, log level, cache TTL).
- **Fail fast**: validate config khi startup, panic nếu thiếu required fields.

```yaml
# config/local.yaml — non-sensitive only
server:
  port: 8080
  read_timeout: 10s
  write_timeout: 10s

jwt:
  access_token_ttl: 15m
  refresh_token_ttl: 7d
  # secret: đọc từ env JWT_SECRET

redis:
  ttl:
    transcript_cache: 1h
    token_blacklist: 15m   # = access_token_ttl
    refresh_token: 7d
```

---

# 🧱 Database Schema

### 🧑 Users
```dbml
Table users {
  id bigserial [primary key]
  email varchar [unique, not null]
  password_hash varchar [not null]
  name varchar
  avatar_url varchar
  is_active boolean [default: true]
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}
```

### 🔐 RBAC Tables
```dbml
Table roles {
  id bigserial [primary key]
  name varchar [unique, not null]   // "admin", "user"
  created_at timestamp
}

Table permissions {
  id bigserial [primary key]
  name varchar [unique, not null]   // "lesson:create", "lesson:delete"
  created_at timestamp
}

Table user_roles {
  user_id bigint [not null]
  role_id bigint [not null]
  indexes { (user_id, role_id) [pk] }
}

Table role_permissions {
  role_id bigint [not null]
  permission_id bigint [not null]
  indexes { (role_id, permission_id) [pk] }
}
```

### 🎬 Lessons
```dbml
Table lessons {
  id bigserial [primary key]
  title varchar
  description text
  video_url varchar [not null]
  thumbnail_url varchar
  level varchar   // "easy" | "medium" | "hard"
  duration float
  created_at timestamp
  updated_at timestamp
  deleted_at timestamp
}
```

### 🧩 Transcripts (Core)
```dbml
Table transcripts {
  id bigserial [primary key]
  lesson_id bigint [not null]
  sequence integer [not null]
  content text [not null]
  phonetic varchar
  vietnamese text
  start_timestamp float
  end_timestamp float

  indexes { lesson_id }
}
```

### 🔁 Attempts
```dbml
Table attempts {
  id bigserial [primary key]
  user_id bigint [not null]
  lesson_id bigint [not null]
  total_score float
  completed boolean [default: false]
  created_at timestamp
  updated_at timestamp

  indexes { (user_id, lesson_id) }
}
```

### ✍️ User Answers
```dbml
Table user_answers {
  id bigserial [primary key]
  attempt_id bigint [not null]
  transcript_id bigint [not null]
  answer_text text
  score float
  is_correct boolean
  created_at timestamp

  indexes { attempt_id }
  indexes { (attempt_id, transcript_id) [unique] }
}
```

### 📊 User Progress
```dbml
Table user_progress {
  id bigserial [primary key]
  user_id bigint [not null]
  lesson_id bigint [not null]
  last_sequence integer [default: 0]
  answer_count integer [default: 0]   // dùng cho running average
  completed boolean [default: false]
  score_avg float [default: 0]
  updated_at timestamp

  indexes { (user_id, lesson_id) [unique] }
}
```

### 🏷️ Categories & Bookmarks
```dbml
Table categories {
  id bigserial [primary key]
  name varchar [unique, not null]
}

Table lesson_categories {
  lesson_id bigint [not null]
  category_id bigint [not null]
  indexes { (lesson_id, category_id) [pk] }
}

Table bookmarks {
  user_id bigint [not null]
  lesson_id bigint [not null]
  created_at timestamp
  indexes { (user_id, lesson_id) [pk] }
}
```

### 🔗 Relationships
```dbml
Ref: transcripts.lesson_id > lessons.id
Ref: attempts.user_id > users.id
Ref: attempts.lesson_id > lessons.id
Ref: user_answers.attempt_id > attempts.id
Ref: user_answers.transcript_id > transcripts.id
Ref: user_progress.user_id > users.id
Ref: user_progress.lesson_id > lessons.id
Ref: lesson_categories.lesson_id > lessons.id
Ref: lesson_categories.category_id > categories.id
Ref: bookmarks.user_id > users.id
Ref: bookmarks.lesson_id > lessons.id
Ref: user_roles.user_id > users.id
Ref: user_roles.role_id > roles.id
Ref: role_permissions.role_id > roles.id
Ref: role_permissions.permission_id > permissions.id
```

---

# 🌐 API Design

**Base path**: `/api/v1`

**Auth header**: `Authorization: Bearer <access_token>`

---

### 🔐 Auth
```
POST /api/v1/auth/register   → đăng ký
POST /api/v1/auth/login      → đăng nhập, trả access_token + refresh_token
POST /api/v1/auth/logout     → blacklist access_token, xoá refresh_token
POST /api/v1/auth/refresh    → đổi refresh_token lấy access_token mới
```

**Login Response:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "access_token": "...",
    "refresh_token": "...",
    "expires_in": 900
  }
}
```

---

### 📚 Lessons
```
GET  /api/v1/lessons              → list (query: page, limit, level, category_id)
GET  /api/v1/lessons/:id          → detail
GET  /api/v1/lessons/:id/transcripts → list transcripts (ordered by sequence)
```

---

### 🔁 Attempts
```
POST /api/v1/attempts             → tạo attempt mới
GET  /api/v1/attempts/:id         → detail attempt kèm answers
```

**POST /attempts Request:**
```json
{ "lesson_id": 1 }
```

---

### ✍️ Answers
```
POST /api/v1/answers              → submit 1 câu
POST /api/v1/answers/bulk         → submit nhiều câu
```

**POST /answers Request:**
```json
{
  "attempt_id": 1,
  "transcript_id": 5,
  "answer_text": "hello world"
}
```

**POST /answers/bulk Request:**
```json
{
  "attempt_id": 1,
  "answers": [
    { "transcript_id": 5, "answer_text": "hello world" },
    { "transcript_id": 6, "answer_text": "how are you" }
  ]
}
```

**Answer Response:**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "score": 0.9,
    "is_correct": false,
    "correct_text": "Hello, world!"
  }
}
```

---

### 📊 Progress
```
GET  /api/v1/progress/:lesson_id  → load tiến độ hiện tại
POST /api/v1/progress             → update thủ công (nếu cần)
```

---

### ⭐ Bookmarks & Categories
```
POST /api/v1/bookmarks            → toggle bookmark
GET  /api/v1/bookmarks            → list bookmarked lessons
GET  /api/v1/categories           → list all categories
GET  /api/v1/categories/:id/lessons → list lessons by category
```

**Toggle Bookmark Response:**
```json
{ "code": 200, "message": "success", "data": { "bookmarked": true } }
```

---

# ⚙️ GORM Models

```go
package models

import (
    "time"
    "gorm.io/gorm"
)

// Base model — embed vào tất cả entity
type Base struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"` // soft delete
}

type Lesson struct {
    Base
    Title        string       `gorm:"type:varchar(255)"`
    Description  string       `gorm:"type:text"`
    VideoURL     string       `gorm:"not null"`
    ThumbnailURL string
    Level        string       `gorm:"type:varchar(20)"`
    Duration     float64
    Transcripts  []Transcript `gorm:"foreignKey:LessonID"`
}

type Transcript struct {
    Base
    LessonID       uint    `gorm:"not null;index"`
    Sequence       int     `gorm:"not null"`
    Content        string  `gorm:"type:text;not null"`
    Phonetic       string
    Vietnamese     string  `gorm:"type:text"`
    StartTimestamp float64
    EndTimestamp   float64
}

type Attempt struct {
    Base
    UserID     uint         `gorm:"not null;index"`
    LessonID   uint         `gorm:"not null;index"`
    TotalScore float64
    Completed  bool         `gorm:"default:false"`
    Answers    []UserAnswer `gorm:"foreignKey:AttemptID"`
}

type UserAnswer struct {
    Base
    AttemptID    uint    `gorm:"not null;index"`
    TranscriptID uint    `gorm:"not null"`
    AnswerText   string  `gorm:"type:text"`
    Score        float64
    IsCorrect    bool
}

type UserProgress struct {
    ID           uint      `gorm:"primaryKey"`
    UserID       uint      `gorm:"not null;uniqueIndex:idx_user_lesson"`
    LessonID     uint      `gorm:"not null;uniqueIndex:idx_user_lesson"`
    LastSequence int       `gorm:"default:0"`
    AnswerCount  int       `gorm:"default:0"`
    Completed    bool      `gorm:"default:false"`
    ScoreAvg     float64   `gorm:"default:0"`
    UpdatedAt    time.Time
}
```

---

# 🧠 Business Logic

## 1. Lesson Management
**Tables**: `lessons`, `transcripts`, `categories`, `lesson_categories`

**Flow**: Admin tạo lesson → gắn transcripts → gắn categories → User fetch & học.

**Cache strategy**: Cache transcript list theo `lesson_id` vào Redis (TTL: 1h). Invalidate khi có update.

---

## 2. Authentication & Authorization (RBAC)

**Login flow:**
1. Nhận `email/password`, verify Bcrypt hash.
2. Tạo `access_token` (TTL: 15m) chứa `user_id`, `role`. Tạo `refresh_token` (TTL: 7d).
3. Lưu `refresh_token` vào Redis: key `refresh:{user_id}`, value = token.

**Logout flow:**
1. Xoá `refresh:{user_id}` khỏi Redis.
2. Đưa `access_token` vào blacklist: key `blacklist:{jti}`, TTL = thời gian còn lại của token.

**Auth Middleware:**
1. Extract JWT từ `Authorization: Bearer <token>`.
2. Verify signature. Check blacklist Redis.
3. `c.Set("user_id", id)` và `c.Set("role", role)` vào Gin context.

**RBAC Middleware:**
```go
// Dùng như sau trong router
protected.GET("/lessons", RequireRole("admin"), ctrl.CreateLesson)
protected.POST("/lessons", RequirePermission("lesson:create"), ctrl.CreateLesson)
```

---

## 3. Dictation & Scoring

**Validation rule**: `transcript_id` phải thuộc về `attempt.lesson_id`. Service phải check trước khi lưu.

**Scoring Algorithm** (`pkg/utils/scoring.go`):
```go
func CalculateScore(answer, correct string) (score float64, isCorrect bool) {
    // 1. Edge case: empty answer
    if strings.TrimSpace(answer) == "" {
        return 0, false
    }
    // 2. Normalize: lowercase, remove punctuation (giữ apostrophe trong contractions)
    a := normalize(answer)
    c := normalize(correct)
    // 3. Levenshtein distance → score [0, 1]
    dist := levenshteinDistance(a, c)
    maxLen := math.Max(float64(len([]rune(a))), float64(len([]rune(c))))
    if maxLen == 0 { return 1, true }
    score = 1 - float64(dist)/maxLen
    // 4. Threshold có thể config qua Viper
    isCorrect = score >= config.ScoringThreshold // default: 0.95
    return
}
```

---

## 4. Attempts & Auto Complete

**Flow:**
1. `POST /attempts` → tạo attempt mới, trả `attempt_id`.
2. User submit từng câu → lưu `user_answers`, update `user_progress`.
3. Sau mỗi submit: Service check nếu số `user_answers` của attempt = tổng số transcripts của lesson → `attempt.completed = true`, tính `total_score`.

---

## 5. Progress Tracking

**Running average** (tránh query lại toàn bộ answers):
```go
// Sau mỗi answer mới:
progress.ScoreAvg = (progress.ScoreAvg*float64(progress.AnswerCount) + newScore) /
                    float64(progress.AnswerCount+1)
progress.AnswerCount++
progress.LastSequence = transcript.Sequence
```

---

## 6. Performance Optimization

- Cache transcript list theo `lesson_id` vào Redis (key: `transcripts:{lesson_id}`, TTL: 1h).
- Dùng `GORM Preload` có chọn lọc, tránh N+1 query.
- List endpoints phải có pagination — không query unbounded.
- DB indexes trên tất cả FK và composite query patterns.

---

# 🧪 Testing Convention

- **Unit test**: test Service với mock Repo (dùng `testify/mock`).
- File test đặt cạnh file source: `lesson_service_test.go`
- **Naming**: `TestLessonService_GetByID_NotFound`, `TestScoringService_Calculate_EmptyInput`
- **Table-driven test** cho scoring algorithm:

```go
func TestCalculateScore(t *testing.T) {
    cases := []struct {
        name      string
        answer    string
        correct   string
        wantScore float64
        wantOK    bool
    }{
        {"exact match",    "hello world", "hello world", 1.0,  true},
        {"empty answer",   "",            "hello world", 0.0,  false},
        {"minor typo",     "helo world",  "hello world", 0.91, false},
        {"contraction",    "i'm fine",    "I'm fine",    1.0,  true},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            score, ok := CalculateScore(tc.answer, tc.correct)
            assert.InDelta(t, tc.wantScore, score, 0.05)
            assert.Equal(t, tc.wantOK, ok)
        })
    }
}
```

---

## 💻 Developer Commands

```bash
make run          # Run server (dev mode)
make test         # Run all tests
make migrate-up   # Apply pending migrations
make migrate-down # Rollback last migration
go mod tidy       # Tidy modules
```

---

> **Note for AI Assistant:**
> - Respect layered architecture: `controller` → `service` → `repo`. Never skip layers.
> - Always define and use **interfaces** for repo and service. Controllers depend on service interfaces only.
> - Always propagate `context.Context` as the **first parameter** through all layers.
> - Wrap errors with `fmt.Errorf("funcName: %w", err)` at service layer. Map to domain errors in `pkg/errors`.
> - Use **structured logging** with `zap.Fields`, never `fmt.Sprintf` in log calls.
> - All DB primary keys use `uint` (bigserial). All models embed `Base` struct with soft delete.
> - Migrations go in `migrations/` as versioned SQL files. Never use `AutoMigrate` in production code.
> - Scoring threshold (`0.95`) is read from config, not hardcoded.
