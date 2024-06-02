# Developer commands
.PHONY: dev-up
dev-up:
    docker-compose -f docker-compose.dev.yml up -d

.PHONY: dev-down
dev-down:
    docker-compose -f docker-compose.dev.yml down

.PHONY: dev-ci
dev-ci:
    docker-compose -f docker-compose.dev.yml up -d --build api

.PHONY: dev-runapi
dev-runapi:
    go run cmd/api/main.go

# Build commands
.PHONY: build-up
build-up:
    docker-compose -f docker-compose.ci.yml up -d

.PHONY: build-down
build-down:
    docker-compose -f docker-compose.ci.yml down

.PHONY: build-ci
build-ci:
    docker-compose -f docker-compose.ci.yml up -d --build api

.PHONY: build-runapi
build-runapi:
    go run cmd/api/main.go
