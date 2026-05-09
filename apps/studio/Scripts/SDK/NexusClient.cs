using Godot;
using NexusStudio.SDK.Services;

namespace NexusStudio.SDK;

/// <summary>
/// Primary singleton entry point for the NexusEngine SDK.
/// Manages platform services and ensures secure access to engine resources.
/// </summary>
public partial class NexusClient : Node
{
    /// <summary>Gets the global SDK instance.</summary>
    public static NexusClient Instance { get; private set; } = null!;

    public NetworkService Network { get; } = new();
    public DataStoreService DataStore { get; } = new();
    public AssetService Assets { get; } = new();

    public override void _EnterTree()
    {
        if (Instance != null && Instance != this) { QueueFree(); return; }
        Instance = this;
        GD.Print("NexusEngine SDK: Initialized.");
    }
}