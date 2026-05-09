package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"github.com/google/uuid"
)

var decoderPool = sync.Pool{
	New: func() interface{} { return json.NewDecoder(nil) },
}

func (s *AssetService) HandleGetUploadURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limitedReader := io.LimitReader(r.Body, 1024)
	dec := decoderPool.Get().(*json.Decoder)
	dec.Reset(limitedReader)
	defer decoderPool.Put(dec)

	var req struct{ FileType string `json:"fileType"` }
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Malformed JSON", http.StatusBadRequest)
		return
	}

	if req.FileType != "model" && req.FileType != "texture" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	assetId := uuid.New().String()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"assetId": assetId,
		"url":     "https://s3.nexusengine.internal/uploads/" + assetId,
	})
}