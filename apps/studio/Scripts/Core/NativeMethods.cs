using System;
using System.Runtime.InteropServices;

namespace NexusStudio.Core;

internal static class NativeMethods
{
    private const string LibName = "luau";

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern IntPtr lua_tolstring(IntPtr L, int index, out UIntPtr len);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_pushcclosure(IntPtr L, IntPtr fn, string name, int n);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    public static extern void lua_setglobal(IntPtr L, string name);
}