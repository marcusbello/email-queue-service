# 📧 Email Queue Microservice (Go)
[![FS3E1TB.md.jpg](https://iili.io/FS3E1TB.md.jpg)](https://freeimage.host/i/FS3E1TB)

## ✅ Features

- Accepts email jobs over HTTP (`POST /send-email`)
- Validates input
- Queues jobs in memory
- Processes jobs with concurrent workers
- Gracefully shuts down on Ctrl+C

## 🚀 Run

```bash
go run main.go
```

- ### Run with command flags
```bash
  -haddr string
    	Set the HTTP address (default "localhost:8080")
  -queue_size int
    	Queue size (default 10)
  -workers int
    	Number of worker goroutines (default 3)
```
- Run `go run main.go -haddr "localhost:11000" -workers 5 -queue_size 100`

## Testing

### Example Payload
```json
{
  "to": "user@example.com",
  "subject": "Welcome!",
  "body": "Thanks for signing up."
}
```
### Testing with curl
```bash
curl -X POST http://localhost:8080/send-email \
-H "Content-Type: application/json" \
-d '{"to":"user@example.com", "subject":"Welcome", "body":"Thanks!"}'
```


## Todo
- Add a Redis Queue
- Add Retry Logic
- Add Prometheus
- Add Testcases

## Tree
```
.
├── go.mod
├── internal
│   ├── email
│   │   └── email.go
│   ├── queue
│   │   ├── memory_queue.go
│   │   └── queue.go
│   ├── server
│   │   └── server.go
│   └── worker
│       └── worker.go
├── main.go
└── readme.md

5 directories, 8 files
```


