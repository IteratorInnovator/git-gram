package config

import (
	"fmt"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

var TelegramCfg TelegramConfig
var FirestoreCfg FirestoreConfig
var AppCfg AppConfig

func InitEnv() {
	TelegramCfg = TelegramConfig{
		TELEGRAM_BOT_API_TOKEN: os.Getenv("TELEGRAM_BOT_API_TOKEN"),
		TELEGRAM_BOT_API_ENDPOINT: os.Getenv("TELEGRAM_BOT_API_ENDPOINT"),
		TELEGRAM_WEBHOOK_URL: os.Getenv("TELEGRAM_WEBHOOK_URL"),
		TELEGRAM_WEBHOOK_SECRET_TOKEN: os.Getenv("TELEGRAM_WEBHOOK_SECRET_TOKEN"),
	}
	TelegramCfg.TELEGRAM_BOT_API_BASE_URL = fmt.Sprintf("%vbot%v/", TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, TelegramCfg.TELEGRAM_BOT_API_TOKEN)

	FirestoreCfg = FirestoreConfig{
		GOOGLE_CLOUD_PROJECT_ID: os.Getenv("GOOGLE_CLOUD_PROJECT_ID"),
		FIRESTORE_DATABASE_ID: os.Getenv("FIRESTORE_DATABASE_ID"),
	}

	AppCfg = AppConfig{
		PORT: fmt.Sprintf(":%v", os.Getenv("PORT")),
	}
}