package main

import (
	"context"
	"log"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/db"
	"github.com/IteratorInnovator/git-gram/server"
)

func main() {
	ctx := context.Background()

	config.InitEnv()

	firestoreClient, err := db.CreateClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	app := server.New()

	if err := app.Listen(config.AppCfg.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
