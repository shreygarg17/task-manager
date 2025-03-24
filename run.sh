#!/bin/bash

set -e

# Colors for output
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Starting Go Tasker Backend Setup...${NC}"

# Set environment variables (adjust if needed)
export PORT=8080
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=tasker
export DB_PASSWORD=password
export DB_NAME=taskerdb

echo -e "${GREEN}✅ Environment variables set.${NC}"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
  echo -e "${GREEN}⚠️  Docker is not running. Please start Docker and rerun this script.${NC}"
  exit 1
fi

# Start Postgres container (if not already running)
if [ "$(docker ps -q -f name=go-tasker-postgres)" ]; then
  echo -e "${GREEN}✅ PostgreSQL container already running.${NC}"
else
  echo -e "${GREEN}📦 Starting PostgreSQL container...${NC}"
  docker run --name go-tasker-postgres -p 5432:5432 \
  -e POSTGRES_USER=$DB_USER \
  -e POSTGRES_PASSWORD=$DB_PASSWORD \
  -e POSTGRES_DB=$DB_NAME \
  -d postgres
fi

# Wait a few seconds for DB to be ready
echo -e "${GREEN}⏳ Waiting for PostgreSQL to be ready...${NC}"
sleep 5

# Create the tasks table if it doesn't exist
echo -e "${GREEN}🛠  Creating tasks table if needed...${NC}"
docker exec -i go-tasker-postgres psql -U $DB_USER -d $DB_NAME <<EOF
CREATE TABLE IF NOT EXISTS tasks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
EOF

echo -e "${GREEN}✅ Tasks table ready.${NC}"

# Run Go server
echo -e "${GREEN}🚀 Running Go server...${NC}"
go run cmd/server/main.go
