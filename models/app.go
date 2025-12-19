package models

import (
	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type App struct {
	Router *fiber.App
	Store  *firestore.Client
}

func (app* App) RegisterRoutes() {
	app.Router.Use(requestid.New())
	app.Router.Use(logger.New())
	app.Router.Use(recover.New())

	telegram := app.Router.Group("/telegram")
	github := app.Router.Group("/github")

	telegram.Post("/webhook", func(c *fiber.Ctx) error {
		return services.HandleTelegramWebhook(c, app.Store)
	})
	
	github.Post("/webhook", func(c *fiber.Ctx) error {
		return services.HandleGitHubWebhook(c, app.Store)
	})
	github.Get("/installation/success", func(c *fiber.Ctx) error {
		return services.HandleSuccessfulInstallation(c, app.Store)
	})
}
