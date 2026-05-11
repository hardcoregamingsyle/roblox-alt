/*
Package auth provides security primitives for the NexusEngine platform.

Includes:
- Argon2id password hashing with constant-time verification.
- RS256 JWT signing and validation logic.
- IP-based rate limiting via sharded Redis/TTL caches.
- NFKC-based Unicode input normalization for usernames.
*/
package auth