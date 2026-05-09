using Godot;
using NexusStudio.UI;
using NexusStudio.Core;

namespace NexusStudio;

/// <summary>
/// Root entry point for NexusStudio.
/// Orchestrates the bootstrap sequence for UI, Sandboxing, and API connectivity.
/// </summary>
public partial class Main : Node
{
    private LuauScriptManager _scriptManager = null!;

    public override void _Ready()
    {
        // 1. Initialize Sandbox Infrastructure
        _scriptManager = new LuauScriptManager();
        AddChild(_scriptManager);
        
        // 2. Initialize UI Workspace
        // The StudioLayout internally handles LayoutManager batching
        var workspace = new StudioLayout();
        AddChild(workspace);
        
        GD.Print("NexusStudio: Editor Workspace successfully initialized.");
    }

    public override void _Notification(int what)
    {
        if (what == NotificationWMCloseRequest)
        {
            // Perform clean shutdown of native resources
            _scriptManager.Dispose();
            GetTree().Quit();
        }
    }
}