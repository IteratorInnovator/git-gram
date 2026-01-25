package telegram

import (
	"context"
	"fmt"
	"net/url"

	"github.com/IteratorInnovator/git-gram/internal/domain"
	"github.com/IteratorInnovator/git-gram/internal/platform/github"
	"github.com/IteratorInnovator/git-gram/internal/platform/telegram"
	"github.com/IteratorInnovator/git-gram/internal/repository"
)

// Service handles Telegram bot commands and interactions.
type Service struct {
	repo         repository.UserRepository
	telegram     *telegram.Client
	github       *github.Client
	stateManager *StateManager
}

// NewService creates a new Telegram service.
func NewService(repo repository.UserRepository, tg *telegram.Client, gh *github.Client, stateSecret string) *Service {
	return &Service{
		repo:         repo,
		telegram:     tg,
		github:       gh,
		stateManager: NewStateManager(stateSecret),
	}
}

// HandleCommand processes a Telegram command.
func (s *Service) HandleCommand(ctx context.Context, command string, chatID int64) error {
	switch command {
	case "/start":
		return s.handleStart(ctx, chatID)
	case "/status":
		return s.handleStatus(ctx, chatID)
	case "/mute":
		return s.handleMute(ctx, chatID)
	case "/unmute":
		return s.handleUnmute(ctx, chatID)
	case "/unlink":
		return s.handleUnlink(ctx, chatID)
	case "/help":
		return s.handleHelp(chatID)
	default:
		return s.handleInvalidCommand(chatID)
	}
}

// HandlePostInstallation processes a successful GitHub App installation.
func (s *Service) HandlePostInstallation(ctx context.Context, installationID int64, stateToken string) error {
	chatID, err := s.stateManager.ParseAndVerifyToken(stateToken)
	if err != nil {
		return err
	}

	accountUsername, err := s.github.GetInstallationAccount(installationID)
	if err != nil {
		return err
	}

	if err := s.repo.UpdateInstallation(ctx, chatID, installationID, accountUsername); err != nil {
		return err
	}

	return s.telegram.SendMessage(chatID, SuccessfulInstallationMessage, nil)
}

func (s *Service) handleStart(ctx context.Context, chatID int64) error {
	user := domain.NewUser(chatID)
	s.repo.Create(ctx, user)

	stateToken, err := s.stateManager.GenerateToken(chatID)
	if err != nil {
		return err
	}

	installationURL := fmt.Sprintf("https://github.com/apps/git-gram-67/installations/new?state=%s", url.QueryEscape(stateToken))

	keyboard := [][]telegram.InlineKeyboardButton{
		{
			telegram.InlineKeyboardButton{
				Text: "Install Git Gram App",
				URL:  installationURL,
			},
		},
	}

	return s.telegram.SendMessage(chatID, InstallationMessage, keyboard)
}

func (s *Service) handleStatus(ctx context.Context, chatID int64) error {
	user, err := s.repo.GetByChatID(ctx, chatID)

	var message string
	if err == repository.ErrUserNotFound {
		message = StatusDocNotFoundMessage
	} else if user.AccountLogin == "" {
		message = StatusNoInstallationMessage
	} else {
		var mutedText string
		var mutedInfoText string
		if user.Muted {
			mutedText = "muted"
			mutedInfoText = "Use /unmute to resume notifications"
		} else {
			mutedText = "unmuted"
			mutedInfoText = "Use /mute to stop notifications"
		}

		message = fmt.Sprintf(
			StatusInstalledTemplateMessage,
			user.AccountLogin,
			mutedText,
			mutedInfoText,
		)
	}

	return s.telegram.SendMessage(chatID, message, nil)
}

func (s *Service) handleMute(ctx context.Context, chatID int64) error {
	err := s.repo.UpdateMuted(ctx, chatID, true)

	var message string
	if err == repository.ErrUserNotFound {
		message = MuteBeforeStartErrorMessage
	} else if err != nil {
		message = DefaultErrorMessage
	} else {
		message = MuteSuccessMessage
	}

	return s.telegram.SendMessage(chatID, message, nil)
}

func (s *Service) handleUnmute(ctx context.Context, chatID int64) error {
	err := s.repo.UpdateMuted(ctx, chatID, false)

	var message string
	if err == repository.ErrUserNotFound {
		message = UnmuteBeforeStartErrorMessage
	} else if err != nil {
		message = DefaultErrorMessage
	} else {
		message = UnmuteSuccessMessage
	}

	return s.telegram.SendMessage(chatID, message, nil)
}

func (s *Service) handleUnlink(ctx context.Context, chatID int64) error {
	message := UnlinkSuccessMessage

	user, err := s.repo.GetByChatID(ctx, chatID)
	if err != nil {
		message = UnlinkFailedMessage
	} else if user.InstallationID == 0 {
		message = UnlinkNotInstalledMessage
	} else {
		if err := s.github.DeleteInstallation(user.InstallationID); err != nil {
			message = UnlinkFailedMessage
		}
		s.repo.UpdateInstallation(ctx, chatID, 0, "")
	}

	return s.telegram.SendMessage(chatID, message, nil)
}

func (s *Service) handleHelp(chatID int64) error {
	return s.telegram.SendMessage(chatID, HelpMessage, nil)
}

func (s *Service) handleInvalidCommand(chatID int64) error {
	return s.telegram.SendMessage(chatID, InvalidCommandMessage, nil)
}
