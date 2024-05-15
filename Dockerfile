FROM golang:alpine AS builder
WORKDIR /app
COPY . .
COPY proto /app/proto
EXPOSE 8082
ENTRYPOINT ["go", "run", "main.go"]
