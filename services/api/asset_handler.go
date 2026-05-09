package main

import (
	"net/http"
	"strings"
)

var allowedFileTypes = map[string]bool{
	"model":   true,
	"texture": true,
	"audio":   true,
}

func HandleAssetUpload(w http.ResponseWriter, r *http.Request) {
	type UploadReq struct {
		FileType string `json:"fileType"`
	}
	
	var req UploadReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Whitelist validation prevents path traversal (e.g., ../../../)
	if !allowedFileTypes[req.FileType] || strings.ContainsAny(req.FileType, "/\\.") {
		http.Error(w, "Invalid file type specified", http.StatusBadRequest)
		return
	}

	// Logic to proceed with safe fileType
}