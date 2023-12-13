# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any dependencies
RUN go mod download

# Build the Go application
RUN go build cmd/main.go

# Expose the port the application runs on
EXPOSE 10000

# Command to run the executable
CMD ["./main"]
