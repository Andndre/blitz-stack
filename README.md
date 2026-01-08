# Dealer Heronusa Stack

Project ini adalah template modern fullstack menggunakan Next.js sebagai frontend, Go sebagai backend, dan PostgreSQL sebagai database. Seluruh aplikasi dikelola menggunakan Docker.

## 🚀 Tech Stack

- **Frontend:** [Next.js](https://nextjs.org/) (App Router) + [Bun](https://bun.sh/)
- **Backend:** [Go](https://go.dev/) (Standard Library for High Performance)
- **Database:** [PostgreSQL 15](https://www.postgresql.org/)
- **Infrastructure:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- **Testing:** [Bun Test](https://bun.sh/docs/cli/test) (Frontend) & [Go Test](https://go.dev/doc/tutorial/add-a-test) (Backend)
- **Type Sync:** [Tygo](https://github.com/gzuidhof/tygo) (Auto-generate TS types from Go structs)

## 📁 Struktur Folder

```text
dealer-heronusa/
├── backend/            # Go Backend
│   ├── cmd/            # Entry points (api, seeder)
│   ├── internal/       # Private code (handlers, database, middleware)
│   └── models/         # Go Structs (Models)
├── frontend/           # Next.js Frontend
│   ├── app/            # Next.js App Router
│   ├── components/     # Reusable React components
│   ├── types/          # Generated TypeScript types
│   └── __tests__/      # Frontend Unit Tests
├── docker-compose.yml  # Orchestration
├── Makefile            # Shortcut commands
└── tygo.yaml           # Type generation config
```

## 🛠️ Cara Menjalankan

### 1. Menggunakan Docker (Rekomendasi)

Jalankan seluruh stack dalam mode background:
```bash
docker-compose up -d --build
```

Akses aplikasi:
- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend: [http://localhost:8080](http://localhost:8080)
- Database: `localhost:5432`

### 2. Mode Development (Hybrid)

Jika ingin fitur Hot Reload tanpa rebuild Docker:
1. Jalankan database: `make db`
2. Jalankan backend: `make dev-backend`
3. Jalankan frontend: `make dev-frontend`

## 🧪 Database & Seeding

Data database tersimpan di volume Docker. Untuk mengisi data awal (seeding):
```bash
make seed
```
*Note: Seeding hanya akan berjalan jika tabel dalam keadaan kosong (aman untuk production).*

## 🔄 Sinkronisasi Tipe Data (Go -> TS)

Jika kamu mengubah struct di `backend/models`, kamu bisa memperbarui interface di TypeScript secara otomatis:
```bash
cd backend && go run github.com/gzuidhof/tygo@latest generate
```
Hasilnya akan muncul di `frontend/types/index.ts`.

## 📈 Testing

### Backend
```bash
cd backend && go test ./...
```

### Frontend
```bash
cd frontend && bun test
```

## ⚡ Performa
Berdasarkan benchmark `ab` (Apache Benchmark), backend Go ini mampu menangani **>1.500 request per detik** dengan rata-rata latency **30ms** pada environment local.

## 🔐 Keamanan
- **Rate Limiting:** Terpasang secara global di Backend (5 req/sec per IP).
- **CORS:** Konfigurasi default aman untuk komunikasi internal Docker.
