# Use the official Golang image as a parent image
FROM golang:1.20.10

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build ./cmd/app.go

# Expose the port the application will run on
EXPOSE 8000

# Command to run the application
CMD ["./app"]
