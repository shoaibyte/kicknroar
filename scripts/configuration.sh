# Create .gitignore
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
server
kicknroar-backend

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# Environment files
.env
.env.local
.env.*.local

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Logs
*.log

# Ent generated code
internal/ent/generated/

# Temporary files
tmp/
temp/
EOF

# Create .env.example
cat > .env.example << 'EOF'
# Server Configuration
PORT=8080
ENV=development
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000

# Database
DATABASE_URL=postgres://user:password@localhost:5432/kicknroar?sslmode=disable

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h

# AWS S3
AWS_ACCESS_KEY_ID=your-aws-access-key
AWS_SECRET_ACCESS_KEY=your-aws-secret-key
AWS_REGION=ap-south-1
AWS_S3_BUCKET=kicknroar-production

# Google Maps
GOOGLE_MAPS_API_KEY=your-google-maps-api-key

# Rate Limiting
RATE_LIMIT_REQUESTS_PER_MINUTE=100
AUTH_RATE_LIMIT_REQUESTS_PER_MINUTE=5
UPLOAD_RATE_LIMIT_REQUESTS_PER_MINUTE=10
EOF

# Create Makefile
cat > Makefile << 'EOF'
.PHONY: help run build test clean docker-up docker-down migrate-up migrate-down ent-gen

help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make ent-gen      - Generate Ent code"
	@echo "  make migrate-up   - Run database migrations"

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

ent-gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema

migrate-up:
	psql $(DATABASE_URL) -f migrations/001_initial_schema.sql

deps:
	go mod download
	go mod tidy
EOF

# Create README.md
cat > README.md << 'EOF'
# Kick&Roar Backend

Backend API for Kick&Roar - Football match coordination platform for Dhaka, Bangladesh.

## Tech Stack

- **Language:** Go 1.24+
- **Framework:** Echo
- **ORM:** Ent
- **Database:** PostgreSQL 15+ with PostGIS
- **Storage:** AWS S3
- **Deployment:** Render

## Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 15+ with PostGIS extension
- AWS account (for S3)
- Google Maps API key

### Setup

1. Clone the repository
```bash
git clone https://github.com/shoaibhassan/kicknroar-backend.git
cd kicknroar-backend
```

2. Install dependencies
```bash
make deps
```

3. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run database migrations
```bash
make migrate-up
```

5. Generate Ent code
```bash
make ent-gen
```

6. Run the application
```bash
make run
```

The API will be available at `http://localhost:8080`

## Development

### Project Structure
```
kicknroar-backend/
├── cmd/server/          # Application entry point
├── internal/
│   ├── api/            # HTTP handlers and routing
│   ├── service/        # Business logic
│   ├── repository/     # Data access layer
│   ├── ent/            # Ent ORM schemas
│   ├── pkg/            # Shared packages
│   └── config/         # Configuration
└── migrations/         # Database migrations
```

### Available Commands

- `make run` - Run development server
- `make build` - Build production binary
- `make test` - Run tests
- `make ent-gen` - Regenerate Ent code after schema changes

## API Documentation

API documentation will be available at `/api/v1/docs` (coming soon)

## Deployment

See `render.yaml` for Render deployment configuration.

## License

Proprietary - All rights reserved
EOF

# Create render.yaml
cat > render.yaml << 'EOF'
services:
  - type: web
    name: kicknroar-api
    env: docker
    region: singapore
    plan: starter
    branch: main

    buildCommand: go build -o server ./cmd/server
    startCommand: ./server
    healthCheckPath: /api/v1/health

    envVars:
      - key: PORT
        value: 8080
      - key: ENV
        value: production
      - key: DATABASE_URL
        fromDatabase:
          name: kicknroar-db
          property: connectionString
      - key: JWT_SECRET
        generateValue: true
      - key: ALLOWED_ORIGINS
        value: https://kicknroar.com
      - key: AWS_ACCESS_KEY_ID
        sync: false
      - key: AWS_SECRET_ACCESS_KEY
        sync: false
      - key: AWS_REGION
        value: ap-south-1
      - key: AWS_S3_BUCKET
        value: kicknroar-production
      - key: GOOGLE_MAPS_API_KEY
        sync: false

databases:
  - name: kicknroar-db
    databaseName: kicknroar
    plan: starter
    postgresMajorVersion: 15
    region: singapore
EOF

# Create Dockerfile
cat > Dockerfile << 'EOF'
# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
EOF

echo "✅ Configuration files created!"