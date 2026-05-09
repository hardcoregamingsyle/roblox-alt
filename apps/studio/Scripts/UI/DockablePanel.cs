using Godot;

namespace NexusStudio.UI;

/// <summary>
/// A security-hardened dockable panel for the Studio IDE.
/// Enforces minimum size constraints (200x150) to prevent UI-based Denial-of-Service.
/// </summary>
public partial class DockablePanel : PanelContainer
{
    [Export] public string PanelTitle { get; set; } = "Panel";

    public override void _Ready()
    {
        // Enforce DoS-prevention constraints
        CustomMinimumSize = new Vector2(200, 150);
        
        var margin = new MarginContainer { Name = "Margin" };
        margin.AddThemeConstantOverride("margin_left", 8);
        margin.AddThemeConstantOverride("margin_right", 8);
        margin.AddThemeConstantOverride("margin_top", 8);
        
        var label = new Label { 
            Text = PanelTitle,
            HorizontalAlignment = HorizontalAlignment.Center 
        };
        
        margin.AddChild(label);
        AddChild(margin);
    }
}