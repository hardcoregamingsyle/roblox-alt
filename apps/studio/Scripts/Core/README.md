# Core Sandbox Infrastructure

This directory contains the security-critical components for the NexusStudio environment.

## Namespace: `NexusStudio.Core`
- **`LuauVM`**: Manages the native `lua_State`. Enforces 64MB memory limits.
- **`SandboxValidator`**: Performs static analysis on Luau scripts to block unauthorized namespaces (`os`, `io`, `debug`).
- **`LuauScriptManager`**: Implements an object pool for `LuauVM` to optimize performance and reduce GC pressure.

## Security Constraints
- **Memory**: 64MB per script context.
- **API Surface**: Restricted to `ILuauSandboxAPI`.
- **Static Analysis**: Regex-based blocking of dangerous globals.