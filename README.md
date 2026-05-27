# SIAKAD — Modul 4: Jadwal & Ruangan

Nama: Alifya Azzahra  
NPM: 714240011  
Mata Kuliah: Pemrograman Web Service  
Modul: 4 — Jadwal & Ruangan  

---

# Deskripsi

Aplikasi web fullstack berbasis REST API untuk pengelolaan data jadwal perkuliahan dan ruangan kelas.  
Dibangun menggunakan Go Fiber v2 sebagai backend, MongoDB Atlas sebagai database, dan Vanilla HTML/CSS/JavaScript sebagai frontend.

---

# Tech Stack

| Layer | Teknologi |
|---|---|
| Backend | Go + Go Fiber v2 |
| Database | MongoDB Atlas |
| Frontend | HTML + CSS + JavaScript (Vanilla) |
| Hosting | Alwaysdata |
| API Testing | Postman |
| Version Control | Git & GitHub |

---

# Struktur Folder

```bash
siakad-2A/
├── config/
│   ├── config.go
│   ├── cors.go
│   └── db.go
│
├── controller/
│   ├── controller.go
│   ├── jadwalController.go
│   └── ruanganController.go
│
├── helper/
│   ├── helper.go
│   └── mongodb.go
│
├── model/
│   ├── model.go
│   └── jadwal.go
│
├── url/
│   ├── jadwalRoute.go
│   ├── ruanganRoute.go
│   └── url.go
│
├── frontend/
│   ├── index.html
│   ├── jadwal_ui/
│   │   └── index.html
│   └── ruangan_ui/
│       └── index.html
│
├── main.go
├── go.mod
├── go.sum
└── .gitignore
```

---

# Alur Pengerjaan

```mermaid
flowchart TD

A([Mulai]) --> B[Setup MongoDB Atlas]

B --> C[Konfigurasi .env]

C --> D[Membuat Model]

D --> E[Membuat Controller]

E --> F[Membuat Route]

F --> G[Membuat Frontend]

G --> H[Testing CRUD]

H --> I[Push GitHub]

I --> J[Create Pull Request]

J --> K([Selesai])
```

---

# Environment Variables

Buat file `.env` di root folder dengan konfigurasi berikut:

```env
MONGOSTRING=mongodb+srv://radityarizkir_db_user:pemrogamman_3_webservice_2026@cluster0.xaive9n.mongodb.net/?appName=Cluster0
PORT=8080
IP=127.0.0.1
JWT_SECRET=IjotroPBEj6MQZDZAeGrsvGcOeXs+LzWsFN+8OYQXWk=
```

---

# Cara Menjalankan

### Install Dependency

```bash
go mod tidy
```

### Jalankan Server

```bash
go run main.go
```

Server akan berjalan secara lokal di:
```
http://127.0.0.1:8080/
```

### Buka Browser

Setelah server berjalan, buka browser dan akses URL berikut untuk melihat tampilan aplikasi Anda:
```
http://127.0.0.1:8080/frontend/jadwal_ui/index.html
```
*(Sesuaikan path frontend dengan routing Anda)*

---

# Dokumentasi API

**Base URL**: `http://127.0.0.1:8080`

## Endpoint Jadwal (`/jadwal`)

### GET `/jadwal`
Mengambil semua data jadwal kuliah.

**Response:**
```json
{
  "status": "success",
  "message": "Jadwal berhasil diambil",
  "data": [
    {
      "_id": "6a144dc2...",
      "mata_kuliah": "Pengantar Strategi Algoritma",
      "dosen": "Roni Habibi, S.Kom., M.T.",
      "ruangan_kode": "LAB-315",
      "hari": "Senin",
      "jam_mulai": "13:00",
      "jam_selesai": "18:50",
      "prodi": "D4 Teknik Informatika"
    }
  ]
}
```

### GET `/jadwal/:id`
Mengambil detail jadwal berdasarkan ID.

**Response:**
```json
{
  "status": "success",
  "message": "Jadwal ditemukan",
  "data": {
    "_id": "6a144dc2...",
    "mata_kuliah": "Pengantar Strategi Algoritma",
    "dosen": "Roni Habibi, S.Kom., M.T.",
    "ruangan_kode": "LAB-315",
    "hari": "Senin",
    "jam_mulai": "13:00",
    "jam_selesai": "18:50",
    "prodi": "D4 Teknik Informatika"
  }
}
```

### POST `/jadwal`
Menambahkan data jadwal baru.

**Request Body:**
```json
{
  "mata_kuliah": "Pengantar Strategi Algoritma",
  "dosen": "Roni Habibi, S.Kom., M.T.",
  "ruangan_kode": "LAB-315",
  "hari": "Senin",
  "jam_mulai": "13:00",
  "jam_selesai": "18:50",
  "prodi": "D4 Teknik Informatika"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Jadwal berhasil ditambahkan",
  "data": "6a144dc2..."
}
```

### PUT `/jadwal/:id`
Mengupdate data jadwal berdasarkan ID.

**Request Body:**
```json
{
  "mata_kuliah": "Pengantar Strategi Algoritma (Update)",
  "dosen": "Roni Habibi, S.Kom., M.T.",
  "ruangan_kode": "LAB-315",
  "hari": "Senin",
  "jam_mulai": "13:00",
  "jam_selesai": "18:50",
  "prodi": "D4 Teknik Informatika"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Jadwal berhasil diupdate"
}
```

### DELETE `/jadwal/:id`
Menghapus data jadwal berdasarkan ID.

**Response:**
```json
{
  "status": "success",
  "message": "Jadwal berhasil dihapus"
}
```

## Endpoint Ruangan (`/ruangan`)

### GET `/ruangan`
Mengambil semua data ruangan.

**Response:**
```json
{
  "status": "success",
  "message": "Data ruangan berhasil diambil",
  "data": [
    {
      "_id": "6a144dc2...",
      "nama": "Laboratorium Komputer 315",
      "kode": "LAB-315",
      "kapasitas": 40,
      "fasilitas": [
        "Komputer",
        "AC",
        "Proyektor"
      ],
      "gedung": "Gedung A"
    }
  ]
}
```

### POST `/ruangan`
Menambahkan data ruangan baru.

**Request Body:**
```json
{
  "nama": "Laboratorium Komputer 315",
  "kode": "LAB-315",
  "kapasitas": 40,
  "fasilitas": [
    "Komputer",
    "AC",
    "Proyektor"
  ],
  "gedung": "Gedung A"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Ruangan berhasil ditambahkan",
  "data": "6a144dc2..."
}
```

### GET `/ruangan/:kode`
Mengambil detail ruangan berdasarkan kode ruangan.

**Response:**
```json
{
  "status": "success",
  "message": "Data ruangan ditemukan",
  "data": {
    "_id": "6a144dc2...",
    "nama": "Laboratorium Komputer 315",
    "kode": "LAB-315",
    "kapasitas": 40,
    "fasilitas": [
      "Komputer",
      "AC",
      "Proyektor"
    ],
    "gedung": "Gedung A"
  }
}
```

### PUT `/ruangan/:kode`
Mengupdate data ruangan berdasarkan kode.

**Request Body:**
```json
{
  "nama": "Laboratorium Komputer 315",
  "kode": "LAB-315",
  "kapasitas": 45,
  "fasilitas": [
    "Komputer",
    "AC",
    "Proyektor",
    "Whiteboard"
  ],
  "gedung": "Gedung A"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Ruangan berhasil diupdate"
}
```

### DELETE `/ruangan/:kode`
Menghapus data ruangan berdasarkan kode ruangan.

**Response:**
```json
{
  "status": "success",
  "message": "Ruangan berhasil dihapus"
}
```

---

# Frontend

Aplikasi ini menggunakan antarmuka berbasis HTML, CSS (Vanilla), dan JavaScript yang berinteraksi dengan REST API backend.

| Halaman | Deskripsi / Fungsi Utama |
|---|---|
| **Dashboard** | Halaman utama (`/frontend/index.html`) yang berisi navigasi menuju setiap modul (Jadwal & Ruangan). |
| **Jadwal UI** | Antarmuka untuk melakukan operasi CRUD data Jadwal Kuliah (`/frontend/jadwal_ui/index.html`). |
| **Ruangan UI** | Antarmuka untuk melakukan operasi CRUD data Ruangan (`/frontend/ruangan_ui/index.html`). |

---

# Fitur

- **Manajemen Jadwal Kuliah**:
  - Menambahkan jadwal mata kuliah baru (Create)
  - Menampilkan daftar jadwal perkuliahan (Read)
  - Mengubah detail jadwal perkuliahan (Update)
  - Menghapus jadwal perkuliahan (Delete)
- **Manajemen Ruangan Kelas**:
  - Menambahkan data ruangan baru beserta kapasitasnya (Create)
  - Menampilkan daftar dan detail ruangan kelas (Read)
  - Mengubah kapasitas, fasilitas, dan detail ruangan (Update)
  - Menghapus data ruangan (Delete)
- **RESTful API backend yang responsif**
- **Terintegrasi dengan MongoDB Atlas**
- **Antarmuka pengguna (Frontend) interaktif menggunakan Vanilla JS**
