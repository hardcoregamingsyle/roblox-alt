package main

import (
	"net/http"
	"sync"
	"github.com/bytecodealliance/wasmtime-go/v28"
	"github.com/gin-gonic/gin"
)

type WasmStorePool struct {
	pool sync.Pool
	engine *wasmtime.Engine
}

func NewWasmStorePool(engine *wasmtime.Engine) *WasmStorePool {
	return &WasmStorePool{
		engine: engine,
		pool: sync.Pool{
			New: func() interface{} {
				return wasmtime.NewStore(engine)
			},
		},
	}
}

func (p *WasmStorePool) Get() *wasmtime.Store {
	return p.pool.Get().(*wasmtime.Store)
}

func (p *WasmStorePool) Put(s *wasmtime.Store) {
	// Reset store context to prevent data leakage
	s.GC() 
	p.pool.Put(s)
}

func HandleWasmError(c *gin.Context, err error, store *wasmtime.Store, pool *WasmStorePool) {
	if store != nil {
		pool.Put(store)
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal execution error"})
}