using Godot;

namespace NexusStudio.Core;

public interface ILuauSandboxAPI
{
    void Print(string message);
    void SetNodePosition(string nodePath, Vector2 position);
    float GetDeltaTime();
}