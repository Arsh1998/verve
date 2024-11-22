# Stage 1: Build the application
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o verve cmd/main.go

# Stage 2: Run the application
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/verve /app/
COPY config/config.yaml /app/config/
EXPOSE 8080
CMD ["./verve"]
