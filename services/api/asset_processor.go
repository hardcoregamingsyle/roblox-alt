package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleUpload(c *gin.Context) {
	// Fix: Apply MaxBytesReader BEFORE any read
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
	
	baseDir := "/tmp/uploads"
	
	// Create sharded directory structure
	id := uuid.New().String()
	hash := sha256.Sum256([]byte(id))
	shard := hex.EncodeToString(hash[:2])
	targetDir := filepath.Join(baseDir, shard)
	os.MkdirAll(targetDir, 0700)

	targetPath := filepath.Join(targetDir, id)
	
	// Read header for content type
	header := make([]byte, 512)
	n, err := io.ReadFull(c.Request.Body, header)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	contentType := http.DetectContentType(header[:n])
	if !strings.Contains(contentType, "model") && !strings.Contains(contentType, "texture") {
		c.AbortWithStatus(http.StatusUnsupportedMediaType)
		return
	}

	out, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, io.MultiReader(bytes.NewReader(header[:n]), c.Request.Body)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}