# Backend: Airline Voucher API

API Service untuk sistem manajemen voucher kru, dibangun menggunakan Go dengan framework Echo.

## 🚀 Fitur Utama
- Validation: Pengecekan ketersediaan voucher per penerbangan dan tanggal.
- Voucher Generation: Pembuatan dan penyimpanan data voucher secara acak.
- CORS Handling: Middleware untuk menangani request dari frontend.

## 🛠 Tech Stack
- Language: Go (v1.26.3+)
- Framework: Echo (https://echo.labstack.com/)
- Database: SQLLite
- Migrations: golang-migrate

## 📦 Struktur Folder
/backend
├── /db
│   └── /migrations  # File SQL migrasi database
├── /handlers        # Logic untuk menangani request API
├── /models          # Definisi struktur data
├── main.go          # Entry point aplikasi
└── Dockerfile       # Konfigurasi container

## ⚙️ Cara Menjalankan (Lokal)
1. Masuk ke folder backend:
   cd backend
2. Pastikan file .env sudah diatur.
3. Jalankan server:
   go run main.go

## 📝 API Endpoints

### 1. Check Voucher
Memvalidasi apakah voucher sudah ada untuk penerbangan tertentu.
- Endpoint: POST /api/check
- Request Body:
  {
    "flightNumber": "GA102",
    "date": "11-07-2026"
  }

### 2. Generate Voucher
Membuat voucher baru jika belum tersedia.
- Endpoint: POST /api/generate
- Request Body:
  {
    "name": "User",
    "id": "CREW123",
    "flightNumber": "GA102",
    "date": "11-07-2026",
    "aircraft": "ATR"
  }

## 🔧 Konfigurasi
- Port: 8080
- Database: Pastikan URL database diset melalui environment variable.
- CORS: Jika diakses di luar Docker, pastikan middleware CORS mengizinkan Origin frontend Anda.