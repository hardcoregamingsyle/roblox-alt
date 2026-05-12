# NexusEngine Platform

NexusEngine is a production-hardened, high-performance game engine ecosystem built on Godot 4.3 (C#/.NET 8) and Go 1.23.

## 🏗️ Architecture
- **API Gateway (Go)**: High-concurrency gateway with JWT (RS256) auth and zero-trust networking.
- **Asset Broker (S3)**: Secure, broker-based asset ingestion with 3-minute TTL presigned URLs.
- **Sandbox (Wasm/Luau)**: Isolated execution environment using Wasmtime with strict 100k instruction budgets and 64MB memory caps.
- **IDE (Godot/C#)**: Secure, dockable studio environment for UGC.

## 🎓 Developer Onboarding
If you are coming from Unity or Unreal, see `docs/ONBOARDING_GUIDE.md` for a mapping of engine paradigms.

## 🚀 Setup
1. **Dependencies**: Requires Go 1.23, .NET 8, and Node 22.
2. **Install**: `pnpm install --frozen-lockfile`
3. **Environment**: Copy `.env.example` to `.env` and configure keys.
4. **Testing**: `pnpm test` (Runs Vitest and Go test suites).

## 🛡️ Security
- **Binary Contract**: All state exchange utilizes `LayoutKind.Explicit` for zero-copy binary alignment.
- **Hardening**: Egress is restricted via Kubernetes `NetworkPolicies`.
- **Error Taxonomy**: All sandbox failures (Traps, OOM, Budgeting) are reported via a standardized JSON schema.

## Deployment
- **Agones**: Game servers are orchestrated via `infra/agones-gs.yaml`.
- **Docker**: Build via `docker build -t nexus-api .` (Exposes 3000).

## API Documentation
See `docs/API_REFERENCE.md` for endpoints and `docs/API_CONTRACT.md` for binary layout definitions.