#!/usr/bin/env bash
set -euo pipefail

# Config
NET_NAME="sbdnet"
VOL_NAME="pg18data"
DB_NAME="pg18"
APP_IMG="orderservice"
APP_CONT="orderservice"
HOST_DB_PORT=5435
APP_PORT=3000

ENV_FILE="./debug.env"                        # must exist next to this script
PGDATA_PATH="/var/lib/postgresql/18/docker"   # persisted inside container

ensure_prereqs() {
  command -v docker >/dev/null || { echo "Docker not found"; exit 1; }
  [[ -f "$ENV_FILE" ]] || { echo "Missing $ENV_FILE"; exit 1; }
}

net_up() {
  docker network create "$NET_NAME" >/dev/null 2>&1 || true
}

vol_up() {
  docker volume create "$VOL_NAME" >/dev/null 2>&1 || true
}

db_up() {
  echo "→ Starting PostgreSQL ($DB_NAME) on host port ${HOST_DB_PORT}…"
  docker rm -f "$DB_NAME" >/dev/null 2>&1 || true

  docker run -d --name "$DB_NAME" \
    --env-file "$ENV_FILE" \
    --network "$NET_NAME" \
    -e PGDATA="$PGDATA_PATH" \
    -p ${HOST_DB_PORT}:5432 \
    -v ${VOL_NAME}:${PGDATA_PATH} \
    postgres:18

  echo "✓ DB container started → localhost:${HOST_DB_PORT}"
}

app_build() {
  echo "→ Building ${APP_IMG} (multi-stage Dockerfile)…"
  docker build -t "$APP_IMG" .
  echo "✓ Image built"
}

app_up() {
  echo "→ Starting ${APP_CONT} on port ${APP_PORT}…"
  docker rm -f "$APP_CONT" >/dev/null 2>&1 || true

  docker run --name "$APP_CONT" \
    --env-file "$ENV_FILE" \
    --network "$NET_NAME" \
    -p ${APP_PORT}:${APP_PORT} \
    "$APP_IMG"

  # note: foreground run (Ctrl+C to stop). For detached, add -d above.
}

logs() {
  echo "=== Logs: $DB_NAME ==="
  docker logs --tail 80 "$DB_NAME" || true
  echo
  echo "=== Logs: $APP_CONT ==="
  docker logs --tail 80 "$APP_CONT" || true
}

down() {
  echo "→ Stopping containers…"
  docker rm -f "$APP_CONT" >/dev/null 2>&1 || true
  docker rm -f "$DB_NAME"   >/dev/null 2>&1 || true
  echo "✓ Containers removed"
  # Uncomment to also remove volume/network:
  # docker volume rm "$VOL_NAME" || true
  # docker network rm "$NET_NAME" || true
}

usage() {
  cat <<EOF
Usage: ./run.sh <command>

Commands:
  up        Create net/volume, start DB, build image, run service (foreground)
  rebuild   Rebuild image and (re)run service (DB left as-is)
  logs      Show last 80 lines of DB and service logs
  down      Stop & remove service and DB containers

Examples:
  ./run.sh up
  ./run.sh rebuild
  ./run.sh logs
  ./run.sh down
EOF
}

case "${1:-}" in
  up)
    ensure_prereqs
    net_up
    vol_up
    db_up
    app_build
    app_up
    ;;
  rebuild)
    ensure_prereqs
    app_build
    app_up
    ;;
  logs)
    logs
    ;;
  down)
    down
    ;;
  *)
    usage
    ;;
esac
