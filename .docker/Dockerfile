# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/swaggo/files
RUN go install github.com/swaggo/gin-swagger
RUN go install github.com/swaggo/swag

# Set the environment variables
ENV GOBIN=/go/bin
ENV PATH=$PATH:$GOBIN

# Copy the rest of the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the application's port
EXPOSE 8080

# Run the compiled binary
CMD ["/app/main"]
