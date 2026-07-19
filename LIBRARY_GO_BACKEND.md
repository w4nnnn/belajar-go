# 📝 Roadmap & Task List: Membangun REST API Perpustakaan (Go 1.23+)

Dokumen ini adalah daftar tugas terstruktur untuk membangun Backend Manajemen Perpustakaan (Buku) yang *production-ready* menggunakan Go, Gin, PostgreSQL, dan sqlc.
Kamu yang akan menulis kodenya secara mandiri, dan saya akan mengevaluasinya setelah kamu selesai.

## Tujuan Pembelajaran
- Menerapkan arsitektur standar Go (`cmd/`, `internal/`).
- Menulis *query* SQL yang aman tipe menggunakan `sqlc`.
- Membangun REST API menggunakan framework `gin-gonic/gin`.
- Menerapkan koneksi database yang efisien menggunakan `pgxpool`.
- Melakukan validasi *request* dan penanganan *error* yang tepat.

---

## 📌 Tahap 1: Inisialisasi Proyek & Tooling

- [x] **Buat folder proyek dan inisialisasi Go Module.**
  - [x] `mkdir library-api && cd library-api`
  - [x] `go mod init <nama-module-mu>` (misal: `github.com/username/library-api`)
- [x] **Buat struktur folder standar.**
  - [x] Buat folder `cmd/api`
  - [x] Buat folder `internal/config`
  - [x] Buat folder `internal/db`
  - [x] Buat folder `internal/handler`
- [x] **Install *tools* yang dibutuhkan (jika belum ada).**
  - [x] `goose` (untuk migrasi database).
  - [x] `sqlc` (untuk *generate* kode dari SQL).
  - [x] `golangci-lint` (untuk *linting`).

---

## 📌 Tahap 2: Database Setup & Migrasi

- [x] **Buat database PostgreSQL lokal.** (Bisa via Docker atau instalasi langsung).
- [x] **Buat file migrasi menggunakan Goose.**
  - [x] Buat folder `sql/schema`.
  - [x] Buat file migrasi (misal: `001_create_books.sql`).
  - [x] Tulis query SQL `CREATE TABLE books` dengan kolom: 
    - `id` (UUID PRIMARY KEY)
    - `title` (VARCHAR NOT NULL)
    - `author` (VARCHAR NOT NULL)
    - `isbn` (VARCHAR UNIQUE)
    - `is_available` (BOOLEAN DEFAULT true)
    - `created_at` (TIMESTAMP)
    - `updated_at` (TIMESTAMP).
  - [x] Tulis query pembatalan (`DROP TABLE books`).
- [x] **Jalankan migrasi (Goose Up).**
  - [x] Buat file `.env` di *root* proyek (`library-api/.env`) dan isi dengan:
    ```env
    GOOSE_DRIVER=postgres
    GOOSE_DBSTRING=postgres://username:password@localhost:5432/namadatabase?sslmode=disable
    ```
  - [x] Install godotenv: `go install github.com/joho/godotenv/cmd/godotenv@latest`
  - [x] Jalankan migrasi: `godotenv goose -dir sql/schema up`
  - [x] Pastikan tabel berhasil terbuat di database.

---

## 📌 Tahap 3: Konfigurasi SQLC & Generate Kode

- [x] **Buat file *query* SQL.**
  - [x] Buat folder `sql/query`.
  - [x] Buat file `books.sql`.
  - [x] Tulis query dengan anotasi `sqlc` untuk operasi CRUD:
    - CreateBook (`INSERT ... RETURNING *`)
    - GetBook (`SELECT ... WHERE id = $1 LIMIT 1`)
    - ListBooks (`SELECT ... ORDER BY created_at DESC`)
    - UpdateBookAvailability (`UPDATE ... SET is_available = $2 ... RETURNING *`)
    - DeleteBook (`DELETE ...`)
- [x] **Buat file konfigurasi `sqlc.yaml`.**
  - [x] Atur *engine* ke `postgresql`, arahkan folder schema dan query, lalu set *output* ke `internal/db`.
- [x] **Generate kode Go.**
  - [x] Jalankan `sqlc generate`.
  - [x] Periksa apakah file `models.go` dan `books.sql.go` muncul di dalam `internal/db`.

---

## 📌 Tahap 4: Setup Framework Gin & Koneksi Database

- [x] **Install *dependencies*.**
  - [x] `go get -u github.com/gin-gonic/gin`
  - [x] `go get github.com/jackc/pgx/v5/pgxpool`
- [ ] **Buat file utama (`cmd/api/main.go`).**
  - [ ] Buat fungsi `main()`.
  - [ ] Konfigurasi koneksi database menggunakan `pgxpool.New`.
  - [ ] Inisialisasi `db.New(pool)` dari kode *sqlc* yang di-*generate*.
  - [ ] Inisialisasi Gin router (`gin.Default()`).
  - [ ] Buat *endpoint* sederhana untuk *health check* (`GET /ping`).
  - [ ] Jalankan server di port tertentu (misal: `8080`).

---

## 📌 Tahap 5: Implementasi Handler (CRUD API)

- [ ] **Buat file *handler* (`internal/handler/book.go`).**
  - [ ] Buat struct `BookHandler` yang menyimpan referensi ke `*db.Queries`.
  - [ ] Buat fungsi pembuatnya (`NewBookHandler(q *db.Queries) *BookHandler`).
- [ ] **Implementasikan *endpoint* Create (`POST /books`).**
  - [ ] Buat struct untuk menampung *request body* (`CreateBookRequest`).
  - [ ] Gunakan validasi bawaan Gin (contoh: ``binding:"required,min=3,max=255"`` untuk judul dan penulis).
  - [ ] Ikat (bind) JSON request, panggil *method* dari *sqlc* untuk menyimpan ke database.
  - [ ] Kembalikan *response* dengan status HTTP 201 (Created).
- [ ] **Implementasikan *endpoint* Read/List (`GET /books`).**
  - [ ] Ambil semua data dari database, kembalikan dengan status HTTP 200 (OK).
- [ ] **Implementasikan *endpoint* Read/Single (`GET /books/:id`).**
  - [ ] Tangkap parameter `:id` dari URL, cari di database. Tangani *error* jika data tidak ditemukan (HTTP 404).
- [ ] **Implementasikan *endpoint* Update Status (`PATCH /books/:id/status`).**
  - [ ] Buat struct *request* khusus untuk update ketersediaan (`is_available` boolean). Update data dan kembalikan *response*.
- [ ] **Implementasikan *endpoint* Delete (`DELETE /books/:id`).**
  - [ ] Tangkap ID, hapus data dari database, kembalikan status HTTP 204 (No Content) atau konfirmasi sukses.
- [ ] **Daftarkan semua *handler* tersebut di rute Gin pada `cmd/api/main.go`.**

---

## 📌 Tahap 6: Eksekusi & Evaluasi

- [ ] **Jalankan aplikasi:** `go run cmd/api/main.go`.
- [ ] **Lakukan uji coba manual** (menggunakan `curl`, Postman, atau ekstensi VS Code Thunder Client/REST Client).
  - [ ] Coba tambahkan Buku (valid).
  - [ ] Coba tambahkan Buku tanpa nama penulis (harus gagal/400).
  - [ ] Ambil daftar Buku.
  - [ ] Update status Buku menjadi dipinjam (`is_available = false`).
  - [ ] Hapus Buku dari sistem.
- [ ] **Laporkan kepada saya (AI)**.
  - [ ] Jika ada *error* saat membuat, salin pesan *error*-nya kepada saya.
  - [ ] Setelah selesai dan berjalan, kamu bisa menyalin kode `main.go` dan `book.go` kamu, lalu minta saya untuk mereviu kodenya dari segi *best practice*, keamanan, dan efisiensi.

Selamat *ngoding*! Saya tunggu laporannya.