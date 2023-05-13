# Use the official Golang image for building the application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the project files into the container
COPY . .

# Build the application
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service ./cmd

# Create an image to run the application based on a lightweight Alpine image
FROM alpine:latest

# Set the working directory and copy the executable file from the previous image
WORKDIR /app
COPY --from=builder /app/auth_service /app/

# Add CA certificates for proper HTTPS functionality
RUN apk --no-cache add ca-certificates

# Expose the port for accessing the application
EXPOSE 8080

# Run the application
CMD ["./auth_service"]
