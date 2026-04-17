.PHONY: up down build logs restart clean swag-init

DOCKER_COMPOSE=docker compose --env-file .env -f docker/docker-compose.yaml

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