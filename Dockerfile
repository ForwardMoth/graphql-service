FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server .

FROM alpine:latest AS runner

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./server"]