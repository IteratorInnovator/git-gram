package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/IteratorInnovator/git-gram/internal/config"
	"github.com/IteratorInnovator/git-gram/internal/platform/github"
	"github.com/IteratorInnovator/git-gram/internal/platform/telegram"
	"github.com/IteratorInnovator/git-gram/internal/repository"
	githubsvc "github.com/IteratorInnovator/git-gram/internal/service/github"
	telegramsvc "github.com/IteratorInnovator/git-gram/internal/service/telegram"
)

// Server represents the HTTP server and its dependencies.
type Server struct {
	cfg      *config.Config
	router   *fiber.App
	repo     repository.UserRepository
	telegram *telegram.Client
	github   *github.Client

	telegramService *telegramsvc.Service
	githubService   *githubsvc.Service
}

// NewServer creates a new HTTP server.
func NewServer(cfg *config.Config, repo repository.UserRepository, tg *telegram.Client) *Server {
	gh := github.NewClient(cfg.GitHub)

	s := &Server{
		cfg:      cfg,
		router:   fiber.New(),
		repo:     repo,
		telegram: tg,
		github:   gh,

		telegramService: telegramsvc.NewService(repo, tg, gh, cfg.App.StateSecret),
		githubService:   githubsvc.NewService(tg),
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	return s.router.Listen(s.cfg.App.Port)
}

func (s *Server) setupMiddleware() {
	s.router.Use(requestid.New())
	s.router.Use(logger.New())
	s.router.Use(recover.New())
}

func (s *Server) setupRoutes() {
	telegramGroup := s.router.Group("/telegram")
	githubGroup := s.router.Group("/github")

	telegramGroup.Post("/webhook", s.handleTelegramWebhook)
	githubGroup.Post("/webhook", s.handleGitHubWebhook)
	githubGroup.Get("/installation/success", s.handleGitHubInstallationSuccess)
}

func (s *Server) handleTelegramWebhook(c *fiber.Ctx) error {
	token := c.Get("X-Telegram-Bot-Api-Secret-Token", "")
	if token == "" || token != s.cfg.Telegram.WebhookSecret {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  "unauthorized",
		})
	}

	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	update := new(telegram.Update)
	if err := c.BodyParser(update); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if update.Message == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	message := update.Message
	chat := &message.Chat

	if err := s.telegramService.HandleCommand(ctx, message.Text, chat.ID); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Server) handleGitHubWebhook(c *fiber.Ctx) error {
	sig := c.Get("X-Hub-Signature-256")
	body := c.Body()

	ok, err := github.VerifyWebhookSignature(body, sig, s.cfg.GitHub.WebhookSecret)
	if err != nil || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  err.Error(),
		})
	}

	var installation struct {
		Installation struct {
			ID int64 `json:"id"`
		} `json:"installation"`
	}

	if err := c.BodyParser(&installation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	user, err := s.repo.GetByInstallationID(ctx, installation.Installation.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if user.Muted {
		return c.SendStatus(fiber.StatusNoContent)
	}

	event := c.Get("X-GitHub-Event")
	if err := s.githubService.HandleWebhookEvent(event, user.ChatID, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Server) handleGitHubInstallationSuccess(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	installationID := int64(c.QueryInt("installation_id"))
	stateToken := c.Query("state")

	if err := s.telegramService.HandlePostInstallation(ctx, installationID, stateToken); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
