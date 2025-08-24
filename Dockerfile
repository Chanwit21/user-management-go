# =========================
# Stage 1: Build the app
# =========================
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for Go modules) and curl (optional)
RUN apk add --no-cache git

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go Fiber app (disable CGO for smaller image)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# =========================
# Stage 2: Run the app
# =========================
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/main .

# Expose the port your Fiber app listens on
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
