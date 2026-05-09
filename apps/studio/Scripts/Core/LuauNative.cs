using System;
using System.Runtime.InteropServices;

namespace NexusStudio.Core;

/// <summary>
/// Provides P/Invoke bindings for the native Luau C library.
/// </summary>
internal static class LuauNative
{
    private const string LibName = "luau";

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern IntPtr luaL_newstate();

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_setmemorylimit(IntPtr L, UIntPtr limit);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl, CharSet = CharSet.Ansi)]
    public static extern int luaL_loadbufferx(IntPtr L, [MarshalAs(UnmanagedType.LPArray)] byte[] buff, UIntPtr sz, string name, int mode);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_close(IntPtr L);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_pushnil(IntPtr L);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_setglobal(IntPtr L, string name);
}