# Requirements Document

## Introduction

SIAKAD Dashboard adalah sistem informasi akademik kampus berbasis web yang dibangun secara kolaboratif oleh 20 mahasiswa, masing-masing mengerjakan satu modul. Sistem ini terdiri dari backend Go Fiber v2 + MongoDB Atlas yang di-deploy di Alwaysdata, dan frontend Vanilla HTML/CSS/JS yang di-host di GitHub Pages.

Tujuan utama adalah menyediakan dashboard yang scalable, terintegrasi, dan memiliki tampilan yang proper — dengan modul mahasiswa sebagai implementasi pertama dan referensi arsitektur untuk 19 modul lainnya (dosen, matkul, jadwal, nilai, absensi, pengumuman, beasiswa, perpustakaan, prestasi, alumni, ormawa, notifikasi, kuesioner, spp, surat, pkl, skripsi, fasilitas, berita).

Base URL API: `https://dituniverse.alwaysdata.net`
Frontend URL: `https://24a-ti-ulbi.github.io/siakad-2A/`

---

## Glossary

- **SIAKAD**: Sistem Informasi Akademik Kampus — aplikasi web fullstack yang dibangun secara kolaboratif.
- **Dashboard**: Halaman utama (`index.html`) yang menampilkan ringkasan status sistem dan navigasi ke semua modul.
- **Modul**: Unit fungsional yang dikerjakan oleh satu mahasiswa, terdiri dari backend endpoint dan halaman frontend.
- **Shared_Layout**: Komponen HTML/CSS/JS yang digunakan bersama oleh semua halaman modul (header, sidebar, footer, style).
- **Auth_System**: Sistem autentikasi berbasis JWT yang melindungi halaman dan endpoint tertentu.
- **JWT_Token**: JSON Web Token yang dihasilkan saat login, disimpan di `localStorage`, dan dikirim sebagai `Authorization: Bearer <token>` pada setiap request terproteksi.
- **API_Client**: Modul JavaScript (`api.js`) yang menangani semua komunikasi HTTP ke backend, termasuk injeksi JWT header.
- **Guard**: Fungsi JavaScript yang memeriksa keberadaan dan validitas JWT_Token sebelum halaman terproteksi dirender.
- **NPM**: Nomor Pokok Mahasiswa — identifier unik mahasiswa.
- **NIDN**: Nomor Induk Dosen Nasional — identifier unik dosen.
- **Base_URL**: URL dasar backend API, yaitu `https://dituniverse.alwaysdata.net`.
- **GitHub_Pages**: Platform hosting frontend statis di `https://24a-ti-ulbi.github.io/siakad-2A/`.

---

## Requirements

### Requirement 1: Shared Layout dan Komponen Bersama

**User Story:** Sebagai mahasiswa pengembang, saya ingin memiliki shared layout yang bisa dipakai semua modul, sehingga tampilan antar modul konsisten dan saya tidak perlu menulis ulang CSS/HTML dari nol.

#### Acceptance Criteria

1. THE Shared_Layout SHALL menyediakan file `frontend/shared/style.css` yang berisi design system global (variabel warna, tipografi, komponen card, tabel, form, button, alert, badge).
2. THE Shared_Layout SHALL menyediakan file `frontend/shared/layout.js` yang meng-inject elemen sidebar navigasi dan header secara programatik ke setiap halaman yang memuatnya.
3. WHEN halaman modul memuat `layout.js`, THE Shared_Layout SHALL merender sidebar yang berisi daftar semua 20 modul beserta ikon dan link navigasinya.
4. THE Shared_Layout SHALL menyediakan file `frontend/shared/api.js` yang mengekspor fungsi `apiFetch(path, options)` untuk semua request HTTP ke Base_URL.
5. WHEN `apiFetch` dipanggil dan JWT_Token tersedia di `localStorage`, THE API_Client SHALL menyertakan header `Authorization: Bearer <token>` secara otomatis pada setiap request.
6. IF JWT_Token tidak ditemukan di `localStorage` saat `apiFetch` dipanggil pada endpoint terproteksi, THEN THE API_Client SHALL melempar error dengan pesan "Sesi habis, silakan login kembali".
7. THE Shared_Layout SHALL menyediakan komponen notifikasi toast (`showToast(type, message)`) yang dapat dipanggil dari halaman manapun.
8. WHILE pengguna berada di halaman manapun, THE Shared_Layout SHALL menampilkan nama modul aktif sebagai highlighted item di sidebar.

---

### Requirement 2: Dashboard Utama (index.html)

**User Story:** Sebagai pengguna, saya ingin melihat dashboard utama yang informatif, sehingga saya dapat mengetahui status sistem dan mengakses semua modul dengan cepat.

#### Acceptance Criteria

1. THE Dashboard SHALL menampilkan halaman `frontend/index.html` sebagai entry point utama aplikasi.
2. WHEN halaman Dashboard dimuat, THE Dashboard SHALL melakukan health check ke `GET /` dan menampilkan status server (aktif/tidak aktif) beserta indikator visual.
3. THE Dashboard SHALL menampilkan grid kartu navigasi untuk semua 20 modul, masing-masing berisi ikon, nama modul, deskripsi singkat, dan link ke halaman modul.
4. WHEN pengguna belum login (JWT_Token tidak ada di `localStorage`), THE Dashboard SHALL menampilkan tombol "Login" yang menonjol di area header.
5. WHEN pengguna sudah login, THE Dashboard SHALL menampilkan nama dan NPM pengguna di header beserta tombol "Logout".
6. THE Dashboard SHALL menampilkan statistik ringkasan (jumlah mahasiswa, dosen, mata kuliah) yang diambil dari API masing-masing modul secara asinkron.
7. IF request statistik ke salah satu endpoint gagal, THEN THE Dashboard SHALL menampilkan tanda "-" pada statistik tersebut tanpa menghentikan render halaman lainnya.

---

### Requirement 3: Sistem Autentikasi JWT (Frontend)

**User Story:** Sebagai mahasiswa, saya ingin bisa login menggunakan nomor HP dan mendapatkan akses ke fitur yang terproteksi, sehingga data akademik saya aman.

#### Acceptance Criteria

1. THE Auth_System SHALL menyediakan halaman `frontend/mahasiswa/login.html` sebagai satu-satunya entry point login.
2. WHEN pengguna mengirimkan nomor HP yang valid ke `POST /auth/login`, THE Auth_System SHALL menyimpan JWT_Token dan data profil pengguna ke `localStorage` dengan key `siakad_token` dan `siakad_user`.
3. WHEN pengguna berhasil login, THE Auth_System SHALL mengarahkan pengguna ke halaman Dashboard (`/`).
4. THE Guard SHALL memeriksa keberadaan `siakad_token` di `localStorage` setiap kali halaman terproteksi dimuat.
5. WHEN Guard memeriksa token dan token tidak ditemukan atau sudah kedaluwarsa, THE Guard SHALL mengarahkan pengguna ke halaman login dengan menyertakan parameter `redirect` berisi URL halaman asal.
6. WHEN pengguna mengklik tombol "Logout", THE Auth_System SHALL menghapus `siakad_token` dan `siakad_user` dari `localStorage` dan mengarahkan pengguna ke halaman login.
7. THE Auth_System SHALL menyediakan fungsi `getUser()` di `api.js` yang mengembalikan objek profil pengguna dari `localStorage` atau `null` jika tidak ada.
8. WHEN JWT_Token yang dikirim ke backend sudah kedaluwarsa dan backend mengembalikan HTTP 401, THE API_Client SHALL secara otomatis menghapus token dari `localStorage` dan mengarahkan pengguna ke halaman login.

---

### Requirement 4: Modul Mahasiswa — Halaman Daftar

**User Story:** Sebagai admin, saya ingin melihat daftar semua mahasiswa dalam tabel yang lengkap dan bisa melakukan aksi CRUD, sehingga pengelolaan data mahasiswa menjadi efisien.

#### Acceptance Criteria

1. WHEN halaman `frontend/mahasiswa/daftar.html` dimuat, THE Dashboard SHALL mengambil data dari `GET /mahasiswa` dan menampilkannya dalam tabel dengan kolom: NPM, Nama, Email, No. HP, Prodi, Angkatan, dan Aksi.
2. THE Dashboard SHALL menampilkan indikator loading saat data sedang diambil dari API.
3. IF `GET /mahasiswa` mengembalikan array kosong, THEN THE Dashboard SHALL menampilkan pesan "Belum ada data mahasiswa" di area tabel.
4. IF `GET /mahasiswa` mengembalikan error, THEN THE Dashboard SHALL menampilkan pesan error yang deskriptif tanpa crash.
5. WHEN admin mengklik tombol "Hapus" pada baris mahasiswa, THE Dashboard SHALL menampilkan dialog konfirmasi sebelum mengirim `DELETE /mahasiswa/:npm`.
6. WHEN `DELETE /mahasiswa/:npm` berhasil, THE Dashboard SHALL menampilkan toast sukses dan memuat ulang data tabel tanpa full page reload.
7. THE Dashboard SHALL menyediakan fitur pencarian client-side yang memfilter baris tabel berdasarkan NPM atau Nama secara real-time.
8. THE Dashboard SHALL menyediakan tombol "Edit" pada setiap baris yang membuka modal form edit inline, mengirim `PUT /mahasiswa/:npm` saat disimpan.
9. WHEN `PUT /mahasiswa/:npm` berhasil, THE Dashboard SHALL menutup modal dan memperbarui baris tabel yang bersangkutan.

---

### Requirement 5: Modul Mahasiswa — Halaman Tambah

**User Story:** Sebagai admin, saya ingin menambahkan data mahasiswa baru melalui form yang tervalidasi, sehingga data yang masuk ke sistem selalu valid.

#### Acceptance Criteria

1. THE Dashboard SHALL menyediakan form di `frontend/mahasiswa/tambah.html` dengan field: NPM (wajib), Nama (wajib), Email, No. HP (wajib), Program Studi, dan Angkatan.
2. WHEN pengguna mengklik "Simpan" dengan field wajib yang kosong, THE Dashboard SHALL menampilkan pesan validasi inline di bawah field yang kosong tanpa mengirim request ke API.
3. WHEN pengguna mengklik "Simpan" dengan semua field wajib terisi, THE Dashboard SHALL mengirim `POST /mahasiswa` dengan payload JSON yang sesuai.
4. WHEN `POST /mahasiswa` berhasil (HTTP 200), THE Dashboard SHALL menampilkan toast sukses dan mereset form ke kondisi kosong.
5. IF `POST /mahasiswa` mengembalikan HTTP 409 (NPM sudah terdaftar), THEN THE Dashboard SHALL menampilkan pesan error "NPM sudah terdaftar" di bawah field NPM.
6. IF `POST /mahasiswa` mengembalikan error lainnya, THEN THE Dashboard SHALL menampilkan toast error dengan pesan dari API.
7. THE Dashboard SHALL memvalidasi format email menggunakan regex standar sebelum mengirim request.
8. THE Dashboard SHALL memvalidasi bahwa field Angkatan berisi angka 4 digit antara 2000 dan 2099 jika diisi.

---

### Requirement 6: Modul Mahasiswa — Halaman Profil & Detail

**User Story:** Sebagai mahasiswa yang sudah login, saya ingin melihat profil saya sendiri, sehingga saya dapat memverifikasi data yang tersimpan di sistem.

#### Acceptance Criteria

1. THE Dashboard SHALL menyediakan halaman `frontend/mahasiswa/profil.html` yang hanya dapat diakses oleh pengguna yang sudah login.
2. WHEN halaman profil dimuat, THE Dashboard SHALL mengambil data dari `GET /auth/profile/:phone` menggunakan nomor HP dari `siakad_user` di `localStorage`.
3. THE Dashboard SHALL menampilkan data profil mahasiswa: NPM, Nama, Email, No. HP, Program Studi, dan Angkatan dalam layout kartu yang rapi.
4. WHEN Guard memeriksa token pada halaman profil dan token tidak ada, THE Guard SHALL mengarahkan pengguna ke halaman login dengan parameter `redirect=/mahasiswa/profil.html`.

---

### Requirement 7: Arsitektur Frontend yang Scalable untuk 20 Modul

**User Story:** Sebagai mahasiswa pengembang modul lain, saya ingin ada panduan dan template yang jelas, sehingga saya bisa membangun modul saya dengan konsisten tanpa harus memahami seluruh codebase.

#### Acceptance Criteria

1. THE Shared_Layout SHALL menyediakan file `frontend/shared/template.html` sebagai template starter yang bisa di-copy oleh setiap mahasiswa untuk memulai halaman modulnya.
2. THE Shared_Layout SHALL mendefinisikan konvensi penamaan folder modul sebagai `frontend/[nama-modul]/` (contoh: `frontend/dosen/`, `frontend/matkul/`).
3. THE API_Client SHALL mengekspor konstanta `BASE_URL` yang nilainya `https://dituniverse.alwaysdata.net` sehingga semua modul menggunakan satu sumber kebenaran untuk URL API.
4. THE Shared_Layout SHALL menyediakan fungsi `formatDate(isoString)` dan `formatCurrency(number)` di file utility `frontend/shared/utils.js` yang dapat digunakan semua modul.
5. WHEN mahasiswa baru menambahkan modul, THE Shared_Layout SHALL secara otomatis menampilkan modul tersebut di sidebar jika entry modul ditambahkan ke array konfigurasi di `layout.js`.
6. THE Shared_Layout SHALL menyediakan komponen tabel generik dengan fitur sorting kolom dan pagination client-side yang dapat digunakan oleh semua modul.

---

### Requirement 8: Keamanan dan Proteksi Endpoint Backend

**User Story:** Sebagai maintainer, saya ingin endpoint-endpoint sensitif dilindungi oleh JWT middleware, sehingga data tidak bisa diakses atau dimodifikasi tanpa autentikasi.

#### Acceptance Criteria

1. THE Auth_System SHALL menyediakan middleware Go Fiber `config.JWTMiddleware()` yang memvalidasi JWT_Token dari header `Authorization: Bearer <token>`.
2. WHEN request masuk ke endpoint terproteksi tanpa header Authorization, THE Auth_System SHALL mengembalikan HTTP 401 dengan pesan "Token tidak ditemukan".
3. WHEN request masuk ke endpoint terproteksi dengan token yang tidak valid atau kedaluwarsa, THE Auth_System SHALL mengembalikan HTTP 401 dengan pesan "Token tidak valid atau kedaluwarsa".
4. THE Auth_System SHALL mengekspor fungsi `config.GetClaimsFromContext(c)` yang memungkinkan controller mengambil data NPM dan Nama dari JWT claims tanpa parsing ulang.
5. WHILE endpoint `POST /mahasiswa`, `PUT /mahasiswa/:npm`, dan `DELETE /mahasiswa/:npm` aktif, THE Auth_System SHALL memproteksi endpoint tersebut dengan `JWTMiddleware`.
6. THE Auth_System SHALL membiarkan endpoint `GET /mahasiswa`, `GET /mahasiswa/:npm`, `POST /auth/login`, dan `GET /auth/profile/:phone` dapat diakses tanpa token (public).

---

### Requirement 9: Responsivitas dan Aksesibilitas UI

**User Story:** Sebagai pengguna, saya ingin tampilan SIAKAD dapat diakses dengan baik di berbagai ukuran layar, sehingga saya bisa menggunakannya dari laptop maupun smartphone.

#### Acceptance Criteria

1. THE Shared_Layout SHALL mengimplementasikan layout responsif menggunakan CSS Grid dan Flexbox yang bekerja pada lebar layar minimal 320px hingga 1920px.
2. WHEN lebar layar kurang dari 768px, THE Shared_Layout SHALL menyembunyikan sidebar dan menampilkan tombol hamburger menu untuk membuka/menutup sidebar sebagai overlay.
3. THE Shared_Layout SHALL memastikan semua elemen interaktif (tombol, link, input) memiliki ukuran touch target minimal 44x44px sesuai standar aksesibilitas.
4. THE Shared_Layout SHALL menggunakan atribut ARIA (`aria-label`, `role`, `aria-live`) pada komponen dinamis seperti alert, modal, dan loading indicator.
5. THE Shared_Layout SHALL memastikan rasio kontras warna teks terhadap background memenuhi standar WCAG 2.1 level AA (minimal 4.5:1 untuk teks normal).
6. WHEN tabel data memiliki lebih dari 5 kolom, THE Shared_Layout SHALL membuat tabel dapat di-scroll secara horizontal pada layar kecil tanpa merusak layout halaman.

---

### Requirement 10: Penanganan Error dan State Loading

**User Story:** Sebagai pengguna, saya ingin mendapatkan feedback yang jelas saat terjadi error atau saat data sedang dimuat, sehingga saya tidak bingung dengan kondisi aplikasi.

#### Acceptance Criteria

1. WHEN request API sedang berlangsung, THE API_Client SHALL menampilkan indikator loading (skeleton screen atau spinner) pada area konten yang menunggu data.
2. IF request API gagal karena koneksi jaringan terputus, THEN THE API_Client SHALL menampilkan pesan "Tidak dapat terhubung ke server. Periksa koneksi internet Anda." beserta tombol "Coba Lagi".
3. IF server mengembalikan HTTP 500, THEN THE API_Client SHALL menampilkan pesan "Terjadi kesalahan pada server. Silakan coba beberapa saat lagi." tanpa menampilkan detail teknis ke pengguna.
4. THE API_Client SHALL mencatat detail error (URL, status code, pesan) ke `console.error` untuk keperluan debugging tanpa mengekspos informasi sensitif ke UI.
5. WHEN operasi CRUD berhasil, THE API_Client SHALL menampilkan toast notifikasi sukses yang otomatis menghilang setelah 4 detik.
6. WHEN operasi CRUD gagal, THE API_Client SHALL menampilkan toast notifikasi error yang tetap tampil hingga pengguna menutupnya secara manual.

