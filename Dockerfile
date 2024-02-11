# Start from the latest golang base image
FROM golang:1.22.0-alpine3.18

# Add Maintainer Info
LABEL maintainer="Dang Viet Hung <dangviethung096@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]