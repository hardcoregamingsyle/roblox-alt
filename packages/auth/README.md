# Auth Package

Core security utilities for NexusEngine.

## Functions
- `SanitizeRequestID(id string) string`: Sanitizes input using NFKC normalization and regex validation.
- `GetLimiter(ip string)`: Returns a rate limiter per IP, sharded to prevent global lock contention.