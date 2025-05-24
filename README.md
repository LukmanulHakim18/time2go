<p align="center">
  <img src="t2g-logo.jpeg" width="200"/>
</p>

<h1 align="center">Time2Go</h1>
<p align="center">⏱️ Time-based event scheduler with HTTP trigger and retry mechanism</p>

# Time2Go

**Time2Go** adalah sistem penjadwalan event ringan berbasis waktu (time-based scheduler) untuk menjalankan HTTP request secara otomatis. Cocok untuk kebutuhan trigger API berkala, webhook tertunda, atau sistem reminder otomatis. Time2Go mendukung retry policy, autentikasi dasar, dan penyimpanan sementara di Redis.

## 🚀 Fitur Utama

- ⏰ Penjadwalan HTTP request berdasarkan waktu (`RFC3339`)
- 🔁 Retry policy (Fixed dan Exponential Backoff)
- 🔐 Dukungan Basic Auth
- 💾 Penyimpanan event terdistribusi via Redis
- 📡 Eksekusi HTTP request otomatis dengan dukungan header, query param, body, dan timeout

---

## 📦 Struktur Konfigurasi Event (Json)

```json
{
  "client_name": "client-1",
  "event_name": "send-webhook",
  "event_id": "uuid-1234",
  "schedule_at": "2025-05-24T13:10:00+07:00",
  "status": "PENDING",
  "last_error": "",
  "request_config": {
    "method": "POST",
    "url": "https://webhook.site/your-endpoint",
    "headers": {
      "Content-Type": "application/json"
    },
    "query_params": {
      "source": "time2go"
    },
    "body": "eyJtZXNzYWdlIjogIlRpbWUgdG8gR08ifQ==",
    "timeout": "10s",
    "auth": {
      "username": "myuser",
      "password": "mypassword"
    }
  },
  "retry_policy": {
    "type": 1, // 1 : Fixed, 2 : Exponential
    "retry_count": 5,
    "max_attempts": 5,
    "attempt_count": 0
  }
}
```
