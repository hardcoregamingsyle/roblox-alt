FROM golang:1.23.6-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /nexus-api ./services/api

FROM alpine:3.21.2
# Pinning to specific hash for supply chain security
RUN apk add --no-cache libstdc++ && rm -rf /var/cache/apk/*
COPY --from=builder /nexus-api /usr/local/bin/nexus-api
USER 10001
CMD ["/usr/local/bin/nexus-api"]