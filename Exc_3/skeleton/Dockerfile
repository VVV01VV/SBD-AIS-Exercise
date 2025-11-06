# ---------- builder: compile the Go app ----------
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
# fix CRLF if needed (Windows)
RUN sed -i 's/\r$//' scripts/build-application.sh
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh

# ---------- runtime: small image with the binary ----------
FROM alpine:3.20 AS run
WORKDIR /app

# (optional but recommended) run as non-root
RUN adduser -D -u 10001 appuser
USER appuser

# compiled binary
COPY --from=builder /app/ordersystem /app/ordersystem

# static assets your binary may serve (Swagger, frontend)
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/frontend /app/frontend

# your app listens on 3000
EXPOSE 3000

ENTRYPOINT ["/app/ordersystem"]
