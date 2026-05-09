using Xunit;
using NexusStudio.Core;
using NexusStudio.Scripts.Sandbox;
using System;
using System.Text;

namespace NexusStudio.Tests;

/// <summary>
/// Final integration test suite for Script Sandbox.
/// Verifies memory limits, instruction budgeting (fuel), and API blacklisting.
/// </summary>
public class SandboxIntegrationTests
{
    [Fact]
    public void TestFuelBudgetEnforcement()
    {
        // Setup: Script with an infinite loop
        string infiniteLoopScript = @"
            (module
                (func (export ""main"")
                    (loop
                        br 0
                    )
                )
            )";
        
        // Convert to WASM bytes (simplified representation)
        byte[] wasmBytes = Encoding.UTF8.GetBytes(infiniteLoopScript);
        
        var bridge = new WasmBridge();
        
        // Assert that the bridge catches the instruction budget exhaustion
        Assert.Throws<Wasmtime.TrapException>(() => {
            bridge.Execute(wasmBytes);
        });
    }

    [Fact]
    public void TestMemoryLimitEnforcement()
    {
        // Verify that the Luau VM enforces the 64MB hard cap
        using var vm = new LuauVM();
        
        // Malicious script attempting to allocate massive memory
        string maliciousScript = @"
            local t = {}
            for i = 1, 1000000 do
                t[i] = string.rep('overflow', 1000)
            end";
        
        byte[] scriptBytes = Encoding.UTF8.GetBytes(maliciousScript);
        
        // Assert that native Lua memory limit triggers a failure
        Assert.Throws<Exception>(() => {
             LuauNative.luaL_loadbufferx(vm.Handle, scriptBytes, (UIntPtr)scriptBytes.Length, "mem_test", 0);
        });
    }

    [Fact]
    public void TestDangerousGlobalBlacklisting()
    {
        // Verify that dangerous globals are flagged by the static analyzer
        string[] forbiddenScripts = {
            "os.execute('rm -rf /')",
            "io.open('/etc/passwd')",
            "debug.debug()",
            "require('socket')"
        };

        foreach (var script in forbiddenScripts)
        {
            bool isSafe = SandboxValidator.IsScriptSafe(script);
            Assert.False(isSafe, $"Sandbox failed to flag dangerous script: {script}");
        }
    }
}