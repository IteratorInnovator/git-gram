package github

import (
	"time"

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
		Issuer:    config.GithubCfg.GITHUB_APP_CLIENT_ID,      // iss
		IssuedAt:  jwt.NewNumericDate(now.Add(-60 * time.Second)), // iat, 60 seconds in the past
		ExpiresAt: jwt.NewNumericDate(now.Add(9 * time.Minute)),   // exp, under 10 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}