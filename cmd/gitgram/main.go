package main

import (
	"context"
	"log"

	"github.com/IteratorInnovator/git-gram/internal/config"
	"github.com/IteratorInnovator/git-gram/internal/handler"
	"github.com/IteratorInnovator/git-gram/internal/platform/telegram"
	"github.com/IteratorInnovator/git-gram/internal/repository"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	telegramClient := telegram.NewClient(cfg.Telegram)
	if err := telegramClient.SetWebhook(); err != nil {
		log.Fatalf("Failed to set Telegram webhook: %v", err)
	}

	repo, err := repository.NewFirestore(ctx, cfg.Firestore)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer repo.Close()

	server := handler.NewServer(cfg, repo, telegramClient)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
