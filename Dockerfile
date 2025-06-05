# Multi-stage build for Go backend
FROM golang:1.22-alpine AS builder

# Install git and ca-certificates
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build the application from the correct path
# Since your main.go is in cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# --- Final stage ---
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .

# Make binary executable and change ownership
RUN chmod +x server && chown appuser:appgroup server

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]