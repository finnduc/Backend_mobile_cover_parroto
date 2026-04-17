-include .env
export

.PHONY: migrate-create up down build logs restart clean swag-init


DOCKER_COMPOSE=docker compose --env-file .env -f docker/docker-compose.yaml

# Usage: make migrate-create name=create_users_table
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Missing migration name. Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	goose -dir internal/database/migrations create $(name) sql

# DB URL build từ .env
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE)

# Run migrations
migrate-up:
	goose -dir internal/database/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir internal/database/migrations postgres "$(DB_URL)" down

# Start containers
up:
	$(DOCKER_COMPOSE) up -d --build

# Stop containers
down:
	$(DOCKER_COMPOSE) down

# Rebuild only
build:
	$(DOCKER_COMPOSE) build

# View logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Restart
restart:
	$(DOCKER_COMPOSE) down && $(DOCKER_COMPOSE) up -d --build

# Clean everything (⚠️ delete DB data)
clean:
	$(DOCKER_COMPOSE) down -v

swag-init:
	@swag init -g cmd/server/main.go -d . -o ./cmd/server/docs --parseDependency --parseInternal