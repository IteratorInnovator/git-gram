package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

// StateManager handles state token generation and verification.
type StateManager struct {
	secret []byte
}

// NewStateManager creates a new state manager with the given secret.
func NewStateManager(secret string) *StateManager {
	return &StateManager{secret: []byte(secret)}
}

// GenerateToken generates a state token for the given chat ID.
func (s *StateManager) GenerateToken(chatID int64) (string, error) {
	chatIDStr := strconv.FormatInt(chatID, 10)

	mac := hmac.New(sha256.New, s.secret)
	if _, err := mac.Write([]byte(chatIDStr)); err != nil {
		return "", err
	}

	signature := mac.Sum(nil)
	raw := append([]byte(chatIDStr+":"), signature...)
	token := base64.RawURLEncoding.EncodeToString(raw)

	return token, nil
}

// ParseAndVerifyToken parses and verifies a state token, returning the chat ID.
func (s *StateManager) ParseAndVerifyToken(token string) (int64, error) {
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return 0, errors.New("invalid state token")
	}

	parts := strings.SplitN(string(raw), ":", 2)
	if len(parts) != 2 {
		return 0, errors.New("invalid state token")
	}

	payload := parts[0]
	sigBytes := []byte(parts[1])

	mac := hmac.New(sha256.New, s.secret)
	mac.Write([]byte(payload))
	expectedSig := mac.Sum(nil)

	if !hmac.Equal(sigBytes, expectedSig) {
		return 0, errors.New("invalid state token")
	}

	chatID, err := strconv.ParseInt(payload, 10, 64)
	if err != nil {
		return 0, errors.New("invalid state token")
	}

	return chatID, nil
}
