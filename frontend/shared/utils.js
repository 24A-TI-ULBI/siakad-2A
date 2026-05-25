/**
 * SIAKAD Shared Utilities
 * Dipakai oleh semua modul frontend.
 * Sertakan file ini di setiap halaman:
 *   <script src="../shared/utils.js"></script>
 */

// ===================== API BASE URL =====================
const API_BASE = (
  window.location.hostname === 'localhost' ||
  window.location.hostname === '127.0.0.1'
) ? 'http://localhost:8080' : 'https://dituniverse.alwaysdata.net';

// ===================== AUTH HELPERS =====================
const Auth = {
  getToken() {
    return localStorage.getItem('siakad_token') || null;
  },
  setToken(token) {
    localStorage.setItem('siakad_token', token);
  },
  removeToken() {
    localStorage.removeItem('siakad_token');
  },
  getUser() {
    try {
      const raw = localStorage.getItem('siakad_user');
      return raw ? JSON.parse(raw) : null;
    } catch {
      return null;
    }
  },
  setUser(user) {
    localStorage.setItem('siakad_user', JSON.stringify(user));
  },
  removeUser() {
    localStorage.removeItem('siakad_user');
  },
  isLoggedIn() {
    return !!this.getToken();
  },
  logout() {
    this.removeToken();
    this.removeUser();
  }
};

// ===================== API FETCH =====================
/**
 * Wrapper fetch dengan base URL dan error handling otomatis.
 * @param {string} path  - path endpoint, contoh: '/mahasiswa'
 * @param {object} opts  - fetch options (method, body, headers, dll)
 * @returns {Promise<any>} - data dari response JSON
 */
async function apiFetch(path, opts = {}) {
  const headers = { 'Content-Type': 'application/json', ...opts.headers };
  const token = Auth.getToken();
  if (token) headers['Authorization'] = 'Bearer ' + token;

  const res = await fetch(API_BASE + path, { ...opts, headers });
  const json = await res.json();

  if (json.status !== 'success') {
    throw new Error(json.message || 'Terjadi kesalahan');
  }
  return json.data;
}

// ===================== TOAST =====================
(function initToast() {
  const container = document.createElement('div');
  container.className = 'toast-container';
  document.body.appendChild(container);
  window._toastContainer = container;
})();

/**
 * Tampilkan toast notification.
 * @param {string} message
 * @param {'success'|'error'|'info'} type
 * @param {number} duration - ms, default 3500
 */
function showToast(message, type = 'success', duration = 3500) {
  const toast = document.createElement('div');
  toast.className = `toast toast-${type}`;
  toast.textContent = message;
  window._toastContainer.appendChild(toast);

  setTimeout(() => {
    toast.style.animation = 'slideOut 0.2s ease forwards';
    setTimeout(() => toast.remove(), 200);
  }, duration);
}

// ===================== HEADER USER INFO =====================
/**
 * Isi info user di header jika sudah login.
 * Panggil setelah DOM ready.
 */
function initHeaderUser() {
  const user = Auth.getUser();
  const avatarEl = document.getElementById('header-avatar');
  const nameEl   = document.getElementById('header-username');
  if (!avatarEl) return;

  if (user && user.nama) {
    const initials = user.nama.split(' ').map(w => w[0]).join('').slice(0, 2).toUpperCase();
    avatarEl.textContent = initials;
    avatarEl.href = '../mahasiswa/profil.html';
    if (nameEl) nameEl.textContent = user.nama;
  } else {
    avatarEl.textContent = '?';
    avatarEl.href = '../mahasiswa/login.html';
    if (nameEl) nameEl.textContent = 'Login';
  }
}

// ===================== DYNAMIC NAV =====================
/**
 * Sembunyikan tab Login di nav jika user sudah login,
 * tampilkan kembali jika belum login.
 * Panggil setelah DOM ready.
 */
function initNav() {
  const loginLinks = document.querySelectorAll('.siakad-nav a[href="login.html"]');
  loginLinks.forEach(link => {
    link.style.display = Auth.isLoggedIn() ? 'none' : '';
  });
}
function formatAngkatan(year) {
  return year ? String(year) : '-';
}

function formatPhone(phone) {
  if (!phone) return '-';
  return phone.replace(/(\d{4})(\d{4})(\d+)/, '$1-$2-$3');
}

function getInitials(nama) {
  if (!nama) return '?';
  return nama.split(' ').map(w => w[0]).join('').slice(0, 2).toUpperCase();
}

// ===================== MODAL HELPERS =====================
function openModal(id) {
  const el = document.getElementById(id);
  if (el) el.classList.add('open');
}

function closeModal(id) {
  const el = document.getElementById(id);
  if (el) el.classList.remove('open');
}

// Tutup modal saat klik overlay
document.addEventListener('click', function(e) {
  if (e.target.classList.contains('modal-overlay')) {
    e.target.classList.remove('open');
  }
});

// Tutup modal saat tekan Escape
document.addEventListener('keydown', function(e) {
  if (e.key === 'Escape') {
    document.querySelectorAll('.modal-overlay.open').forEach(el => {
      el.classList.remove('open');
    });
  }
});
