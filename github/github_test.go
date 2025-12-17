package github

import (
	"testing"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func TestGenerateGitHubAppJWT(t *testing.T) {
	// For the test, config.GithubCfg must be initialised.
	_ = godotenv.Load("../.env")

	config.InitEnv()
	if err := config.InitEnv() ; err != nil {
		t.Fatalf("Error loading environment variables: %v", err)
	}
	if config.GithubCfg.GITHUB_APP_PRIVATE_KEY == "" {
		t.Fatal("GITHUB_APP_PRIVATE_KEY not set for tests")
	}
	if config.GithubCfg.GITHUB_APP_CLIENT_ID == "" {
		t.Fatal("GITHUB_APP_CLIENT_ID not set for tests")
	}

	tokenStr, err := generateGitHubAppJWT()
	if err != nil {
		t.Fatalf("generateGitHubAppJWT error: %v", err)
	}
	if tokenStr == "" {
		t.Fatal("expected non empty JWT")
	}

	// Parse without verifying signature just to inspect claims
	parser := jwt.NewParser()
	var claims jwt.RegisteredClaims
	_, _, err = parser.ParseUnverified(tokenStr, &claims)
	if err != nil {
		t.Fatalf("failed to parse JWT: %v", err)
	}

	if claims.Issuer != config.GithubCfg.GITHUB_APP_CLIENT_ID {
		t.Fatalf("unexpected iss: got %q, want %q",
			claims.Issuer, config.GithubCfg.GITHUB_APP_CLIENT_ID)
	}

	if claims.ExpiresAt == nil || claims.IssuedAt == nil {
		t.Fatal("expected iat and exp to be set")
	}
	t.Log(tokenStr)
}

func TestFetchInstallationAccountLogin(t *testing.T) {
	_ = godotenv.Load("../.env")

	if err := config.InitEnv() ; err != nil {
		t.Fatalf("Error loading environment variables: %v", err)
	}
	if config.GithubCfg.GITHUB_APP_PRIVATE_KEY == "" {
		t.Fatal("GITHUB_APP_PRIVATE_KEY not set for tests")
	}
	if config.GithubCfg.GITHUB_APP_CLIENT_ID == "" {
		t.Fatal("GITHUB_APP_CLIENT_ID not set for tests")
	}

	const installationID int64 = 99379392 // change to existing installation id
	account_username, err := FetchInstallationAccountUsername(installationID)

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	} 
	t.Log(account_username)
}