cat > Makefile << 'EOF'
.PHONY: help run build test clean docker-up docker-down migrate-up migrate-down ent-gen install-ent

help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make ent-gen      - Generate Ent code"
	@echo "  make install-ent  - Install Ent CLI"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make deps         - Install dependencies"

run:
	go run cmd/server/main.go

build:
	go build -o server cmd/server/main.go

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
EOF

echo "✅ Makefile updated!"