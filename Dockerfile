# --- STAGE 1 ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

# --- STAGE 2 ---
FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app .
COPY migration ./migration
COPY .env .env

EXPOSE 8080
CMD ["./app"]
