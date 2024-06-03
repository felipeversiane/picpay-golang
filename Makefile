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

.PHONY: dev-test
dev-test: dev-ci
	docker-compose -f docker-compose.dev.yml run api go test -v ./internal/...
	docker-compose -f docker-compose.dev.yml down

.PHONY: dev-e2e-test
dev-e2e-test: dev-ci
	docker-compose -f docker-compose.dev.yml run api go test -v ./e2e/...
	docker-compose -f docker-compose.dev.yml down

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

.PHONY: build-test
build-test: build-ci
	docker-compose -f docker-compose.ci.yml run api go test -v ./internal/...
	docker-compose -f docker-compose.ci.yml down

.PHONY: build-e2e-test
build-e2e-test: build-ci
	docker-compose -f docker-compose.ci.yml run api go test -v ./e2e/...
	docker-compose -f docker-compose.ci.yml down

.PHONY: build-runapi
build-runapi:
	go run cmd/api/main.go
