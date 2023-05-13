FROM golang:1.20 as builder

WORKDIR /app

# Copy files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build the application
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service ./cmd

# An image to run the application based on a lightweight image
FROM alpine:latest

WORKDIR /app
# Copy the executable file from the previous image
COPY --from=builder /app/auth_service /app/

# Add CA certificates for proper HTTPS functionality
RUN apk --no-cache add ca-certificates

# Expose the port for accessing the application
EXPOSE 8080

# Run
CMD ["./auth_service"]
