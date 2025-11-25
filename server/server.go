package server

import (
	"github.com/IteratorInnovator/git-gram/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// New builds a Fiber app with standard middleware and route registration.
func New() *fiber.App {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	registerRoutes(app)

	return app
}

func registerRoutes(app *fiber.App) {
	telegram := app.Group("/telegram")
	github := app.Group("/github")

	telegram.Post("/webhook", services.HandleTelegramWebhook)
	github.Post("/webhook", services.HandleGitHubWebhook)
}
