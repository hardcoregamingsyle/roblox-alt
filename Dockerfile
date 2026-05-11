FROM golang:1.23.6-alpine AS builder
RUN addgroup -g 10001 nexus && adduser -u 10001 -G nexus -D nexus
WORKDIR /app

# Copy dependency manifests for all relevant packages
COPY go.work go.sum ./
COPY services/api/go.mod services/api/go.sum ./services/api/
COPY packages/auth/go.mod packages/auth/go.sum ./packages/auth/

RUN go mod download

# Copy source
COPY services/api/ ./services/api/
COPY packages/auth/ ./packages/auth/

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /api ./services/api/

FROM gcr.io/distroless/static-debian12:latest
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /api /api
USER 10001
EXPOSE 3000
ENTRYPOINT ["/api"]