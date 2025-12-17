package main

import (
	"context"
	"log"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/db"
	"github.com/IteratorInnovator/git-gram/models"
	"github.com/IteratorInnovator/git-gram/services"
	"github.com/gofiber/fiber/v2"
)



func main() {
	ctx := context.Background()

	if err := config.InitEnv() ; err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	err := services.SetTelegramWebHook()
	if err != nil {
		log.Fatalf("Failed to set Telegram Webhook: %v", err)
	}

	firestoreClient, err := db.CreateClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	app := &models.App {
		Router: fiber.New(),
		Store: firestoreClient,
	}

	app.RegisterRoutes()

	if err := app.Router.Listen(config.AppCfg.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
