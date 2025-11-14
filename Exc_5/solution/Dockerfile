FROM golang:1.25 AS builder
WORKDIR /app

COPY . .
RUN chmod +x scripts/build-application.sh
RUN ./scripts/build-application.sh

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ordersystem /app/ordersystem
EXPOSE 3000
CMD ["/app/ordersystem"]