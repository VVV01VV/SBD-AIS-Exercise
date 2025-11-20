FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
# Normalize Windows line endings and run the build script
RUN sed -i 's/\r$//' scripts/build-application.sh \
    && chmod +x scripts/build-application.sh \
    && ./scripts/build-application.sh

FROM alpine AS run
WORKDIR /
COPY --from=builder /app/ordersystem /app/ordersystem
# EXPOSE doesn't actually do anything!
EXPOSE 3000
CMD ["/app/ordersystem"]