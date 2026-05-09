package main

import (
	"encoding/json"
	"sync"
)

// Centralized decoder pool to minimize heap allocations
var DecoderPool = sync.Pool{
	New: func() interface{} { return json.NewDecoder(nil) },
}