# Download alpine image, copy the project and download the dependencies
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Run the docker
FROM alpine:latest
RUN apk add --no-cache sqlite
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
