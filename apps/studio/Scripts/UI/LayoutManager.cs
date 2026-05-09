using System;

public class LayoutManager {
    public Vector2 Clamp(Vector2 input, Vector2 min, Vector2 max) {
        if (float.IsNaN(input.X) || float.IsInfinity(input.X)) input.X = min.X;
        if (float.IsNaN(input.Y) || float.IsInfinity(input.Y)) input.Y = min.Y;
        return new Vector2(
            Math.Clamp(input.X, min.X, max.X),
            Math.Clamp(input.Y, min.Y, max.Y)
        );
    }
}