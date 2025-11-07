# Stage 1: build binary
FROM golang:1.25.3 AS builder

# Set working directory ke root module (yang berisi go.mod)
WORKDIR /app

# Copy seluruh proyek backend
COPY . .

# Jalankan go mod tidy
RUN go mod tidy

# Build dari folder cmd/app
WORKDIR /app/cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/cmd/app/main .

RUN chmod +x main

EXPOSE 5000
ENTRYPOINT ["./main"]
