# SIAKAD — Modul 4: Jadwal & Ruangan

Nama: Alifya Azzahra  
NPM: [Isi NPM Kamu]  
Mata Kuliah: Pemrograman Web Service  
Modul: 4 — Jadwal & Ruangan  

---

# Deskripsi

Aplikasi web fullstack berbasis REST API untuk pengelolaan data jadwal perkuliahan dan ruangan kelas di kampus.  
Dibangun menggunakan Go Fiber v2 sebagai backend dan Vanilla HTML/CSS/JavaScript sebagai frontend dashboard yang modern.

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

# Struktur Folder Modul

```bash
siakad-2A/
├── controller/
│   └── jadwalController.go     # Logic CRUD Jadwal & Ruangan
│
├── model/
│   └── jadwal.go               # Struct model data MongoDB
│
├── url/
│   ├── jadwalRoute.go          # Registrasi Group Endpoint /api
│   └── url.go                  # Integrasi route utama
│
├── frontend/
│   └── jadwal/
│       └── index.html          # UI Dashboard interaktif
```

---

# Cara Menjalankan Secara Lokal

## 1. Install Dependency
```bash
go mod tidy
```

## 2. Jalankan Server
```bash
go run main.go
```

## 3. Akses Halaman
* Halaman Portal Utama: `http://127.0.0.1:8080/`
* Halaman Modul Jadwal: `http://127.0.0.1:8080/jadwal/`

---

# Dokumentasi API (Base URL: http://127.0.0.1:8080)

## Endpoint Jadwal Kuliah

### 1. GET `/api/jadwal`
Mengambil semua data jadwal kuliah (bisa difilter berdasarkan query parameter `prodi` dan `dosen`).
* **Query Params (Opsional):** `prodi=Teknik` atau `dosen=Budi`
* **Response Sukses (200 OK):**
  ```json
  {
    "status": "success",
    "message": "Jadwal berhasil diambil",
    "data": [
      {
        "_id": "60c72b2f9b1d8b2c8c8b4567",
        "mata_kuliah": "Pemrograman Web Service",
        "dosen": "Romi",
        "prodi": "D4 Teknik Informatika",
        "hari": "Selasa",
        "jam_mulai": "08:00",
        "jam_selesai": "09:40",
        "ruangan_kode": "LAB-312"
      }
    ]
  }
  ```

### 2. POST `/api/jadwal`
Menambahkan data jadwal kuliah baru.
* **Request Body:**
  ```json
  {
    "mata_kuliah": "Kriptografi",
    "dosen": "Indra",
    "prodi": "D4 Teknik Informatika",
    "hari": "Rabu",
    "jam_mulai": "10:00",
    "jam_selesai": "11:40",
    "ruangan_kode": "LAB-315"
  }
  ```

### 3. GET `/api/jadwal/:id`
Mengambil data jadwal spesifik berdasarkan ObjectID.

### 4. PUT `/api/jadwal/:id`
Memperbarui data jadwal berdasarkan ObjectID.

### 5. DELETE `/api/jadwal/:id`
Menghapus data jadwal berdasarkan ObjectID.

---

## Endpoint Ruangan Kelas

### 1. GET `/api/ruangan`
Mengambil semua daftar ruangan kelas.

### 2. POST `/api/ruangan`
Menambahkan ruangan kelas baru.
* **Request Body:**
  ```json
  {
    "nama": "Lab Vokasi 312",
    "kode": "LAB-312",
    "kapasitas": 35,
    "gedung": "School of Information Technology",
    "fasilitas": ["AC", "Proyektor", "Komputer"]
  }
  ```

### 3. GET `/api/ruangan/:kode`
Mengambil data ruangan spesifik berdasarkan kode ruangan.

### 4. PUT `/api/ruangan/:kode`
Memperbarui data ruangan berdasarkan kode ruangan.

### 5. DELETE `/api/ruangan/:kode`
Menghapus data ruangan berdasarkan kode ruangan.

---

# Fitur Modul 4

✅ **CRUD Jadwal Kuliah**: Tambah, edit, hapus, dan cari jadwal kuliah secara dinamis.  
✅ **CRUD Ruangan Kelas**: Kelola ruangan kelas beserta kapasitas, gedung, dan fasilitas.  
✅ **Validasi ObjectID**: Penanganan validasi ObjectID hex secara aman.  
✅ **Fallback Static Routing**: Pencegahan tabrakan route statis `/jadwal/index.html` dengan API.  
✅ **Responsive & Beautiful Dashboard UI**: Desain antarmuka premium dengan animasi transisi yang halus dan dark mode elements.  
✅ **API Gateway /api**: Route API terorganisir di bawah prefix `/api` yang bersih.  