.PHONY: help run run-web build test clean docker-up docker-down migrate-up migrate-down ent-gen install-ent swagger-gen run-with-docs

help:
	@echo "Available commands:"
	@echo "  make run          - Run the backend (Go server)"
	@echo "  make run-web      - Run the web client (Vite dev server at http://localhost:3000)"
	@echo "  make build        - Build the Go binary (use build-full to embed frontend)"
	@echo "  make build-full   - Build web then Go (production binary with SPA)"
	@echo "  make build-web    - Build frontend only (web/dist)"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make ent-gen      - Generate Ent code"
	@echo "  make install-ent  - Install Ent CLI"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make deps         - Install dependencies"
	@echo "  make swagger-gen  - Generate Swagger documentation"
	@echo "  make run-with-docs - Generate docs and run the application"

run:
	go run ./cmd/server

run-web:
	cd web && yarn install && yarn dev

build:
	go build -o server ./cmd/server

# Build frontend then Go (for production binary with embedded SPA)
build-full: build-web build

build-web:
	cd web && yarn install && yarn build

test:
	go test -v ./...

clean:
	rm -f server
	go clean

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

install-ent:
	go get entgo.io/ent/cmd/ent
	go install entgo.io/ent/cmd/ent

ent-gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

migrate-up:
	@echo "Running migrations..."
	@if [ -f migrations/001_initial_schema.sql ]; then \
		psql $(DATABASE_URL) -f migrations/001_initial_schema.sql; \
	else \
		echo "No migration file found"; \
	fi

deps:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run

swagger-gen:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
	@echo "Swagger documentation generated in docs/"

run-with-docs: swagger-gen run
