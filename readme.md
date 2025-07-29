# 📧 Email Queue Microservice (Go)

## ✅ Features

- Accepts email jobs over HTTP (`POST /send-email`)
- Validates input
- Queues jobs in memory
- Processes with concurrent workers
- Gracefully shuts down on Ctrl+C

## 🚀 Run

```bash
go run main.go
```


## Todo
- Add a Redis Queue
- Add Retry Logic
- Add Prometheus



