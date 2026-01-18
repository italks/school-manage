#!/bin/bash

# Exit on error
set -e

echo "==========================================="
echo "  School Management System - Build Script  "
echo "==========================================="

# Build Backend
echo "[1/3] Building Backend (Go)..."
cd api
if [ -f "main" ]; then
    rm main
fi
# Check if Go is installed
if command -v go &> /dev/null; then
    CGO_ENABLED=0 GOOS=linux go build -o main .
    echo "Backend binary built successfully."
else
    echo "Warning: Go not found. Skipping local binary build (Docker will handle it)."
fi
cd ..

# Build Frontend
echo "[2/3] Building Frontend (Vue)..."
cd web
if command -v pnpm &> /dev/null; then
    pnpm install
    pnpm build
    echo "Frontend built successfully."
else
    echo "Warning: pnpm not found. Skipping local frontend build (Docker will handle it)."
fi
cd ..

# Build Docker Images
echo "[3/3] Building Docker Images..."
docker-compose build

echo "==========================================="
echo "  Build Complete!                          "
echo "  Run 'docker-compose up -d' to start.     "
echo "==========================================="
