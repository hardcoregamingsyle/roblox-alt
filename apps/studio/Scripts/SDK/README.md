# NexusEngine SDK

The `NexusEngine SDK` is the official interface for developers to interact with the platform backend.

## Architecture
- **Memory Pooling:** All network operations utilize `ArrayPool<byte>` to maintain stable frame rates.
- **Serialization:** Uses `protobuf-net` for efficient, binary-compact state replication.
- **Security:** Capability-based access ensures game scripts cannot access host system APIs directly.

## Usage
```csharp
using NexusStudio.SDK;

// Send a game state update
var state = new PlayerState { Position = new Vector3(0, 10, 0) };
NexusEngineSDK.Instance.Network.Send(state);
```

## Security Limits
- **Packet Size:** 8KB Max.
- **Serialization:** Protobuf-only to prevent injection vulnerabilities.