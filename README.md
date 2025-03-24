# GoTasker

Production-grade Go project scaffolded for clean architecture and scalability.

## Folder Structure
```
cmd/                → Entry point (main.go)
internal/
  handler/          → HTTP handlers
  service/          → Business logic
  repository/       → DB abstraction
  model/            → Structs/models
  config/           → App config
  middleware/       → Middleware (auth, logging)
cli/                → CLI tool
pkg/                → Reusable utils
test/               → Tests
```

## Running Locally

```bash
docker-compose up --build
```
