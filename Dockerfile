# Use official Golang image as a build stage
FROM golang:1.23 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy Go modules first (Optimized caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build the Go application
RUN go build -o main ./main.go


# === Final stage ===
FROM golang:1.23 AS final

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Install curl for health checks
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Expose application port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
