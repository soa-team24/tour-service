FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8081
ENTRYPOINT ["go", "run", "main.go"]
