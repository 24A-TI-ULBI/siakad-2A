# Deployment Guide — Siakad 2A
**Stack:** Go Fiber + MongoDB Atlas + GitHub Actions + Alwaysdata  
**Boilerplate:** [gocroot/alwaysdata](https://github.com/gocroot/alwaysdata)

---

## Daftar Isi
1. [Persiapan Lokal](#1-persiapan-lokal)
2. [Setup MongoDB Atlas](#2-setup-mongodb-atlas)
3. [Setup Alwaysdata](#3-setup-alwaysdata)
4. [Setup GitHub Secrets](#4-setup-github-secrets)
5. [Deploy ke Production](#5-deploy-ke-production)
6. [Verifikasi Production](#6-verifikasi-production)
7. [Troubleshooting](#7-troubleshooting)

---

## 1. Persiapan Lokal

### Prasyarat
| Tool | Versi | Cek |
|------|-------|-----|
| Go | ≥ 1.22 | `go version` |
| Git | ≥ 2.x | `git --version` |

### Setup .env Lokal
Buat file `.env` di root project (jangan di-commit):
```env
MONGOSTRING=mongodb+srv://<user>:<password>@cluster0.xxxxx.mongodb.net/?appName=Cluster0
PORT=8080
IP=127.0.0.1
JWT_SECRET=<string acak 32+ karakter>
```

Generate JWT Secret:
```bash
openssl rand -base64 32
```

### Test Build Lokal
```bash
go env -w CGO_ENABLED=0
go mod tidy
go vet ./...
go build -o gocroot .
./gocroot   # test jalan
rm gocroot  # cleanup
```

### Test Koneksi MongoDB
```bash
go run main.go
# Buka browser: http://localhost:8080
# Harus muncul JSON: {"status":"success","data":{...}}
```

---

## 2. Setup MongoDB Atlas

### Buat Cluster (jika belum ada)
1. Login ke [cloud.mongodb.com](https://cloud.mongodb.com)
2. **Create** → pilih **M0 Free Tier**
3. Pilih region terdekat → **Create Cluster**

### Buat Database User
1. Sidebar → **Database Access** → **Add New Database User**
2. Authentication: **Password**
3. Username & password → catat keduanya
4. Role: **Atlas Admin** (atau **readWriteAnyDatabase**)
5. Klik **Add User**

### Whitelist IP
1. Sidebar → **Network Access** → **Add IP Address**
2. Klik **Allow Access from Anywhere** (`0.0.0.0/0`)
   > Diperlukan karena GitHub Actions dan Alwaysdata punya IP dinamis
3. Klik **Confirm**

### Ambil Connection String
1. Klik **Connect** pada cluster
2. Pilih **Drivers** → Driver: **Go**
3. Copy connection string:
   ```
   mongodb+srv://<user>:<password>@cluster0.xxxxx.mongodb.net/?appName=Cluster0
   ```
4. Ganti `<password>` dengan password user yang dibuat tadi

### Ganti Password (jika sudah ada user lama)
1. **Database Access** → **Edit** user
2. **Edit Password** → generate password baru
3. **Update User**

---

## 3. Setup Alwaysdata

### Buat Akun
1. Daftar di [alwaysdata.com/en/register](https://www.alwaysdata.com/en/register/)
2. Pilih plan **Free** (100MB, free for life)
3. Pilih username dengan hati-hati — akan jadi URL: `https://<username>.alwaysdata.net`

### Buat Site Baru
1. Login ke [admin.alwaysdata.com](https://admin.alwaysdata.com)
2. Sidebar → **Web > Sites** → **Add a site**
3. Isi form:

   | Field | Nilai |
   |-------|-------|
   | Name | `siakad` (bebas) |
   | Addresses | `<username>.alwaysdata.net` (default) |
   | Type | **User program** |
   | Command | `/home/<username>/gocroot` |
   | Working directory | `/home/<username>/` |

4. Scroll ke bagian **Environment** → tambahkan variable berikut:

   | Key | Value |
   |-----|-------|
   | `MONGOSTRING` | connection string MongoDB kamu |
   | `JWT_SECRET` | JWT secret yang sama dengan lokal |

   > ⚠️ **Jangan isi `PORT` dan `IP` secara manual.**  
   > Berdasarkan [dokumentasi resmi alwaysdata](https://help.alwaysdata.com/en/web-hosting/languages/go/configuration/), alwaysdata otomatis meng-inject environment variable `IP` dan `PORT` sesuai konfigurasi internal mereka. App harus listen ke nilai tersebut — kode kita sudah melakukan ini via `os.Getenv("PORT")` dan `os.Getenv("IP")` di `helper/helper.go`.

5. Klik **Save**
6. **Catat APPID** dari URL browser:
   ```
   https://admin.alwaysdata.com/site/XXXXX/
   ```
   Angka `XXXXX` = APPID

### Aktifkan SSH Password Login
1. Sidebar → **Remote Access > SSH**
2. Klik **Modify**
3. Centang **"Password login"** → **Save**
4. Catat:
   - **Host**: `ssh-<username>.alwaysdata.com`
   - **Port**: `22`
   - **Username**: username alwaysdata
   - **Password**: password akun alwaysdata

### Generate API Token
1. Klik nama akun di **pojok kanan atas** → **Profile**
2. Scroll ke **Managing Tokens**
3. Klik **"Add a token"** → beri nama `github-actions`
4. **Catat token** — hanya tampil sekali

---

## 4. Setup GitHub Secrets

Pergi ke repo GitHub → **Settings > Secrets and variables > Actions > New repository secret**

Tambahkan 7 secrets berikut:

| Secret Name | Nilai | Cara Dapat |
|-------------|-------|------------|
| `sshhost` | `ssh-<username>.alwaysdata.com` | Halaman SSH di alwaysdata |
| `sshusername` | username alwaysdata | Username akun |
| `sshpassword` | password alwaysdata | Password akun |
| `sshport` | `22` | Default SSH port |
| `apikey` | token dari Managing Tokens | Langkah Generate API Token |
| `appid` | angka ID dari URL site | URL halaman Modify site |
| `folder` | `/home/<username>/` | Home folder di alwaysdata |

---

## 5. Deploy ke Production

### Trigger Deployment
Deployment otomatis berjalan setiap kali ada push atau merge ke branch `main`:

```bash
git checkout main
git push origin main
```

### Pantau GitHub Actions
1. Buka repo GitHub → tab **Actions**
2. Klik workflow run terbaru
3. Pantau setiap step:
   - ✅ Get latest code
   - ✅ Setup Go
   - ✅ Set CGO_ENABLED=0
   - ✅ Install dependencies
   - ✅ Build and chmod
   - ✅ Copy binary via SCP
   - ✅ SSH restart

### Apa yang Terjadi di Balik Layar
```
Push ke main
    ↓
GitHub Actions build binary gocroot (linux/amd64, CGO_ENABLED=0)
    ↓
SCP binary ke /home/<username>/ di server alwaysdata
    ↓
SSH: curl POST ke API alwaysdata untuk restart site
    ↓
Alwaysdata menjalankan /home/<username>/gocroot
    ↓
App live di https://<username>.alwaysdata.net
```

---

## 6. Verifikasi Production

Setelah workflow selesai, test endpoint berikut:

```bash
# Homepage — harus return JSON status success
curl https://<username>.alwaysdata.net/

# IP Server
curl https://<username>.alwaysdata.net/ip
```

Expected response:
```json
{
  "status": "success",
  "data": {
    "name": "Portal Informasi Akademik Kampus",
    "version": "1.0.0",
    "status": "Server is running"
  }
}
```

---

## 7. Troubleshooting

### Build Gagal di GitHub Actions
```
# Cek error di tab Actions
# Pastikan go.mod tidak ada dependency yang hilang
go mod tidy
git add go.mod go.sum
git commit -m "fix: update dependencies"
git push origin main
```

### App Tidak Jalan Setelah Deploy
1. Cek log di alwaysdata: **Web > Sites > Logs**
2. Pastikan Command benar: `/home/<username>/gocroot`
3. Pastikan environment variable `PORT=8080` sudah diset
4. Pastikan binary ada: SSH ke server → `ls -la ~/gocroot`

### Koneksi MongoDB Gagal
1. Cek `MONGOSTRING` di environment variable alwaysdata
2. Pastikan IP `0.0.0.0/0` sudah di-whitelist di MongoDB Atlas Network Access
3. Test connection string lokal dulu sebelum deploy

### SSH/SCP Gagal di Actions
1. Pastikan password login SSH sudah diaktifkan di alwaysdata
2. Cek semua GitHub Secrets sudah terisi dengan benar
3. Pastikan `folder` diakhiri dengan `/` → `/home/<username>/`

### Port Conflict (Prefork Mode)
Project ini menggunakan `Prefork: true` — alwaysdata akan mengelola proses secara otomatis. Jika ada error terkait port, pastikan tidak ada proses lain yang berjalan di port yang sama.

---

## Referensi
- [Boilerplate gocroot/alwaysdata](https://github.com/gocroot/alwaysdata)
- [Alwaysdata Dashboard](https://admin.alwaysdata.com)
- [MongoDB Atlas](https://cloud.mongodb.com)
- [Go Fiber Docs](https://gofiber.io)
- [GitHub Actions Docs](https://docs.github.com/actions)
