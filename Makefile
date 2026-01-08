.PHONY: db dev-backend dev-frontend dev stop seed

# Jalankan database saja
db:
	docker-compose up -d db

# Jalankan backend di host
dev-backend:
	cd backend && DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_PASSWORD=password DB_NAME=dealer_heronusa go run main.go

# Jalankan frontend di host
dev-frontend:
	cd frontend && bun dev

# Matikan database
stop:
	docker-compose down

# Seed data (Run inside Docker container)
seed:
	docker-compose exec backend ./seeder

# Seed data local (Run on host)
seed-local:
	cd backend && DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_PASSWORD=password DB_NAME=dealer_heronusa go run cmd/seeder/main.go