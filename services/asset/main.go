package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/**
 * @file main.go
 * @description Entry point for the Asset Storage Service.
 * Acts as a security broker for generating S3-compatible pre-signed upload URLs.
 */

func main() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("Failed to load AWS configuration: " + err.Error())
	}

	s3Client := s3.NewFromConfig(cfg)
	svc := &AssetService{
		s3Client: s3Client,
		presign:  s3.NewPresignClient(s3Client),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/assets/upload", svc.HandleGetUploadURL)

	server := &http.Server{
		Addr:         ":3001",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}