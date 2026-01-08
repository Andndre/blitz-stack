# Blitz Stack Starter

Project ini adalah template fullstack modern yang menggabungkan performa **Go** di backend, fleksibilitas **Next.js** di frontend, kecepatan runtime **Bun**, dan kehandalan **PostgreSQL**. Dirancang untuk pengembangan aplikasi web yang cepat, scalable, dan type-safe.

## 🚀 Tech Stack

- **Frontend:** [Next.js 15](https://nextjs.org/) (App Router) + [Bun](https://bun.sh/) + [Tailwind CSS](https://tailwindcss.com/)
- **Backend:** [Go 1.24](https://go.dev/) (Standard Library + Clean Architecture)
- **Database:** [PostgreSQL 15](https://www.postgresql.org/)
- **Infrastructure:** [Docker Compose](https://docs.docker.com/compose/)
- **Testing:** [Bun Test](https://bun.sh/docs/cli/test) & [Go Test](https://go.dev/doc/tutorial/add-a-test)
- **Type Sync:** [Tygo](https://github.com/gzuidhof/tygo) (Go Structs -> TypeScript Interfaces)

## 📋 Prasyarat

Sebelum memulai, pastikan Anda telah menginstal tools berikut:

- [Docker](https://www.docker.com/products/docker-desktop/) & Docker Compose
- [Go](https://go.dev/dl/) (versi 1.20+)
- [Bun](https://bun.sh/) (versi 1.0+)
- [Make](https://www.gnu.org/software/make/) (biasanya sudah ada di Linux/macOS)

## ⚡ Quick Start (Docker)

Cara termudah untuk menjalankan seluruh stack adalah menggunakan Docker Compose.

1. **Jalankan Stack:**

   ```bash
   docker-compose up -d --build
   ```

2. **Isi Data Awal (Seeding):**

   ```bash
   make seed
   ```

   _Perintah ini akan mengisi database dengan data dummy._

3. **Akses Aplikasi:**

   - **Frontend:** [http://localhost:3000](http://localhost:3000)
   - **Backend API:** [http://localhost:8080](http://localhost:8080)
   - **Database:** Port `5432` (User: `user`, Pass: `password`, DB: `blitz_db`)

4. **Hentikan Aplikasi:**
   ```bash
   make stop
   ```

## 💻 Local Development

Jika Anda ingin menjalankan Backend atau Frontend secara lokal (di host machine) untuk development yang lebih cepat (misal: hot-reload), ikuti langkah ini:

### 1. Jalankan Database

Gunakan Docker hanya untuk database PostgreSQL:

```bash
make db
```

### 2. Jalankan Backend (Go)

Di terminal baru:

```bash
make dev-backend
```

_Backend akan berjalan di `http://localhost:8080` dan terhubung ke database lokal._

### 3. Jalankan Frontend (Next.js)

Di terminal baru lainnya:

```bash
make dev-frontend
```

_Frontend akan berjalan di `http://localhost:3000`._

### 4. Setup Data Lokal

Jika Anda menjalankan backend secara lokal, gunakan perintah seed lokal:

```bash
make seed-local
```

## 📁 Struktur Folder

```text
├── backend/
│   ├── cmd/            # Entry points aplikasi (server, seeder)
│   ├── internal/       # Private application code
│   │   ├── database/   # Koneksi dan setup DB
│   │   ├── handlers/   # HTTP handlers (Controllers)
│   │   └── middleware/ # HTTP middlewares
│   ├── models/         # Struct data (Shared dengan frontend via Tygo)
│   └── main.go         # Entry point utama server
├── frontend/
│   ├── app/            # Next.js App Router pages
│   ├── components/     # React components
│   ├── types/          # TypeScript types (generated)
│   └── package.json    # Dependencies frontend
├── Makefile            # Shortcut perintah-perintah penting
└── docker-compose.yml  # Konfigurasi container Docker
```

## 🛠️ Perintah Make yang Tersedia

File `Makefile` menyediakan shortcut untuk tugas-tugas umum:

| Command             | Deskripsi                                                                 |
| :------------------ | :------------------------------------------------------------------------ |
| `make db`           | Menjalankan container database PostgreSQL saja                            |
| `make dev-backend`  | Menjalankan backend Go secara lokal dengan hot-reload (via `go run`)      |
| `make dev-frontend` | Menjalankan frontend Next.js secara lokal dengan Bun |
| `make dev` | Menjalankan database, backend, dan frontend secara bersamaan |
| `make production` | Build dan jalankan seluruh stack dalam Docker container (rebuild images) |
| `make stop` | Menghentikan dan menghapus semua container Docker |
| `make seed`         | Menjalankan seeder data di dalam container backend (saat Docker berjalan) |
| `make seed-local`   | Menjalankan seeder data secara lokal (saat `make dev-backend` digunakan)  |
| `make sync-types`   | Sinkronisasi tipe data dari Go structs ke TypeScript interfaces           |
| `make test-all`     | Menjalankan unit test untuk Backend dan Frontend                          |

## ⚙️ Konfigurasi (Environment Variables)

Aplikasi menggunakan environment variables untuk konfigurasi.

- **Backend:** Secara default dikonfigurasi di `Makefile` untuk local dev (`DB_HOST=localhost`, dll). Di Docker, diatur via `docker-compose.yml`.
- **Frontend:** Menggunakan `NEXT_PUBLIC_API_URL` (client-side) dan `API_URL` (server-side fetching).

## 🔄 Type Synchronization

Salah satu fitur unggulan template ini adalah sinkronisasi tipe data otomatis dari Go ke TypeScript.
Jika Anda mengubah struct di `backend/models/`, jalankan perintah ini untuk memperbarui tipe di frontend:

```bash
make sync-types
```

_Ini akan memperbarui file `frontend/types/index.ts`._

## 🧪 Testing

Jalankan test suite untuk memastikan kode Anda berjalan dengan baik:

```bash
make test-all
```

- **Backend:** Menggunakan `go test` standar + `go-sqlmock` untuk mocking database.
- **Frontend:** Menggunakan `bun test` + `react-testing-library`.

## ❓ Troubleshooting

- **Port Conflict:** Pastikan port `3000`, `8080`, dan `5432` tidak sedang digunakan oleh aplikasi lain.
- **Database Connection Error:** Pastikan container database sudah berjalan (`make db` atau `docker-compose up db`) sebelum menjalankan backend.
- **Bun Error:** Jika `bun install` atau `bun dev` gagal, pastikan Anda menggunakan versi Bun terbaru (`bun upgrade`).
