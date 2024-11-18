# Build stage
FROM golang:1.23.3-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -o server cmd/api/main.go

# Final stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 8100
CMD ["/app/server"]
