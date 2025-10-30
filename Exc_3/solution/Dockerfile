  # todo
# syntax=docker/dockerfile:1

FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ordersystem

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/ordersystem .
EXPOSE 3000
CMD ["./ordersystem"]