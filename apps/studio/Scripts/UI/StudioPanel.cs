using Godot;

namespace NexusStudio.UI;

/// <summary>
/// Base class for all dockable editor windows.
/// Enforces 200x150 min size and 64-char title sanitization.
/// </summary>
public partial class StudioPanel : PanelContainer
{
    private string _title = "Nexus Studio";
    
    [Export]
    public string Title 
    { 
        get => _title;
        set => _title = value.Length > 64 ? value.Substring(0, 64) : value; 
    }

    public override void _Ready()
    {
        // Enforce DoS-prevention constraints
        CustomMinimumSize = new Vector2(200, 150);
        
        var margin = new MarginContainer();
        margin.AddThemeConstantOverride("margin_left", 8);
        margin.AddChild(new Label { Text = _title.Replace("\n", " ") });
        AddChild(margin);
    }
}