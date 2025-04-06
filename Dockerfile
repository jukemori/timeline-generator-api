FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o timeline-generator ./cmd/server

FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/timeline-generator /app/

# Verify the file exists (for debugging)
RUN ls -la /app/

# Run with the full path
CMD ["/app/timeline-generator"]