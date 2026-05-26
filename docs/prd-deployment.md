# PRD: Gocroot Alwaysdata — Pre-Deploy Setup
**Version:** 1.0.0  
**Target Executor:** AI Agent / Claude Code  
**Stack:** Go + Fiber + MongoDB + GitHub Actions + Alwaysdata  

---

## 1. Overview

Dokumen ini mendefinisikan seluruh task pra-deploy yang harus dieksekusi oleh AI agent untuk mempersiapkan project berbasis [gocroot boilerplate](https://github.com/gocroot/alwaysdata) agar siap di-deploy ke platform **alwaysdata.com** menggunakan CI/CD via GitHub Actions.

### 1.1 Tujuan

- Menyiapkan struktur project yang bersih dan siap build
- Memvalidasi semua dependency dan konfigurasi
- Menghasilkan file konfigurasi yang lengkap dan benar
- Memastikan pipeline GitHub Actions dapat berjalan tanpa error

### 1.2 Batasan Scope

> Agent **HANYA** mengerjakan setup & pra-deploy di sisi lokal dan konfigurasi file.  
> Agent **TIDAK** melakukan: push ke GitHub, akses SSH ke server, atau menyentuh kredensial production.

---

## 2. Prerequisites (Wajib Ada Sebelum Agent Mulai)

Agent harus memverifikasi semua item berikut tersedia di environment sebelum mengeksekusi task apapun.

| Item | Perintah Verifikasi | Expected Output |
|------|---------------------|-----------------|
| Go 1.22.x | `go version` | `go1.22.x` |
| Git | `git --version` | `git version 2.x` |
| Akses internet | `curl https://icanhazip.com/` | IP address |
| Direktori project | `ls go.mod` | File exists |

Jika salah satu gagal, **agent harus berhenti dan melaporkan** item yang tidak terpenuhi.

---

## 3. Task List

### TASK-01 — Validasi Struktur Folder Project

**Deskripsi:** Pastikan semua folder dan file wajib dari boilerplate tersedia.

**Aksi:**
```
Cek keberadaan file/folder berikut:
- main.go
- go.mod
- go.sum
- config/
- config/api.go
- config/config.go
- config/cors.go
- config/db.go
- config/token.go
- controller/
- controller/controller.go
- controller/whatsauth.go
- controller/auth.go
- helper/
- helper/helper.go
- helper/api.go
- helper/mongodb.go
- helper/whatsauth.go
- model/
- model/model.go
- model/mongodb.go
- model/whatsauth.go
- url/
- url/url.go
- .github/workflows/
```

**Kriteria Sukses:** Semua path di atas exist.  
**Jika Gagal:** Buat file/folder yang hilang berdasarkan template boilerplate. Laporkan file apa yang dibuat.

---

### TASK-02 — Setup File `.env`

**Deskripsi:** Buat file `.env` dengan semua environment variable yang dibutuhkan aplikasi.

**Aksi:** Buat file `.env` di root project dengan template berikut. Nilai yang kosong (`""`) wajib **diisi oleh user** sebelum menjalankan lokal.

```env
# === DATABASE ===
MONGOSTRING=

# === SERVER ===
PORT=8080
IP=

# === WHATSAUTH ===
WAQRKEYWORD=
WEBHOOKURL=
WEBHOOKSECRET=
WAPHONENUMBER=

# === AUTH ===
PUBLICKEY=

# === INTERNAL ===
INTERNALHOST=
URLAPIWABUTTON=
```

**Aksi Tambahan:**
- Pastikan `.env` masuk ke `.gitignore` (cek apakah sudah ada, jika belum tambahkan baris `*.env` dan `.env`)
- Buat `.env.example` sebagai copy dari template di atas (tanpa nilai sensitif) untuk referensi tim

**Kriteria Sukses:** File `.env` dan `.env.example` ada di root, `.env` masuk `.gitignore`.

---

### TASK-03 — Validasi & Update `go.mod`

**Deskripsi:** Pastikan module name dan Go version sudah benar, lalu sync semua dependency.

**Aksi:**
```bash
# 1. Cek module name di go.mod — harus "gocroot"
head -1 go.mod

# 2. Cek Go version — harus >= 1.22
grep "^go " go.mod

# 3. Set CGO_ENABLED=0 (required untuk build di alwaysdata)
go env -w CGO_ENABLED=0

# 4. Install dan tidy dependency
go get .
go mod tidy

# 5. Verifikasi tidak ada error
go vet ./...
```

**Kriteria Sukses:** `go mod tidy` dan `go vet ./...` selesai tanpa error.

---

### TASK-04 — Test Build Binary

**Deskripsi:** Pastikan project bisa di-build menjadi binary executable, mensimulasi proses yang sama dengan GitHub Actions.

**Aksi:**
```bash
# Build binary dengan nama "gocroot"
go build -o gocroot .

# Pastikan binary ada dan executable
ls -la gocroot

# Cleanup binary setelah test (jangan commit binary)
rm gocroot
```

**Kriteria Sukses:** Build berhasil, binary `gocroot` terbentuk, tidak ada compile error.  
**Jika Gagal:** Tampilkan output error build lengkap. Analisis penyebabnya dan perbaiki jika memungkinkan (missing import, syntax error, dll).

---

### TASK-05 — Setup GitHub Actions Workflow

**Deskripsi:** Pastikan file workflow CI/CD untuk alwaysdata sudah ada dan valid.

**Aksi:** Buat atau validasi file `.github/workflows/alwaysdata.yml` dengan konten berikut:

```yaml
name: AlwaysData.com Deployment
on:
  push:
    branches:
      - main
jobs:
  web-deploy:
    name: Deploy
    runs-on: ubuntu-latest
    steps:
    - name: Get latest code
      uses: actions/checkout@v4

    - name: Setup Go 1.22.x
      uses: actions/setup-go@v5
      with:
        go-version: '1.22.x'

    - name: Set Env for build CGO_ENABLED=0
      run: go env -w CGO_ENABLED=0

    - name: Install dependencies
      run: go get .

    - name: Build and chmod
      run: |
        go build -o gocroot
        chmod a+x gocroot

    - name: Copy binary file via ssh password
      uses: appleboy/scp-action@v0.1.7
      with:
        host: ${{ secrets.sshhost }}
        username: ${{ secrets.sshusername }}
        password: ${{ secrets.sshpassword }}
        port: ${{ secrets.sshport }}
        source: "gocroot"
        target: ${{ secrets.folder }}

    - name: Check binary file, ipaddress, and restart sites
      uses: appleboy/ssh-action@v1.0.3
      with:
        host: ${{ secrets.sshhost }}
        username: ${{ secrets.sshusername }}
        password: ${{ secrets.sshpassword }}
        port: ${{ secrets.sshport }}
        script: |
          ls -l gocroot
          curl https://icanhazip.com/
          curl -X POST --basic --user "${{ secrets.apikey }}:" \
            https://api.alwaysdata.com/v1/site/${{ secrets.appid }}/restart/
```

**Kriteria Sukses:** File `.github/workflows/alwaysdata.yml` ada dan valid YAML.

---

### TASK-06 — Generate GitHub Secrets Checklist

**Deskripsi:** Buat file dokumentasi berisi daftar GitHub Secrets yang wajib diisi secara manual oleh user di repo GitHub.

**Aksi:** Buat file `DEPLOYMENT_SECRETS.md` di root project dengan konten:

```markdown
# GitHub Secrets — Deployment Checklist

Semua secret berikut wajib diisi di:
**GitHub Repo > Settings > Secrets and variables > Actions > New repository secret**

## SSH Credentials (dari alwaysdata.com)

| Secret Name    | Keterangan                                      | Contoh            |
|----------------|-------------------------------------------------|-------------------|
| `sshhost`      | SSH host alwaysdata                             | `ssh-xxx.alwaysdata.com` |
| `sshusername`  | Username SSH alwaysdata                         | `namaakun`        |
| `sshpassword`  | Password SSH yang sudah di-set di alwaysdata    | `P@ssw0rd!`       |
| `sshport`      | Port SSH (default alwaysdata: 22)               | `22`              |

## Alwaysdata API

| Secret Name | Keterangan                                           | Cara Dapat |
|-------------|------------------------------------------------------|------------|
| `apikey`    | API key dari Customer Area > Profile > Managing Tokens | Generate token di dashboard |
| `appid`     | ID angka dari halaman Web > Sites (ada di URL/title) | Lihat di URL halaman Modify site |

## Server Path

| Secret Name | Keterangan                                     | Contoh         |
|-------------|------------------------------------------------|----------------|
| `folder`    | Home folder target binary di server alwaysdata | `/home/namaakun/` |

## Cara Setup di Alwaysdata

1. Login ke https://admin.alwaysdata.com
2. **Web > Sites > Modify** → catat APPID dari URL
3. Set environment variable di bagian Environment:
   ```
   MONGOSTRING=<your-mongo-uri>
   WAQRKEYWORD=<keyword>
   WEBHOOKURL=https://<appname>.alwaysdata.net/whatsauth/webhook
   WEBHOOKSECRET=<secret>
   WAPHONENUMBER=62811xxxxxx
   ```
4. **Remote Access > SSH > Modify** → aktifkan password-based login
5. **Customer Area > Profile > Managing Tokens** → generate API key

## Status Checklist

- [ ] `sshhost` diisi
- [ ] `sshusername` diisi
- [ ] `sshpassword` diisi
- [ ] `sshport` diisi
- [ ] `apikey` diisi
- [ ] `appid` diisi
- [ ] `folder` diisi
- [ ] Environment variable di alwaysdata sudah diset
```

**Kriteria Sukses:** File `DEPLOYMENT_SECRETS.md` terbuat di root project.

---

### TASK-07 — Validasi CORS Configuration

**Deskripsi:** Pastikan konfigurasi CORS di `config/cors.go` sudah mencakup domain yang dibutuhkan project ini.

**Aksi:**
1. Baca isi `config/cors.go`
2. Tampilkan daftar `origins` yang saat ini terdaftar
3. Tanyakan kepada user: *"Apakah ada domain tambahan yang perlu ditambahkan ke CORS?"*
4. Jika user memberikan domain baru, tambahkan ke slice `origins` di file tersebut

**Format tambahan domain:**
```go
// Tambahkan di dalam var origins = []string{ ... }
"https://domain-baru.com",
```

**Kriteria Sukses:** `config/cors.go` berisi domain yang sesuai kebutuhan project.

---

### TASK-08 — Validasi Route & Controller

**Deskripsi:** Pastikan semua route yang didefinisikan di `url/url.go` memiliki controller yang valid.

**Aksi:**
```bash
# Cek semua import dan fungsi controller terpanggil dengan benar
grep -n "controller\." url/url.go
```

Verifikasi setiap fungsi yang dipanggil di `url/url.go` **benar-benar ada** di file controller yang sesuai:

| Route | Controller File | Fungsi |
|-------|----------------|--------|
| `GET /` | controller/controller.go | `Homepage` |
| `GET /ip` | controller/controller.go | `GetIPServer` |
| `GET /whatsauth/refreshtoken` | controller/whatsauth.go | `RefreshWAToken` |
| `POST /whatsauth/webhook` | controller/whatsauth.go | `WhatsAuthReceiver` |
| `GET /auth/phonenumber/:login` | controller/auth.go | `GetPhoneNumber` |

**Kriteria Sukses:** Semua fungsi controller exist dan tidak ada route yang mengarah ke fungsi yang tidak ada.

---

### TASK-09 — Final Pre-Deploy Report

**Deskripsi:** Setelah semua task selesai, buat laporan ringkas dalam file `PRE_DEPLOY_REPORT.md`.

**Format laporan:**

```markdown
# Pre-Deploy Report
Generated: <tanggal dan waktu eksekusi>

## Summary

| Task | Status | Catatan |
|------|--------|---------|
| TASK-01: Validasi Struktur Folder | ✅ / ❌ | ... |
| TASK-02: Setup .env | ✅ / ❌ | ... |
| TASK-03: Validasi go.mod | ✅ / ❌ | ... |
| TASK-04: Test Build Binary | ✅ / ❌ | ... |
| TASK-05: GitHub Actions Workflow | ✅ / ❌ | ... |
| TASK-06: Secrets Checklist | ✅ / ❌ | ... |
| TASK-07: CORS Configuration | ✅ / ❌ | ... |
| TASK-08: Route & Controller | ✅ / ❌ | ... |

## Next Steps (Manual — Harus Dilakukan oleh User)

1. [ ] Isi nilai di file `.env` dengan kredensial yang sebenarnya
2. [ ] Setup alwaysdata.com: Web > Sites > Modify > set Command & Environment
3. [ ] Isi semua GitHub Secrets sesuai `DEPLOYMENT_SECRETS.md`
4. [ ] Push ke branch `main` untuk trigger deployment otomatis
5. [ ] Monitor GitHub Actions tab untuk memastikan workflow berhasil
6. [ ] Verifikasi app live di `https://<appname>.alwaysdata.net`

## Issues Found

<daftar masalah yang ditemukan selama setup, jika ada>

## Files Created/Modified

<daftar file yang dibuat atau dimodifikasi oleh agent>
```

**Kriteria Sukses:** File `PRE_DEPLOY_REPORT.md` terbuat dan berisi status semua task.

---

## 4. Urutan Eksekusi

Agent wajib menjalankan task dalam urutan berikut. Jika satu task gagal dan tidak bisa diperbaiki otomatis, agent harus **berhenti dan melaporkan** sebelum lanjut ke task berikutnya.

```
TASK-01 → TASK-02 → TASK-03 → TASK-04 → TASK-05 → TASK-06 → TASK-07 → TASK-08 → TASK-09
```

---

## 5. File Output yang Dihasilkan Agent

Setelah semua task selesai, file-file berikut harus ada di root project:

| File | Keterangan |
|------|------------|
| `.env` | Environment variable template (kosong, siap diisi) |
| `.env.example` | Contoh env tanpa nilai sensitif |
| `.gitignore` | Sudah include `.env` |
| `.github/workflows/alwaysdata.yml` | GitHub Actions workflow |
| `DEPLOYMENT_SECRETS.md` | Checklist secrets untuk GitHub |
| `PRE_DEPLOY_REPORT.md` | Laporan hasil pre-deploy agent |

---

## 6. Hal yang TIDAK Boleh Dilakukan Agent

- Commit atau push ke Git repository
- Menyimpan nilai kredensial nyata ke dalam file apapun
- Mengakses atau memodifikasi database production
- Menjalankan binary hasil build di environment production
- Memodifikasi logic bisnis di luar scope konfigurasi

---

## 7. Referensi

- Boilerplate repo: https://github.com/gocroot/alwaysdata
- Alwaysdata dashboard: https://admin.alwaysdata.com
- Go Fiber docs: https://gofiber.io
- MongoDB Go Driver: https://www.mongodb.com/docs/drivers/go/
- WhatsAuth docs: https://whatsauth.my.id/docs/