# Base image with Go installed
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy application source code and configuration
COPY . .

# Install dependencies and build the application
RUN go mod tidy
RUN go build -o verve cmd/main.go

# Copy Redis configuration file (if any)
COPY config/config.yaml /app/config/

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./verve"]
