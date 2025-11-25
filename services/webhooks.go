package services

import (
	"github.com/IteratorInnovator/git-gram/telegram"
	"github.com/gofiber/fiber/v2"
)

func HandleTelegramWebhook(c *fiber.Ctx) error {
	update := new(telegram.Update)

	if err := c.BodyParser(update); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "could not parse request body",
		})
	}

	if update.Message == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	var message *telegram.Message = update.Message
	var chat *telegram.Chat = &message.Chat

	telegram.HandleCommands(message.Text, chat.ID)

	return c.SendStatus(fiber.StatusOK)
}

func HandleGitHubWebhook(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
