# Stage 1: Build frontend (web/dist)
FROM node:20-alpine AS frontend-builder

WORKDIR /web

COPY web/package.json web/yarn.lock ./
RUN yarn install --frozen-lockfile

COPY web/ ./
ENV DISABLE_PWA=true
RUN yarn build

# Stage 2: Build Go binary (with embedded web/dist)
FROM golang:1.24-alpine AS go-builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=frontend-builder /web/dist ./web/dist

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# Stage 3: Minimal runtime
FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

WORKDIR /root/

COPY --from=go-builder /app/server .

EXPOSE 8000

CMD ["./server"]
