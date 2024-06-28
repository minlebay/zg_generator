# Use the official Golang image as a build stage
FROM golang:1.22-alpine AS builder

# Set the maintainer label
LABEL maintainer="ilshatminnibaev@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/config.yaml .

# Expose port 8888 and 50051 to the outside world
EXPOSE 8888
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
