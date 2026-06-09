# SIAKAD — Modul 3: Mata Kuliah & KRS

Nama: Affifah Putri Deza  
NPM: 714240014  
Mata Kuliah: Pemrograman Web Service  
Modul: 3 — Mata Kuliah & KRS  

---

# Deskripsi

Aplikasi web fullstack berbasis REST API untuk pengelolaan data mata kuliah dan pengisian KRS mahasiswa.  
Dibangun menggunakan Go Fiber v2 sebagai backend dan Vanilla HTML/CSS/JavaScript sebagai frontend.

---

# Tech Stack

| Layer | Teknologi |
|---|---|
| Backend | Go + Go Fiber v2 |
| Database | MongoDB Atlas |
| Frontend | HTML + CSS + JS (Vanilla) |
| Hosting | Alwaysdata (Free for Life) |
| CI/CD | GitHub Actions |
| Boilerplate | github.com/gocroot/alwaysdata |

---

# Struktur Folder

```bash
714240014/
├── controller/
│   ├── controller.go
│   └── matkulController.go
│
├── model/
│   ├── model.go
│   └── matkul.go
│
├── url/
│   ├── matkulRoute.go
│   └── url.go
│
├── frontend/
│   └── matkul/
│       └── index.html
│
├── config/
├── helper/
├── main.go
├── go.mod
├── go.sum
└── .env
```

---

# Alur Pengerjaan

```mermaid
flowchart TD
    A([Mulai]) --> B[Clone Repository SIAKAD]
    B --> C[Checkout Branch matkul]
    C --> D[Menjalankan Project Go Fiber]
    D --> E[Membuat Model Matkul]
    E --> F[Membuat Controller Matkul]
    F --> G[Membuat Route Matkul]
    G --> H[Integrasi Route ke url.go]
    H --> I[Membuat Frontend Matkul]
    I --> J[Membuat Fitur Tambah Data]
    J --> K[Membuat Fitur Edit Data]
    K --> L[Membuat Fitur Delete Data]
    L --> M[Testing CRUD Matkul]
    M --> N[Git Add & Commit]
    N --> O[Push Branch matkul]
    O --> P[Create Pull Request]
    P --> Q([Selesai])
```

---

# Environment Variables

Buat file `.env` di root project:

```env
MONGOSTRING=mongodb+srv://username:password@cluster.mongodb.net/
MONGODB_NAME=kampus
PORT=8080
IP=127.0.0.1
```

---

# Cara Menjalankan

## Install Dependency

```bash
go mod tidy
```

---

## Jalankan Server

```bash
go run main.go
```

---

## Buka Browser

```bash
http://127.0.0.1:8080/matkul
```

---

# Dokumentasi API

## Base URL

```bash
http://127.0.0.1:8080
```

---

# Endpoint Matkul Mata Kuliah

## GET /api/matkul

Mengambil semua data mata kuliah.

### Response

```json
[
  {
    "kode": "IF101",
    "nama": "Pemrograman Web",
    "dosen": "Pak Budi",
    "hari": "Senin",
    "jam": "08:00 - 10:00",
    "ruangan": "Lab Komputer 1",
    "sks": 3,
    "semester": 2
  }
]
```

---

## POST /api/matkul

Menambahkan data mata kuliah baru.

### Request Body

```json
{
  "kode": "IF106",
  "nama": "Networking",
  "dosen": "Pak Andi",
  "hari": "Rabu",
  "jam": "10:00 - 12:00",
  "ruangan": "Lab Jaringan",
  "sks": 3,
  "semester": 1
}
```

---

## PUT /api/matkul/:kode

Mengupdate data mata kuliah berdasarkan kode.

### Request Body

```json
{
  "kode": "IF101",
  "nama": "Rekayasa Perangkat Lunak",
  "dosen": "Bu Sinta",
  "hari": "Kamis",
  "jam": "13:00 - 15:00",
  "ruangan": "Lab Software",
  "sks": 2,
  "semester": 1
}
```

---

## DELETE /api/matkul/:kode

Menghapus data mata kuliah berdasarkan kode.

---

# Frontend

| Halaman | URL | Fungsi |
|---|---|---|
| Matkul | /matkul | Kelola data mata kuliah & pengisian KRS |

---

# Fitur

✅ CRUD Data Mata Kuliah  
✅ Tambah Mata Kuliah  
✅ Edit Mata Kuliah  
✅ Delete Mata Kuliah  
✅ Pengisian KRS Mahasiswa  
✅ Search Mata Kuliah  
✅ Responsive UI  
✅ REST API Fiber v2  
✅ Frontend Interaktif  