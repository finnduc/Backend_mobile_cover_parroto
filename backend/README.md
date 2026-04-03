# Parroto Backend - Listening Practice System

Parroto là hệ thống Backend hỗ trợ việc học tiếng Anh qua video (dictation/listening practice), được xây dựng bằng ngôn ngữ Go với hiệu năng cao và kiến trúc sạch sẽ.

## 🚀 Công nghệ sử dụng

- **Dự án**: Go (Golang) 1.26+
- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL (GORM)
- **Caching**: Redis (Lưu trữ Token Blacklist & Cache data)
- **Dependency Injection**: [Google Wire](https://github.com/google/wire)
- **Logging**: Uber Zap & Lumberjack (Rotate logs)
- **Config Management**: Viper
- **Documentation**: Swagger (swaggo)
- **Containerization**: Docker & Docker Compose

## 🏗 Kiến trúc dự án

Dự án tuân theo kiến trúc 3 lớp (3-tier architecture):
- **Controller**: Tiếp nhận request, validate dữ liệu đầu vào.
- **Service**: Xử lý logic nghiệp vụ chính.
- **Repository**: Tương tác trực tiếp với Database và Cache.

Sơ đồ thư mục chính:
- `cmd/server`: Chứa file `main.go` khởi chạy ứng dụng.
- `internal/controller`: Xử lý HTTP handlers.
- `internal/service`: Business logic.
- `internal/repo`: Database persistence layer.
- `internal/models`: Định nghĩa các cấu trúc dữ liệu GORM.
- `internal/middlewares`: Các bộ lọc (Auth, Logger, CORS, Rate Limit...).
- `pkg/`: Các thư viện tiện ích dùng chung (utils, response, logger...).

## 🌟 Tính năng chính

- **Xác thực người dùng**: Đăng ký, đăng nhập, Refresh Token, Logout (với Blacklist Token trong Redis).
- **Quản lý bài học (Lessons)**: Danh sách bài học, chi tiết bài học kèm theo transcripts lồng nhau.
- **Hệ thống Transcript**: Hỗ trợ đồng bộ thời gian (timestamps), phiên âm (phonetics) và dịch thuật.
- **Tiến độ học tập**: Theo dõi quá trình hoàn thành bài học, điểm số trung bình.
- **Luyện tập (Attempts/Answers)**: Ghi lại các lần làm bài và chấm điểm câu trả lời.
- **Danh mục & Yêu thích**: Phân loại bài học theo category và lưu bài học yêu thích (bookmarks).

## 🛠 Hướng dẫn cài đặt & Chạy ứng dụng

### 1. Sử dụng Docker (Khuyên dùng)

Hệ thống đã được cấu hình sẵn Docker Compose bao gồm App, Postgres và Redis.

```bash
# Rebuild và chạy hệ thống
docker compose up -d --build
```

- **API Server**: http://localhost:8002
- **Swagger UI**: http://localhost:8002/swagger/index.html (Sau khi chạy app)
- **Postgres**: localhost:5432 (User: `backend_parroto`, Pass: `nguyen123`)
- **Redis**: localhost:6380 (Pass: `nguyen123`)

### 2. Chạy local (Dành cho Developer)

Yêu cầu đã cài đặt Go và khởi chạy sẵn Postgres/Redis.

```bash
# Cài đặt dependencies
go mod tidy

# Chạy migrations (nếu cần)
make migrate-up

# Khởi chạy ứng dụng
make run
```

## 📝 Tài liệu API

Dự án sử dụng Swagger để tự động hóa tài liệu API. Sau khi ứng dụng khởi chạy, bạn có thể truy cập tại:
`http://localhost:8002/swagger/index.html`

Để cập nhật tài liệu sau khi sửa code:
```bash
make swagger
```

## 📜 Giấy phép

Dự án được phát triển cho mục đích học tập và xây dựng hệ thống hỗ trợ học tiếng Anh.
