# Asset Pipeline: Direct-to-S3 Workflow

This document details the secure asset upload pipeline for NexusEngine.

## Overview
To prevent DoS attacks and minimize server bandwidth costs, NexusEngine uses a broker-based architecture for user-uploaded assets.

## Workflow
1. **Client Request:** `NexusStudio` requests an upload URL via `POST /assets/upload`.
2. **Broker Validation:** The `Asset Storage Service` validates the metadata (type/size).
3. **Presigned URL:** The server returns a 3-minute TTL S3 URL.
4. **Direct Upload:** The client uploads the binary directly to S3.

## Security Constraints
- **MIME Whitelisting:** Only `model` and `texture` types are accepted.
- **TTL:** URLs expire after 3 minutes.
- **Payload Limits:** Metadata requests are strictly capped at 1KB to prevent DoS.