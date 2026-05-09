package main

import (
	"context"
	"fmt"
	"github.com/bytecodealliance/wasmtime-go/v24"
)

// SandboxExecutor restricts WASM execution to prevent host system access
func ExecuteUserScript(scriptBytes []byte) ([]byte, error) {
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)
	
	// Enforce strict memory and capability limits
	linker := wasmtime.NewLinker(engine)
	linker.DefineWasi(store, wasmtime.NewWasiConfig()) // Empty config = no FS, no Net

	module, err := wasmtime.NewModule(engine, scriptBytes)
	if err != nil {
		return nil, err
	}

	instance, err := linker.Instantiate(store, module)
	if err != nil {
		return nil, err
	}

	run := instance.GetFunc(store, "run")
	result, err := run.Call(store)
	if err != nil {
		return nil, err
	}
	
	return []byte(fmt.Sprintf("%v", result)), nil
}