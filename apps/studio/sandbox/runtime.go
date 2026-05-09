// Wasmtime implementation for secure script execution
package sandbox

import (
	"github.com/bytecodealliance/wasmtime-go/v2"
)

func NewSandbox() (*wasmtime.Store, error) {
	engine := wasmtime.NewEngine()
	return wasmtime.NewStore(engine), nil
}
// Note: Host bindings must be explicitly registered here via Linker