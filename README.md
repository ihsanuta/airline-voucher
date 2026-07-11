# Airline Voucher Management System

Sistem terintegrasi untuk manajemen voucher kru penerbangan. Aplikasi ini terdiri dari Frontend (React + Vite) dan Backend (Golang).

## 📂 Struktur Proyek
- `/frontend`: Antarmuka aplikasi berbasis React.
- `/backend`: API server berbasis Golang.
- `docker-compose.yml`: Konfigurasi multi-container Docker.
- `Makefile`: Shortcut perintah untuk operasional Docker.

## 🛠 Persyaratan
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- Make (Opsional, untuk mempermudah eksekusi perintah)

## 🚀 Perintah Cepat (Menggunakan Makefile)

| Perintah | Fungsi |
| :--- | :--- |
| `make build` | Membangun ulang (rebuild) seluruh service |
| `make up` | Menjalankan aplikasi di latar belakang |
| `make logs` | Melihat log aplikasi secara real-time |
| `make down` | Menghentikan dan menghapus container |
| `make clean` | Reset total (hapus container & volume data) |

## 🌐 Pengembangan
- **Akses Aplikasi**: Setelah menjalankan `make up`, aplikasi akan tersedia di [http://localhost:8000](http://localhost:8000).
- **Frontend**: Berjalan di port 8000 (dilayani oleh Nginx).
- **Backend**: Berjalan di port 8080 (internal container).

## 💡 Troubleshooting
1. **Host not found in upstream**: Pastikan backend sudah menyala sebelum Nginx mencoba melakukan proxy. Gunakan `depends_on` di `docker-compose.yml`.
2. **CORS Error**: Jika terjadi error CORS, pastikan request dari frontend mengarah ke `/api` (path relatif) sehingga Nginx dapat melakukan *proxying* ke backend.
3. **Database Migration**: Pastikan folder `/db/migrations` sudah tercopy ke dalam container backend di `Dockerfile` backend Anda.
