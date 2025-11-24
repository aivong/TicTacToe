# API Contracts

This directory is intentionally empty.

**Reason**: The CLI Tic-Tac-Toe game is a standalone terminal application with no external APIs, network interfaces, or inter-process communication. All interaction happens through stdin/stdout within a single process.

**If this changes in the future**, this directory would contain:
- OpenAPI/Swagger specifications (for REST APIs)
- GraphQL schemas (for GraphQL APIs)
- Protocol buffer definitions (for gRPC)
- JSON Schema files (for data validation)

**Current Architecture**: Single-process, in-memory game state with terminal UI.
