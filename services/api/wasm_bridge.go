package main

import (
	"github.com/bytecodealliance/wasmtime-go/v28"
)

// ConfigureWasmSandbox enforces a 256MB memory limit to prevent DoS
func ConfigureWasmSandbox() *wasmtime.Config {
	config := wasmtime.NewConfig()
	// Set memory limit to 256 MiB
	config.SetMemoryLimit(256 * 1024 * 1024)
	return config
}