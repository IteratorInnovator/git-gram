package config

type FirestoreConfig struct {
	GOOGLE_CLOUD_PROJECT_ID string
	FIRESTORE_DATABASE_ID string
}

type TelegramConfig struct {
	TELEGRAM_BOT_API_TOKEN string
	TELEGRAM_BOT_API_ENDPOINT string
	TELEGRAM_BOT_API_BASE_URL string
}

type AppConfig struct {
	PORT string
}