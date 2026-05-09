using System;
using Wasmtime;

public class SandboxEnvironment : IDisposable
{
    private readonly Engine _engine;
    private readonly Store _store;
    private readonly Module _module;
    private readonly Linker _linker;
    private readonly Instance _instance;

    public SandboxEnvironment(byte[] wasmBytes)
    {
        _engine = new Engine();
        _store = new Store(_engine);
        _module = Module.FromBytes(_engine, "user_script", wasmBytes);
        _linker = new Linker(_engine);
        
        // Explicitly deny access to OS/FileSystem by not defining imports
        _instance = _linker.Instantiate(_store, _module);
    }

    public void Execute(string entryPoint)
    {
        var run = _instance.GetFunction(_store, entryPoint);
        if (run == null) throw new InvalidOperationException("Invalid entry point");
        run.Invoke(_store);
    }

    public void Dispose()
    {
        _store.Dispose();
        _engine.Dispose();
    }
}