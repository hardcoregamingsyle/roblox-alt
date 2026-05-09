using System;
using System.Text.Json;

public class SceneSerializer {
    private const int MAX_DEPTH = 32;

    public static void Deserialize(string json) {
        var options = new JsonSerializerOptions {
            MaxDepth = MAX_DEPTH
        };
        
        try {
            // AGGRESSIVE FIX: Enforce recursion depth to prevent StackOverflow/DoS
            JsonDocument.Parse(json, new JsonDocumentOptions {
                MaxDepth = MAX_DEPTH
            });
        } catch (JsonException ex) {
            throw new Exception("Security violation: Payload exceeds recursion depth", ex);
        }
    }
}