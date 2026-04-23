# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
# This leverages Docker's layer caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .
COPY .env .env

# Build the Go app, creating a static binary
RUN go build -o /SimpleUserRegistry -ldflags="-w -s" ./cmd

# Stage 2: Create the final, minimal, non-root image
FROM gcr.io/distroless/base AS final

WORKDIR /

# Copy the static binary from the builder stage
COPY --from=builder /SimpleUserRegistry /SimpleUserRegistry

# Set the PORT environment variable
ENV PORT=${PORT}

# Expose port 8080
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/SimpleUserRegistry"]