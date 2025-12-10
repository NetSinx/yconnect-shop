# Y-Connect Shop

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Angular](https://img.shields.io/badge/angular-%23DD0031.svg?style=for-the-badge&logo=angular&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

> **Y-Connect Shop** adalah aplikasi e-commerce sederhana yang dibangun untuk mendemonstrasikan implementasi arsitektur **Microservices** menggunakan Go dan Angular.

## 📂 Struktur Project
Project ini dibagi menjadi dua bagian utama (Monorepo):

* **`/server`**: Backend services yang dibangun menggunakan **Go (Golang)**.
* **`/app`**: Frontend aplikasi yang dibangun menggunakan **Angular**.

## 🚀 Fitur Utama
* Arsitektur Microservices.
* Pemisahan Frontend (Client) dan Backend (Server).
* Implementasi RESTful API dengan Go.
* Antarmuka pengguna yang reaktif dengan Angular.
* Dukungan containerization dengan Docker.

## 🛠️ Teknologi yang Digunakan

### Backend
* **Bahasa:** Go (Golang)
* **Arsitektur:** Microservices

### Frontend
* **Framework:** Angular
* **Bahasa:** TypeScript, HTML, CSS/SCSS

### DevOps & Tools
* **Docker:** Untuk containerization aplikasi.

## 💻 Cara Menjalankan (Local Development)

Ikuti langkah-langkah di bawah ini untuk menjalankan aplikasi di komputer lokal kamu.

### Prasyarat
Pastikan kamu sudah menginstall:
* [Go](https://go.dev/dl/) (versi terbaru)
* [Node.js & npm](https://nodejs.org/)
* [Angular CLI](https://angular.io/cli) (`npm install -g @angular/cli`)
* *(Opsional)* [Docker Desktop](https://www.docker.com/products/docker-desktop)

### 1. Clone Repository
```bash
git clone https://github.com/NetSinx/yconnect-shop.git
cd yconnect-shop
```

### 2. Menjalankan Backend (Server)
Buka terminal baru, lalu masuk ke folder server:

```bash
cd server
# Install dependencies (jika menggunakan Go modules)
go mod tidy

# Jalankan server
go run main.go
# Catatan: Sesuaikan nama file utama jika bukan main.go
```

### 3. Menjalankan Frontend (App)
Buka terminal **baru** (biarkan backend berjalan), lalu masuk ke folder app:

```bash
cd app
# Install dependencies
npm install

# Jalankan development server
ng serve
```
Buka browser dan akses `http://localhost:4200/`. Aplikasi akan otomatis me-reload jika kamu mengubah source code.

### 4. Menjalankan dengan Docker (Opsional)
Jika kamu ingin menjalankan menggunakan Docker (pastikan Docker Daemon sudah berjalan):

```bash
# Build dan jalankan container (sesuaikan perintah build docker kamu)
docker-compose up --build
```

## 🤝 Kontribusi
Kontribusi selalu terbuka! Jika kamu ingin meningkatkan project ini untuk pembelajaran:
1.  Fork repository ini.
2.  Buat branch fitur baru (`git checkout -b fitur-keren`).
3.  Commit perubahan kamu (`git commit -m 'Menambahkan fitur keren'`).
4.  Push ke branch (`git push origin fitur-keren`).
5.  Buat Pull Request.

## 📝 Lisensi
[MIT License](LICENSE)
