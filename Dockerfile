# Select Golang Version
FROM golang:1.21.3-alpine

# Setup Directory project in container
WORKDIR /src/app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o ./dist/main ./main.go

# Run the application
CMD ["./dist/main"]
