#!/bin/sh
set -e
cd /app
go mod download
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem