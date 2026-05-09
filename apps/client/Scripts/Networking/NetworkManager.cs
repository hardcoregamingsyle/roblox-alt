using Godot;

namespace NexusEngine.Networking;

public partial class NetworkManager : Node
{
    private ENetMultiplayerPeer _peer = new();
    private const int MaxClients = 32;

    public override void _Ready()
    {
        var error = _peer.CreateServer(3000, MaxClients);
        if (error != Error.Ok)
        {
            GD.PrintErr($"Failed to initialize ENet server: {error}");
            return;
        }
        Multiplayer.MultiplayerPeer = _peer;
    }

    public override void _Process(double delta)
    {
        if (_peer != null && _peer.GetConnectionStatus() == MultiplayerPeer.ConnectionStatus.Connected)
        {
            _peer.Poll();
        }
    }
}