package repository

import (
	"context"

	"github.com/IteratorInnovator/git-gram/internal/domain"
)

// UserRepository defines the interface for user data access.
type UserRepository interface {
	// Create creates a new user record.
	Create(ctx context.Context, user *domain.User) error

	// GetByChatID retrieves a user by their Telegram chat ID.
	GetByChatID(ctx context.Context, chatID int64) (*domain.User, error)

	// GetByInstallationID retrieves a user by their GitHub installation ID.
	GetByInstallationID(ctx context.Context, installationID int64) (*domain.User, error)

	// UpdateInstallation updates the GitHub installation details for a user.
	UpdateInstallation(ctx context.Context, chatID int64, installationID int64, accountLogin string) error

	// UpdateMuted updates the muted status for a user.
	UpdateMuted(ctx context.Context, chatID int64, muted bool) error

	// Close closes the repository connection.
	Close() error
}
