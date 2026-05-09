using System;
using System.Runtime.InteropServices;

namespace NexusStudio.Core;

/**
 * @class LuauBridge
 * @description Bridges the Godot C# runtime with the native Luau VM.
 * Provides a strictly whitelisted API surface for user scripts.
 */
public static class LuauBridge
{
    private static ILuauSandboxAPI? _sandboxApi;

    /// <summary>
    /// Initializes the bridge with a whitelisted API instance.
    /// </summary>
    /// <param name="luaState">Pointer to the native lua_State.</param>
    /// <param name="api">The implementation of allowed engine methods.</param>
    public static void Initialize(IntPtr luaState, ILuauSandboxAPI api)
    {
        _sandboxApi = api;
        // Register whitelisted functions (e.g., 'print', 'SetPosition')
    }

    [UnmanagedFunctionPointer(CallingConvention.Cdecl)]
    private delegate int LuaCFunction(IntPtr L);
}