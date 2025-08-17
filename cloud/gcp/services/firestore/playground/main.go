package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	ProjectID string `env:"PROJECT_ID" envDefault:"druchtx-local"`
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

	skillsRef := defaultClient.Collection("skills")
	if _, err = skillsRef.Doc("test").Set(ctx, map[string]any{
		"name":  "go",
		"level": 1000,
	}); err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}

	_, _, _ = skillsRef.Add(ctx, map[string]any{
		"name":  "elixir",
		"level": 9999,
	})
	refs, _ := skillsRef.DocumentRefs(ctx).GetAll()

	snaps, _ := defaultClient.GetAll(ctx, refs)

	for _, snap := range snaps {
		log.Println(snap.Data())
	}
}
