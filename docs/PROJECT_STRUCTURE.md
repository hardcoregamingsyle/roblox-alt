# Project Structure: NexusEngine

This repository uses a monorepo structure to manage the engine, services, and shared libraries.

- `/apps`: End-user applications.
  - `/studio`: Godot 4 editor application.
  - `/client`: Godot 4 game runtime.
- `/services`: Microservices.
  - `/api`: API Gateway (Go).
  - `/matchmaker`: Session management (Go).
- `/packages`: Shared code.
  - `/auth`: Security, middleware, and sanitization utilities.
- `/infra`: Kubernetes manifests and CI/CD pipelines.
- `/docs`: Architectural Decision Records (ADRs).