# API Contract & Integration Guide

## Request Headers
- `X-Correlation-ID`: **Mandatory**. Every request must include a UUID v4. Used for distributed tracing.

## Authentication
- **Scheme**: Bearer Token (RS256 JWT).
- **Validation**: Services verify tokens using the public key retrieved from the internal Vault or cached key store.

## Rate Limiting
- **Policy**: Sharded, IP-based sliding window (100 req/min).

## Security Constraints
- **Max Payload**: 64KB (JSON), 1MB (Asset Metadata).
- **TLS**: Enforced `verify-full` for all database and service-to-service communication.
- **Input Sanitization**: All strings are normalized via `NFKC` and filtered for control characters.
- **Zero-Allocation Paths**: All high-frequency endpoints utilize `sync.Pool` (Go) or `ArrayPool<byte>` (C#) to mitigate GC jitter.