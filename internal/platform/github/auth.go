package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/IteratorInnovator/git-gram/internal/config"
)

// GenerateAppJWT generates a JWT for GitHub App authentication.
func GenerateAppJWT(cfg config.GitHubConfig) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(cfg.PrivateKey))
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := jwt.RegisteredClaims{
		Issuer:    cfg.AppClientID,
		IssuedAt:  jwt.NewNumericDate(now.Add(-60 * time.Second)),
		ExpiresAt: jwt.NewNumericDate(now.Add(9 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}

// VerifyWebhookSignature verifies the HMAC-SHA256 signature from GitHub webhooks.
func VerifyWebhookSignature(body []byte, signature string, secret string) (bool, error) {
	if signature == "" {
		return false, fmt.Errorf("missing X-Hub-Signature-256")
	}

	const prefix = "sha256="
	if !strings.HasPrefix(signature, prefix) {
		return false, fmt.Errorf("bad X-Hub-Signature-256 format")
	}

	givenHex := strings.TrimPrefix(signature, prefix)
	given, err := hex.DecodeString(givenHex)
	if err != nil {
		return false, fmt.Errorf("signature not hex: %w", err)
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expected := mac.Sum(nil)

	return hmac.Equal(given, expected), nil
}
