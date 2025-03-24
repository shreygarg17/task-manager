# 🧠 Go Tasker – Task Management API in Golang

Go Tasker is a simple, clean, and modular task management REST API built with Golang.  
It uses PostgreSQL for persistence and follows a layered architecture with clearly separated concerns (Handler → Service → Repository).

---

## 🚀 Features
- RESTful CRUD APIs for tasks
- PostgreSQL database integration
- Lightweight routing with [Chi](https://github.com/go-chi/chi)
- Clean, scalable folder structure
- Environment-based configuration
- Testable architecture

---

## 📂 Project Structure

```
go-tasker/
├── cmd/
│   └── server/             → Application entrypoint
├── internal/
│   ├── config/             → App config & DB connection
│   ├── handler/
│   │   └── http/           → HTTP handlers and routes
│   ├── repository/
│   │   └── postgres/       → Database logic (CRUD)
│   ├── service/            → Business logic
│   └── model/              → Data models
├── .env                    → Environment variables file (not to be pushed)
├── .gitignore              → Git ignore rules
└── README.md
```

---

## 🛠 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/michaelorina/go-tasker.git
cd go-tasker
```

### 2. Setup environment variables

Create a `.env` file in the root directory:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=tasker
DB_PASSWORD=password
DB_NAME=taskerdb
```

### 3. Run the application

```bash
go run cmd/server/main.go
```

### 4. Run tests

```bash
go test ./...
```

---

## 📦 API Endpoints

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | /tasks           | Get all tasks           |
| POST   | /tasks           | Create a new task       |
| GET    | /tasks/{id}      | Get task by ID          |
| PUT    | /tasks/{id}      | Update a task by ID     |
| DELETE | /tasks/{id}      | Delete a task by ID     |

---
