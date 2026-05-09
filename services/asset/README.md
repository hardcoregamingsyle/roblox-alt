# Asset Storage Service

The `Asset Storage Service` provides a secure, scalable interface for user-generated content (UGC). It bridges the Godot Studio client with cloud storage via S3.

## Security Architecture
- **Presigned URLs:** Assets are uploaded directly to S3. The backend never touches raw binary data, preventing memory exhaustion and reducing the platform's attack surface.
- **Expiration:** Presigned URLs are strictly limited to a 3-minute TTL.
- **Validation:** Only `model` and `texture` MIME types are permitted.

## API Endpoints
### `POST /assets/upload`
- **Description:** Requests a secure URL for binary upload.
- **Request:** `{ "fileType": "model" | "texture" }`
- **Response:** `{ "assetId": "uuid", "url": "string" }`