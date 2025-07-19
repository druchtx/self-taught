package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ProjectID string `env:"PROJECT_ID" envDefault:"druchtx-dev"`
}

func main() {
	ctx := context.Background()
	cfg, _ := env.ParseAs[Config]()

	// connect to the default database if not specified
	defaultClient, err := firestore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// specificClient, err := firestore.NewClient(ctx, cfg.ProjectID,database)
	defer func() { _ = defaultClient.Close() }()

	if _, err = defaultClient.Collection("skills").Doc("test").Set(ctx, map[string]any{
		"name":  "go",
		"level": 1000,
	}); err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

}
