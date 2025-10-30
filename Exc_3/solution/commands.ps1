# ===============================================
# Exercise 3 - Order Service + PostgreSQL Setup
# ===============================================

Write-Host "🔨 Building order-service image..."
docker build -t order-service:local .

Write-Host "`n🐘 Starting PostgreSQL container..."
docker stop order-db -t 2 2>$null
docker rm order-db 2>$null
docker run -d `
  --name order-db `
  --env-file debug.env `
  -p 5432:5432 `
  -v order_data:/var/lib/postgresql/data `
  postgres:16

Write-Host "`n🌐 Creating Docker network (if not exists)..."
docker network create order-net 2>$null | Out-Null
docker network connect order-net order-db 2>$null

Write-Host "`n🚀 Starting order-service container..."
docker stop order-service -t 2 2>$null
docker rm order-service 2>$null
docker run -d `
  --name order-service `
  --network order-net `
  -p 3000:3000 `
  -e POSTGRES_DB=order `
  -e POSTGRES_USER=docker `
  -e POSTGRES_PASSWORD=docker `
  -e POSTGRES_TCP_PORT=5432 `
  -e DB_HOST=order-db `
  order-service:local

Write-Host "Containers working!"
# http://localhost:3000
# http://localhost:3000/openapi/index.html
# on terminal: .\run_all.ps1