# NexusEngine Binary Contract
## Memory Layout (Pack=8)
To prevent info disclosure and ensure deterministic cross-architecture alignment, all structs MUST use `Pack = 8` and explicit padding.

### C# Contract (Fixed)
```csharp
[StructLayout(LayoutKind.Sequential, Pack = 8)]
public struct EntityState {
    public ulong EntityId;
    public WasmVector3 Position;
    public WasmVector3 Velocity;
    public long ReservedPadding; // Explicit 8-byte alignment
}
```