# API Documentation
This document outlines the core interface for the NexusEngine backend.

## Authentication (RS256)
- `POST /auth/login`: Authenticates user; returns a signed JWT.
- `POST /auth/register`: Creates an account; enforces alphanumeric constraints.

## Asset Broker
- `POST /assets/upload`: Requests a 3-minute TTL presigned URL.
- **Constraints**: 1KB metadata limit; MIME type must be 'model' or 'texture'.

## Game Discovery
- `GET /discovery/trending`: Paginated list of public game instances.
- **Rate Limiting**: IP-based sliding window (100 req/min).

## Security Headers
- `X-Correlation-ID`: Required for all requests to enable distributed tracing.
- `Strict-Transport-Security`: Enforced for all endpoints.