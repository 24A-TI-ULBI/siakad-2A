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
   JWT_SECRET=<jwt-secret-kuat>
   WAQRKEYWORD=<keyword>
   WEBHOOKURL=https://<appname>.alwaysdata.net/whatsauth/webhook
   WEBHOOKSECRET=<secret>
   WAPHONENUMBER=62811xxxxxx
   ```
4. Set **Command** di bagian Site configuration:
   ```
   /home/<username>/gocroot
   ```
5. **Remote Access > SSH > Modify** → aktifkan password-based login
6. **Customer Area > Profile > Managing Tokens** → generate API key

## Status Checklist

- [ ] `sshhost` diisi
- [ ] `sshusername` diisi
- [ ] `sshpassword` diisi
- [ ] `sshport` diisi
- [ ] `apikey` diisi
- [ ] `appid` diisi
- [ ] `folder` diisi
- [ ] Environment variable di alwaysdata sudah diset
