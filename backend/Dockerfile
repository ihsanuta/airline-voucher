# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Install build tools untuk sqlite3 CGO
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Pastikan CGO_ENABLED=1 untuk sqlite
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Stage 2: Runtime
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]