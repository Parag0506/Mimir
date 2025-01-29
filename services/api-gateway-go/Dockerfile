# Stage 1: Build the Go binary
FROM golang:1.19-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o api-gateway main.go

# Stage 2: Create the final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/api-gateway .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./api-gateway"]
