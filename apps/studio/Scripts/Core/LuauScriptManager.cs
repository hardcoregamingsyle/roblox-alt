using Godot;
using System.Collections.Concurrent;

namespace NexusStudio.Core;

/// <summary>
/// Manages a thread-safe pool of LuauVM instances to minimize GC pressure.
/// </summary>
public partial class LuauScriptManager : Node
{
    private readonly ConcurrentStack<LuauVM> _pool = new();

    public override void _Ready()
    {
        for (int i = 0; i < 64; i++) _pool.Push(new LuauVM());
    }

    public LuauVM Acquire() => _pool.TryPop(out var vm) ? vm : new LuauVM();
    
    public void Release(LuauVM vm)
    {
        vm.ResetState();
        _pool.Push(vm);
    }
}