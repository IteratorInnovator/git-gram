package github

import (
	"time"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/golang-jwt/jwt/v5"
)

func generateGitHubAppJWT() (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.GithubCfg.GITHUB_APP_PRIVATE_KEY))
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := jwt.RegisteredClaims {
		Issuer:    config.GithubCfg.GITHUB_APP_CLIENT_ID,          // iss
		IssuedAt:  jwt.NewNumericDate(now.Add(-60 * time.Second)), // iat, 60 seconds in the past
		ExpiresAt: jwt.NewNumericDate(now.Add(9 * time.Minute)),   // exp, under 10 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}

func VerifyHMAC256Signature(body []byte, headerVal string, secret string) (bool, error) {
	if headerVal == "" {
		return false, fmt.Errorf("missing X-Hub-Signature-256")
	}

	const prefix = "sha256="
	if !strings.HasPrefix(headerVal, prefix) {
		return false, fmt.Errorf("bad X-Hub-Signature-256 format")
	}

	givenHex := strings.TrimPrefix(headerVal, prefix)
	given, err := hex.DecodeString(givenHex)
	if err != nil {
		return false, fmt.Errorf("signature not hex: %w", err)
	}

	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write(body)
	expected := mac.Sum(nil)

	return hmac.Equal(given, expected), nil
}