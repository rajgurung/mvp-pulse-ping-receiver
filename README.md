```markdown
# MVP Pulse Ping Receiver

[![Go Version](https://img.shields.io/badge/go-1.22-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A lightweight Go server that receives heartbeat pings from multiple clients to monitor the health and uptime of their systems.

## 🚀 Project Purpose

The Ping Receiver accepts incoming `POST /ping/{token}` requests from various clients to:
- Confirm client liveness
- Optionally forward pings to the core API (Rails backend) for tracking
- Respond quickly with 200 OK

## 📦 Project Structure

```plaintext
mvp-pulse-ping-receiver/
├── main.go           # Entry point of the server
├── router/
│   └── router.go     # Route definitions
├── internal/
│   └── handlers/
│       └── ping.go   # Ping handler logic
├── go.mod            # Go module file
└── README.md         # Project documentation
```

## 🛠 Getting Started

### Prerequisites

- Go 1.22+ installed
- Git installed

### Clone the repository

```bash
git clone git@github.com:rajgurung/mvp-pulse-ping-receiver.git
cd mvp-pulse-ping-receiver
```

### Install dependencies

```bash
go mod tidy
```

### Run the server

```bash
go run .
```

The server will start on [http://localhost:8080](http://localhost:8080)

## 📜 API Endpoint

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | `/ping/{token}`    | Client sends a heartbeat ping |

### Example Request

```bash
curl -X POST http://localhost:8080/ping/your-client-token
```

✅ The server will respond with `200 OK`.

## 🧩 Future Improvements

- Forward pings to Rails Core API
- Add metrics (e.g., Prometheus integration)
- Add authentication and token validation
- Graceful shutdown handling
- Healthcheck endpoints

## 🧠 Tech Stack

- Go
- Gorilla Mux (HTTP router)

## 📄 License

This project is licensed under the MIT License.

## ✨ Author

- [Raj Gurung](https://github.com/rajgurung)