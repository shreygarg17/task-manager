#!/bin/bash

set -e

echo "📁 Creating folder structure..."

# Create folder structure
mkdir -p \
  cmd/server \
  internal/handler/http \
  internal/service \
  internal/repository/postgres \
  internal/model \
  internal/config \
  internal/middleware \
  cli/taskctl \
  pkg/utils \
  test

# Create main entry point files
touch cmd/server/main.go
touch cli/taskctl/main.go

# Create internal files
touch internal/handler/http/task_handler.go
touch internal/service/task_service.go
touch internal/repository/postgres/task_repo.go
touch internal/repository/interface.go
touch internal/model/task.go
touch internal/config/config.go
touch internal/middleware/logger.go

# Create utility and test files
touch pkg/utils/validators.go
touch test/task_test.go

# Create .gitignore
cat <<EOF > .gitignore
/bin
*.log
*.exe
*.out
*.test
.env
EOF

# Create Dockerfile
cat <<EOF > Dockerfile
# Simple Go Dockerfile
FROM golang:1.21

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o server ./cmd/server

CMD ["./server"]
EOF

# Create docker-compose.yml
cat <<EOF > docker-compose.yml
version: '3.9'

services:
  db:
    image: postgres:15
    restart: always
    environment:
      POSTGRES_USER: tasker
      POSTGRES_PASSWORD: password
      POSTGRES_DB: tasker_db
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=tasker
      - DB_PASSWORD=password
      - DB_NAME=tasker_db

volumes:
  pgdata:
EOF

# Create empty .env file
touch .env

# Create README.md
cat <<EOF > README.md
# GoTasker

Production-grade Go project scaffolded for clean architecture and scalability.

## Folder Structure
\`\`\`
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
\`\`\`

## Running Locally

\`\`\`bash
docker-compose up --build
\`\`\`
EOF

echo "✅ Scaffold complete!"
