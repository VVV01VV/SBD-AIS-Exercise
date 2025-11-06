# ---------- builder: compile the Go app ----------
FROM golang:1.25 AS builder
WORKDIR /app

# 1) Better build cache: pull modules first
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy the rest of the sources
COPY . .

# 3) Fix CRLF if coming from Windows and build the static binary
RUN sed -i 's/\r$//' scripts/build-application.sh \
 && chmod +x scripts/build-application.sh \
 && ./scripts/build-application.sh
# -> produces /app/ordersystem

# ---------- runtime: minimal image with the binary ----------
FROM alpine:3.20 AS run
WORKDIR /app

# CA certs for HTTPS calls
RUN apk add --no-cache ca-certificates

# Run as non-root
RUN adduser -D -u 10001 appuser
USER appuser

# Copy compiled binary
COPY --from=builder /app/ordersystem /app/ordersystem

# Copy static assets served by the binary (Swagger UI, frontend)
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/frontend /app/frontend

# App listens on 3000 inside the container
EXPOSE 3000

ENTRYPOINT ["/app/ordersystem"]

