FROM golang:1.25 AS builder
WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y dos2unix
RUN dos2unix scripts/build-application.sh

RUN chmod +x scripts/build-application.sh
RUN ./scripts/build-application.sh
# EXPOSE doesn't actually do anything!
EXPOSE 3000
CMD ["/app/ordersystem"]