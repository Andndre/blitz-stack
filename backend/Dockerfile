# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod ./
# COPY go.sum ./ # You will need this later when you have dependencies

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .
RUN go build -o seeder ./cmd/seeder/main.go

# Stage 2: Runner
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/seeder .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
