# Variables
APP_NAME = server
BIN_DIR = bin
ENTRY_POINT = cmd/server/main.go

# Go Commands
.PHONY: run test tidy build migrate-up migrate-down swagger

run:
	go run $(ENTRY_POINT)

test:
	go test ./...

swagger:
	~/go/bin/swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal


tidy:
	go mod tidy

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(ENTRY_POINT)

# Migrations (Manual via psql for simplicity, or use golang-migrate if preferred)
# Adjust PG connection string as needed
DB_URL = "postgres://postgres:postgres@localhost:5432/dictation_db?sslmode=disable"

migrate-up:
	@echo "Applying migrations..."
	@psql $(DB_URL) -f migrations/001_init.sql
	@psql $(DB_URL) -f migrations/002_seed.sql

migrate-down:
	@echo "Dropping all tables (reset)..."
	@psql $(DB_URL) -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"