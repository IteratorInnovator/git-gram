package github

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/IteratorInnovator/git-gram/internal/platform/github/events"
	"github.com/IteratorInnovator/git-gram/internal/platform/telegram"
)

// Service handles GitHub webhook events.
type Service struct {
	telegram *telegram.Client
}

// NewService creates a new GitHub service.
func NewService(tg *telegram.Client) *Service {
	return &Service{telegram: tg}
}

// HandleWebhookEvent processes a GitHub webhook event.
func (s *Service) HandleWebhookEvent(event string, chatID int64, ctx *fiber.Ctx) error {
	switch event {
	case "branch_protection_configuration":
		return s.handleBranchProtectionConfigurationEvent(ctx, chatID)
	case "create":
		return s.handleCreateEvent(ctx, chatID)
	case "delete":
		return s.handleDeleteEvent(ctx, chatID)
	case "push":
		return s.handlePushEvent(ctx, chatID)
	case "pull_request":
		fmt.Printf("github event handler: pull_request event for chat id=%d\n", chatID)
		return nil
	case "repository":
		return s.handleRepositoryEvent(ctx, chatID)
	default:
		fmt.Printf("github event handler: unknown event=%v chat id=%d\n", event, chatID)
		return nil
	}
}

func (s *Service) handleBranchProtectionConfigurationEvent(ctx *fiber.Ctx, chatID int64) error {
	var event events.BranchProtectionConfiguration
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}

	keyboard := events.BuildBranchProtectionConfigurationInlineKeyboard(&event)
	message := events.BuildBranchProtectionConfigurationMessage(&event)

	return s.telegram.SendMessage(chatID, message, keyboard)
}

func (s *Service) handleCreateEvent(ctx *fiber.Ctx, chatID int64) error {
	var event events.CreateEvent
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}

	keyboard := events.BuildCreateInlineKeyboard(&event)
	message := events.BuildCreateMessage(&event)

	return s.telegram.SendMessage(chatID, message, keyboard)
}

func (s *Service) handleDeleteEvent(ctx *fiber.Ctx, chatID int64) error {
	var event events.DeleteEvent
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}

	keyboard := events.BuildDeleteInlineKeyboard(&event)
	message := events.BuildDeleteMessage(&event)

	return s.telegram.SendMessage(chatID, message, keyboard)
}

func (s *Service) handlePushEvent(ctx *fiber.Ctx, chatID int64) error {
	var event events.PushEvent
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}

	keyboard := events.BuildPushInlineKeyboard(&event)
	message := events.BuildPushMessage(&event)

	return s.telegram.SendMessage(chatID, message, keyboard)
}

func (s *Service) handleRepositoryEvent(ctx *fiber.Ctx, chatID int64) error {
	var event events.RepositoryEvent
	if err := ctx.BodyParser(&event); err != nil {
		return err
	}

	keyboard := events.BuildRepositoryInlineKeyboard(&event)
	message := events.BuildRepositoryMessage(&event)

	return s.telegram.SendMessage(chatID, message, keyboard)
}
