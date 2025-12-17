package config

import (
	"fmt"
	"os"
	"encoding/base64"

	_ "github.com/joho/godotenv/autoload"
)

var TelegramCfg TelegramConfig
var FirestoreCfg FirestoreConfig
var GithubCfg GithubConfig
var AppCfg AppConfig

func InitEnv() error {
	TelegramCfg = TelegramConfig {
		TELEGRAM_BOT_API_TOKEN: os.Getenv("TELEGRAM_BOT_API_TOKEN"),
		TELEGRAM_BOT_API_ENDPOINT: os.Getenv("TELEGRAM_BOT_API_ENDPOINT"),
		TELEGRAM_WEBHOOK_URL: os.Getenv("TELEGRAM_WEBHOOK_URL"),
		TELEGRAM_WEBHOOK_SECRET_TOKEN: os.Getenv("TELEGRAM_WEBHOOK_SECRET_TOKEN"),
	}
	TelegramCfg.TELEGRAM_BOT_API_BASE_URL = fmt.Sprintf("%vbot%v/", TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, TelegramCfg.TELEGRAM_BOT_API_TOKEN)

	FirestoreCfg = FirestoreConfig {
		GOOGLE_CLOUD_PROJECT_ID: os.Getenv("GOOGLE_CLOUD_PROJECT_ID"),
		FIRESTORE_DATABASE_ID: os.Getenv("FIRESTORE_DATABASE_ID"),
	}

	GithubCfg = GithubConfig {
		GITHUB_APP_CLIENT_ID: os.Getenv("GITHUB_APP_CLIENT_ID"),
		GITHUB_WEBHOOK_SECRET_TOKEN: os.Getenv("GITHUB_WEBHOOK_SECRET_TOKEN"),
	}
	fmt.Println(os.Getenv("GITHUB_APP_PRIVATE_KEY_B64")) // debug print statement
	decodedKey, err := base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_PRIVATE_KEY_B64"))
	if err != nil {
		return err 
	}
	GithubCfg.GITHUB_APP_PRIVATE_KEY = string(decodedKey)

	AppCfg = AppConfig {
		PORT: fmt.Sprintf(":%v", os.Getenv("PORT")),
		STATE_SECRET: fmt.Sprintf("%v", os.Getenv("STATE_SECRET")),
	}
	return nil
}