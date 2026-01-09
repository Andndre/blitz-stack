.PHONY: db dev-backend dev-frontend dev production stop seed seed-local sync-types test-all

# Jalankan database dan adminer GUI
db:
	docker-compose up -d db adminer

# Jalankan backend di host
dev-backend:
	cd backend && DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_PASSWORD=password DB_NAME=blitz_db go run main.go

# Jalankan frontend di host
dev-frontend:
	cd frontend && bun dev

# Jalankan database, backend, dan frontend secara bersamaan
dev: db
	@$(MAKE) -j 2 dev-backend dev-frontend

# Build dan jalankan container untuk produksi
production:
	docker-compose up -d --build

# Matikan database
stop:
	docker-compose down

# Seed data (Run inside Docker container)
seed:
	docker-compose exec backend ./seeder

# Seed data local (Run on host)
seed-local:
	cd backend && DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_PASSWORD=password DB_NAME=blitz_db go run cmd/seeder/main.go

# Sinkronisasi tipe data (Go -> TypeScript)
sync-types:
	cd backend && go run github.com/gzuidhof/tygo@latest generate

# Jalankan seluruh testing
test-all:
	cd backend && go test ./...
	cd frontend && bun run test
