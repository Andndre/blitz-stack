# Dealer Heronusa Stack

Project ini adalah template modern fullstack menggunakan Next.js sebagai frontend, Go sebagai backend, dan PostgreSQL sebagai database. Seluruh aplikasi dikelola menggunakan Docker dengan arsitektur High Performance.

## 🚀 Tech Stack

- **Frontend:** [Next.js](https://nextjs.org/) (App Router) + [Bun](https://bun.sh/)
- **Backend:** [Go](https://go.dev/) (Standard Library + Clean Architecture)
- **Database:** [PostgreSQL 15](https://www.postgresql.org/)
- **Infrastructure:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- **Testing:** [Bun Test](https://bun.sh/docs/cli/test) (Frontend) & [Go Test](https://go.dev/doc/tutorial/add-a-test) (Backend)
- **Type Sync:** [Tygo](https://github.com/gzuidhof/tygo) (Auto-generate TS types from Go structs)

## 🏗️ Architecture & Performance Strategy

Project ini menggunakan strategi 3-Layer Performance Protection:

1.  **Layer 1 (Frontend Cache - ISR):**
    - Next.js menggunakan *Incremental Static Regeneration* (revalidate: 60s).
    - Traffic user biasa (read-only) dilayani instan oleh cache HTML Next.js tanpa menyentuh backend.
    - **Benchmark:** ~850 req/sec (HTML Page).

2.  **Layer 2 (Backend Rate Limiting):**
    - Global Rate Limiter terpasang di level middleware Go.
    - Membatasi 5 request/detik per IP (Burst 10).
    - Melindungi server dari spam/DDoS.
    - **Benchmark:** Menahan beban >8.000 req/sec (Rejected/Protected).

3.  **Layer 3 (Efficient Go Backend):**
    - Logic bisnis dieksekusi oleh Go yang sangat efisien memori.
    - **Benchmark:** >1.500 transaksi valid per detik (Direct DB access).

## 📁 Struktur Folder

```text
dealer-heronusa/
├── backend/
│   ├── cmd/            # Entry points (api, seeder)
│   ├── internal/       # Private code (handlers, database, middleware)
│   └── models/         # Go Structs (Shared Models)
├── frontend/
│   ├── app/            # Next.js App Router (Server Components)
│   ├── components/     # UI Components (Client Components)
│   └── __tests__/      # Frontend Unit Tests
├── docker-compose.yml  # Orchestration
├── Makefile            # Shortcut commands
└── tygo.yaml           # Type generation config
```

## 🛠️ Cara Menjalankan

### 1. Menggunakan Docker (Production Ready)

Jalankan seluruh stack dalam mode background:
```bash
docker-compose up -d --build
```

Akses aplikasi:
- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend: [http://localhost:8080](http://localhost:8080)

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
*Note: Seeding menggunakan binary terpisah (`cmd/seeder`) dan aman dijalankan di production (hanya mengisi jika tabel kosong).*

## 🔄 Sinkronisasi Tipe Data (Go -> TS)

Jika kamu mengubah struct di `backend/models`, perbarui interface TypeScript secara otomatis:
```bash
cd backend && go run github.com/gzuidhof/tygo@latest generate
```

## 📈 Testing

### Backend (Unit Test with Mock DB)
```bash
cd backend && go test ./...
```

### Frontend (Bun Native Test)
```bash
cd frontend && bun run test
```
*Atau gunakan panel Testing di VS Code.*

## 🔐 Keamanan
- **Rate Limiting:** Terpasang secara global.
- **Middleware:** Logika request dipisahkan di `internal/middleware`.
- **Environment:** Konfigurasi via Environment Variables di `docker-compose.yml`.