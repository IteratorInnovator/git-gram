package config

import (
	"encoding/base64"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// Config holds all application configuration.
type Config struct {
	App       AppConfig
	Telegram  TelegramConfig
	Firestore FirestoreConfig
	GitHub    GitHubConfig
}

// AppConfig contains general application settings.
type AppConfig struct {
	Port        string
	StateSecret string
}

// TelegramConfig contains Telegram bot settings.
type TelegramConfig struct {
	BotToken      string
	APIEndpoint   string
	WebhookURL    string
	WebhookSecret string
}

// FirestoreConfig contains Firestore database settings.
type FirestoreConfig struct {
	ProjectID  string
	DatabaseID string
}

// GitHubConfig contains GitHub App settings.
type GitHubConfig struct {
	AppClientID   string
	PrivateKey    string
	WebhookSecret string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_PRIVATE_KEY_B64"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode GitHub private key: %w", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")

	cfg := &Config{
		App: AppConfig{
			Port:        fmt.Sprintf(":%s", os.Getenv("PORT")),
			StateSecret: os.Getenv("STATE_SECRET"),
		},
		Telegram: TelegramConfig{
			BotToken:      botToken,
			APIEndpoint:   fmt.Sprintf("https://api.telegram.org/bot%s", botToken),
			WebhookURL:    os.Getenv("TELEGRAM_WEBHOOK_URL"),
			WebhookSecret: os.Getenv("TELEGRAM_WEBHOOK_SECRET_TOKEN"),
		},
		Firestore: FirestoreConfig{
			ProjectID:  os.Getenv("GOOGLE_CLOUD_PROJECT_ID"),
			DatabaseID: os.Getenv("FIRESTORE_DATABASE_ID"),
		},
		GitHub: GitHubConfig{
			AppClientID:   os.Getenv("GITHUB_APP_CLIENT_ID"),
			PrivateKey:    string(decodedKey),
			WebhookSecret: os.Getenv("GITHUB_WEBHOOK_SECRET_TOKEN"),
		},
	}

	return cfg, nil
}
