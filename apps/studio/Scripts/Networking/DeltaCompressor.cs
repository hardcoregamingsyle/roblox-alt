using System;
using Godot;

public class DeltaCompressor {
    public Vector2 Sanitize(Vector2 input) {
        if (float.IsNaN(input.X) || float.IsInfinity(input.X)) input.X = 0;
        if (float.IsNaN(input.Y) || float.IsInfinity(input.Y)) input.Y = 0;
        return input;
    }
}