# Use an official Go image as the base image
FROM golang:1.22.9

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o /talon-server

# Command to run the executable
CMD ["/talon-server"]
