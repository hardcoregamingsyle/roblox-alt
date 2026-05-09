using System.Collections.Generic;
using NexusStudio.Core;

namespace NexusStudio.Persistence;

public class SceneValidator
{
    private const int MaxDepth = 64;

    public bool IsValid(SceneBundle bundle)
    {
        // Pre-allocate based on estimated count to avoid re-allocations
        var visited = new HashSet<ulong>(bundle.RootEntities.Count);
        foreach (var entity in bundle.RootEntities)
        {
            if (!IterativeValidate(entity, visited)) return false;
        }
        return true;
    }

    private bool IterativeValidate(EntityData root, HashSet<ulong> visited)
    {
        var stack = new Stack<(EntityData, int)>(MaxDepth);
        stack.Push((root, 0));
        while(stack.Count > 0) {
            var (node, depth) = stack.Pop();
            if (depth > MaxDepth) return false;
            if (!visited.Add(node.Id)) return false;
            foreach (var child in node.Children) stack.Push((child, depth + 1));
        }
        return true;
    }
}