public void Dispose() {
    Dispose(true);
    GC.SuppressFinalize(this);
}

protected virtual void Dispose(bool disposing) {
    if (_luaState != IntPtr.Zero) {
        LuauNative.lua_close(_luaState);
        _luaState = IntPtr.Zero;
    }
}

~LuauVM() {
    Dispose(false);
}