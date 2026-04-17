# Pioneer AI Doc API

A Go-based REST API project using Gin framework and Gorm ORM.

## Tech Stack

- **Language**: Go 1.25.3
- **Framework**: Gin
- **ORM**: Gorm
- **Documentation**: Swagger

## Project Structure

```
api/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── configs/             # Configuration files
│   ├── core/                # Shared utilities and constants
│   │   ├── constants.go
│   │   ├── enums/
│   │   └── response/        # Response helpers (pagination, errors, base-response)
│   ├── database/            # Database layer
│   │   └── models/           # Database models and migrations
│   ├── modules/             # Feature modules (each module is a package)
│   │   └── example/
│   │       ├── example.module.go   # Module route registration
│   │       └── example.controller.go
│   ├── storage/             # Storage interfaces and implementations
│   │   ├── storage.interface.go
│   │   ├── storage.s3.go
│   │   └── storage.mock.go
│   └── utils/               # Utility functions (DTOs, etc.)
└── go.mod
```

## Module Convention

Each module in `@internal/modules/` follows the naming convention:
- `{module-name}.module.go` - Contains route registration function `RegisterRoutes()`
- `{module-name}.controller.go` - Contains HTTP handlers

Example:
```go
// internal/modules/example/example.module.go
package example

func RegisterRoutes() {
}
```

## API Endpoints

- **Base URL**: `http://localhost:3001/api`
- **Swagger Docs**: `http://localhost:3001/swagger/index.html`

## Getting Started

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

3. Access Swagger documentation at `http://localhost:3001/swagger/index.html`