#!/bin/sh
set -e
cd /app
go mod download
go get ordersystem
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem