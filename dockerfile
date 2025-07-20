FROM golang:1.24.3 as builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o frontend .

# Use a minimal base image for the final container
FROM debian:bookworm-slim

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/frontend .

# Copy template files
# These directories should exist in the build context
# created by the GitHub Workflow
COPY templates/ ./templates/
COPY pages/ ./pages/
COPY assets/ ./assets/
COPY config/ ./config/

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./frontend"]