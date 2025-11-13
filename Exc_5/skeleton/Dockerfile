# ---------- builder: compile the Go app ----------
FROM golang:1.25 AS builder
WORKDIR /app

# 1) Copy module files and download deps (better caching)
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy the rest of the source code
COPY . .

# 3) Normalize Windows line endings and run the build script
RUN sed -i 's/\r$//' scripts/build-application.sh \
 && chmod +x scripts/build-application.sh \
 && ./scripts/build-application.sh

# ---------- runtime: minimal image with the binary ----------
FROM alpine:3.20 AS run
WORKDIR /app

# Optional but nice: certs + non-root user
RUN apk add --no-cache ca-certificates \
 && adduser -D -u 10001 appuser
USER appuser

# compiled binary
COPY --from=builder /app/ordersystem /app/ordersystem

# (if your app serves Swagger/frontend from these folders)
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/frontend /app/frontend

EXPOSE 3000
ENTRYPOINT ["/app/ordersystem"]
