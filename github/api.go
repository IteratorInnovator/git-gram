package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/IteratorInnovator/git-gram/github/events"
)

func FetchInstallationAccountUsername(installation_id int64) (string, error) {
	jwt, err := generateGitHubAppJWT()
    if err != nil {
        return "", err
    }

	url := fmt.Sprintf("https://api.github.com/app/installations/%d", installation_id)

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+ jwt)
    req.Header.Set("Accept", "application/vnd.github+json")
    req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("github get installation failed: status=%d body=%s", resp.StatusCode, string(body))
	}

	var inst events.InstallationResponse
	err = json.NewDecoder(resp.Body).Decode(&inst)
	if err != nil {
		return "", err
	}

    return inst.Account.Login, nil
}

func DeleteAppInstallation(installation_id int64) error {
    jwt, err := generateGitHubAppJWT()
    if err != nil {
        return err
    }

	url := fmt.Sprintf("https://api.github.com/app/installations/%d", installation_id)

    req, err := http.NewRequest(http.MethodDelete, url, nil)
    if err != nil {
        return err
    }

    req.Header.Set("Authorization", "Bearer "+ jwt)
    req.Header.Set("Accept", "application/vnd.github+json")
    req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

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