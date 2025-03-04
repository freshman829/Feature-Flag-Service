# Use official Golang image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main ./main.go

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
