using Godot;

namespace NexusStudio.UI;

/// <summary>
/// Constructs the primary Studio IDE interface.
/// </summary>
public partial class StudioLayout : Control
{
    public override void _Ready()
    {
        SetAnchorsPreset(LayoutPreset.FullRect);
        
        var root = new HSplitContainer { AnchorRight = 1.0f, AnchorBottom = 1.0f };
        
        var explorer = new StudioPanel { Title = "Explorer" };
        var inspector = new StudioPanel { Title = "Properties" };
        
        root.AddChild(explorer);
        root.AddChild(inspector);
        
        AddChild(root);
    }
}