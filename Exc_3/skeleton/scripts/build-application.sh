#!/usr/bin/env bash
set -euo pipefail

# build the single main package at repo root
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/ordersystem .

