.PHONY: up
up:
	docker compose -f docker-compose.ci.yml up -d

.PHONY: down
down:
	docker compose -f docker-compose.ci.yml down

.PHONY: ci
ci:
	docker compose -f docker-compose.ci.yml up -d --build api

.PHONY: runapi
runapi:
	go run cmd/api/main.go
