# Kick&Roar

Football match coordination platform for Dhaka, Bangladesh. Monolithic repo: Go API + React frontend, single deployment.

## Tech stack

- **Backend:** Go 1.24+, Echo, Ent, PostgreSQL 15+ (PostGIS), AWS S3
- **Frontend:** React, Vite, TypeScript (in `web/`)
- **Deploy:** Render (single web service), or Docker Compose locally

## Project layout

```
kicknroar/
├── cmd/server/       # Go entrypoint
├── internal/         # Backend (api, config, database, service, …)
├── web/              # React app (Vite); build output in web/dist (embedded in server)
├── migrations/
├── docs/             # API docs (Swagger)
├── scripts/
├── Dockerfile        # Multi-stage: build web, then Go with embedded SPA
├── docker-compose.yml
├── render.yaml       # Single-service Render deploy
└── go.mod
```

Production: one server serves `/api/v1/*` and the SPA at `/*` (same origin).

## Prerequisites

- Go 1.24+
- Node 20+ and Yarn (for frontend dev and production build)
- PostgreSQL 15+ with PostGIS (or use Docker for DB)

## Local development

### Option A: Docker Compose (app + DB)

```bash
cp .env.example .env
# Set JWT_SECRET and optionally AWS/Google keys
docker-compose up --build
```

- App + API: http://localhost:8000  
- Postgres: localhost:5432 (user `kicknroar`, password `kicknroar123`, db `kicknroar`)

### Option B: Backend and frontend separately

1. **Database:** start Postgres (e.g. `docker-compose up -d postgres`).
2. **Backend:**
   ```bash
   cp .env.example .env
   # Edit .env (DATABASE_URL, JWT_SECRET, …)
   go run ./cmd/server
   ```
   API: http://localhost:8000
3. **Frontend (separate terminal):**
   ```bash
   cd web && yarn install && yarn dev
   ```
   App: http://localhost:3000 (proxies `/api` to backend).

## Production build (monolith)

Frontend must be built before the Go binary so `web/dist` is embedded:

```bash
cd web && yarn install && yarn build && cd ..
go build -o server ./cmd/server
./server
```

Or use the Dockerfile (builds web then Go in one image):

```bash
docker build -t kicknroar .
docker run -p 8000:8000 -e DATABASE_URL=… -e JWT_SECRET=… kicknroar
```

## Deploy to Render

Single web service using the repo Dockerfile:

1. Connect the repo to Render; use `render.yaml` (Blueprint).
2. Set env vars (or use `sync: false` / `generateValue` in the blueprint).
3. Render runs `docker build` then runs the container; one URL for app + API.

Health check: `/api/v1/health`.

## API docs

- Scalar UI: `/docs`
- Swagger JSON: `/api/v1/swagger/*`

## Commands (backend)

- `make run` – run dev server  
- `make build` – build binary (build frontend first for full SPA embed)  
- `make test` – run tests  
- `make ent-gen` – regenerate Ent after schema changes  

## License

Proprietary – All rights reserved.
