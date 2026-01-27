# --- Stage 1: Build ---
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build API binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api-bin ./cmd/api/main.go

# Build Worker/Scheduler binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o worker-bin ./cmd/worker/main.go

# --- Stage 2: Runtime ---
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copy binaries from builder
COPY --from=builder /app/api-bin .
COPY --from=builder /app/worker-bin .

# Expose the API port
EXPOSE 8080
