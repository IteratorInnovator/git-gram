package repository

import (
	"context"
	"errors"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/IteratorInnovator/git-gram/internal/config"
	"github.com/IteratorInnovator/git-gram/internal/domain"
)

const usersCollection = "users"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUpdateFailed = errors.New("failed to update user")
)

// FirestoreRepository implements UserRepository using Google Cloud Firestore.
type FirestoreRepository struct {
	client *firestore.Client
}

// NewFirestore creates a new FirestoreRepository instance.
func NewFirestore(ctx context.Context, cfg config.FirestoreConfig) (*FirestoreRepository, error) {
	client, err := firestore.NewClientWithDatabase(ctx, cfg.ProjectID, cfg.DatabaseID)
	if err != nil {
		return nil, err
	}
	return &FirestoreRepository{client: client}, nil
}

// Close closes the Firestore client connection.
func (r *FirestoreRepository) Close() error {
	return r.client.Close()
}

// Create creates a new user record in Firestore.
func (r *FirestoreRepository) Create(ctx context.Context, user *domain.User) error {
	docID := strconv.FormatInt(user.ChatID, 10)
	_, err := r.client.Collection(usersCollection).Doc(docID).Create(ctx, user)
	return err
}

// GetByChatID retrieves a user by their Telegram chat ID.
func (r *FirestoreRepository) GetByChatID(ctx context.Context, chatID int64) (*domain.User, error) {
	docID := strconv.FormatInt(chatID, 10)
	snap, err := r.client.Collection(usersCollection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, ErrUserNotFound
	}
	if !snap.Exists() {
		return nil, ErrUserNotFound
	}

	var user domain.User
	if err := snap.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByInstallationID retrieves a user by their GitHub installation ID.
func (r *FirestoreRepository) GetByInstallationID(ctx context.Context, installationID int64) (*domain.User, error) {
	query := r.client.Collection(usersCollection).Where("installation_id", "==", installationID).Limit(1)
	iter := query.Documents(ctx)

	snap, err := iter.Next()
	if err == iterator.Done {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	var user domain.User
	if err := snap.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateInstallation updates the GitHub installation details for a user.
func (r *FirestoreRepository) UpdateInstallation(ctx context.Context, chatID int64, installationID int64, accountLogin string) error {
	docID := strconv.FormatInt(chatID, 10)
	docRef := r.client.Collection(usersCollection).Doc(docID)

	snap, _ := docRef.Get(ctx)
	if !snap.Exists() {
		return ErrUserNotFound
	}

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "installation_id", Value: installationID},
		{Path: "github_account_username", Value: accountLogin},
	})
	if err != nil {
		return ErrUpdateFailed
	}
	return nil
}

// UpdateMuted updates the muted status for a user.
func (r *FirestoreRepository) UpdateMuted(ctx context.Context, chatID int64, muted bool) error {
	docID := strconv.FormatInt(chatID, 10)
	docRef := r.client.Collection(usersCollection).Doc(docID)

	snap, _ := docRef.Get(ctx)
	if !snap.Exists() {
		return ErrUserNotFound
	}

	_, err := docRef.Update(ctx, []firestore.Update{
		{Path: "muted", Value: muted},
	})
	if err != nil {
		return ErrUpdateFailed
	}
	return nil
}
