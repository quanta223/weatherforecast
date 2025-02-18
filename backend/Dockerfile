FROM golang:1.23-alpine

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

EXPOSE 3000

CMD ["./main"]