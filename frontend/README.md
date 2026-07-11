# Frontend: Crew Voucher UI

Aplikasi client-side dibangun dengan React, TypeScript, dan Tailwind CSS.

## 🚀 Fitur Utama
- Form input yang responsif untuk pendaftaran voucher kru.
- Validasi data sisi klien sebelum dikirim ke API.
- Status loading dan error handling yang user-friendly.
- Integrasi otomatis dengan Nginx sebagai Reverse Proxy.

## 🛠 Tech Stack
- Framework: React 18+ (Vite)
- Language: TypeScript
- Styling: Tailwind CSS v4
- State Management: React useState
- API Client: Fetch API

## 📦 Struktur Folder
/frontend
├── /src
│   ├── /components  # Komponen UI (Form, Status, dll)
│   ├── /api         # Service untuk komunikasi dengan Backend
│   ├── /types       # Definisi interface TypeScript
│   └── App.tsx      # Main application component
├── nginx.conf       # Konfigurasi Nginx untuk Reverse Proxy
└── Dockerfile       # Konfigurasi multi-stage build

## ⚙️ Cara Menjalankan (Lokal)
1. Masuk ke folder frontend:
   cd frontend
2. Install dependensi:
   npm install
3. Jalankan development server:
   npm run dev

## 🌐 Konfigurasi
- Environment: Gunakan file .env untuk mengatur VITE_API_BASE_URL (disarankan menggunakan '/api').
- Proxy: Nginx di dalam Docker secara otomatis melakukan proxy dari jalur '/api' ke container backend (port 8080).

## 💡 Troubleshooting
- CORS Error: Pastikan pemanggilan API menggunakan path '/api' agar diproses oleh Nginx, bukan menembak langsung ke port 8080 backend.
- Build Error: Pastikan versi Node.js di environment Anda sesuai dengan versi yang digunakan di Dockerfile (Node 20+).

---
*Dokumentasi ini digunakan untuk mempermudah pengembangan antarmuka pengguna.*