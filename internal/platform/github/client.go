package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/IteratorInnovator/git-gram/internal/config"
)

const (
	apiBaseURL   = "https://api.github.com"
	apiVersion   = "2022-11-28"
	acceptHeader = "application/vnd.github+json"
)

// Client provides methods to interact with the GitHub API.
type Client struct {
	cfg config.GitHubConfig
}

// NewClient creates a new GitHub API client.
func NewClient(cfg config.GitHubConfig) *Client {
	return &Client{cfg: cfg}
}

// InstallationAccount represents the account associated with an installation.
type InstallationAccount struct {
	Login string `json:"login"`
}

// InstallationResponse represents a GitHub App installation.
type InstallationResponse struct {
	ID      int64               `json:"id"`
	Account InstallationAccount `json:"account"`
}

// GetInstallationAccount fetches the account username for an installation.
func (c *Client) GetInstallationAccount(installationID int64) (string, error) {
	jwt, err := GenerateAppJWT(c.cfg)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/app/installations/%d", apiBaseURL, installationID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	c.setHeaders(req, jwt)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("github get installation failed: status=%d body=%s", resp.StatusCode, string(body))
	}

	var inst InstallationResponse
	if err := json.NewDecoder(resp.Body).Decode(&inst); err != nil {
		return "", err
	}

	return inst.Account.Login, nil
}

// DeleteInstallation removes a GitHub App installation.
func (c *Client) DeleteInstallation(installationID int64) error {
	jwt, err := GenerateAppJWT(c.cfg)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/app/installations/%d", apiBaseURL, installationID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	c.setHeaders(req, jwt)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("could not delete installation")
	}
	return nil
}

// setHeaders sets the standard headers for GitHub API requests.
func (c *Client) setHeaders(req *http.Request, jwt string) {
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
}
