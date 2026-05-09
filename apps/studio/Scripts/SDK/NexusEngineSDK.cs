using System;
using System.Buffers;
using NexusStudio.SDK.Services;

namespace NexusStudio.SDK;

/// <summary>
/// The primary SDK entry point for game developers.
/// Enforces strict memory pooling and capability-based access.
/// </summary>
public sealed class NexusEngineSDK
{
    private static readonly Lazy<NexusEngineSDK> _instance = new(() => new NexusEngineSDK());
    public static NexusEngineSDK Instance => _instance.Value;

    public NetworkService Network { get; } = new();
    public DataStoreService DataStore { get; } = new();
    public AssetService Assets { get; } = new();

    private NexusEngineSDK() { }
}