package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/db"
)

func HandlePostInstallation(ctx context.Context, client *firestore.Client, installation_id int64, stateToken string) error {
	chatId, err := parseAndVerifyStateToken(stateToken)
	if err != nil {
		return err
	}
	
	go db.SaveInstallation(ctx, client, chatId, installation_id)

	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	payload := struct {
		ChatID      int                  `json:"chat_id"`
		ParseMode   string               `json:"parse_mode"`
		Text        string               `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: SuccessfulInstallationMessage,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func HandleCommand(ctx context.Context, client *firestore.Client, command string, chatId int64) error {
	var err error = nil

	switch command {
	case "/start":
		err = handleStart(ctx, client, chatId)
	case "/status":
		err = handleStatus(chatId)
	case "/mute":
		err = handleMute(ctx, client, chatId)
	case "/unmute":
		err = handleUnmute(ctx, client, chatId)
	case "/unlink":
		err = handleUnlink(chatId)
	case "/help":
		err = handleHelp(chatId)
	default:
		err = handleInvalidCommand(chatId)
	}
	return err
}

func handleStart(ctx context.Context, client *firestore.Client, chatId int64) error {
	go db.SaveChat(ctx, client, chatId)

	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	stateToken, err := generateStateToken(chatId)
	if err != nil {
		return err
	}

	installationURL := fmt.Sprintf("https://github.com/apps/git-gram-67/installations/new?state=%s", neturl.QueryEscape(stateToken))

	keyboardButtons := [][]InlineKeyboardButton {
		{ 
			InlineKeyboardButton { 
				Text: "Install Git Gram App", 
				URL: installationURL,
			},
		},
	}
	payload := struct {
		ChatID      int                  `json:"chat_id"`
		ParseMode   string               `json:"parse_mode"`
		Text        string               `json:"text"`
		ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: InstallationMessage,
		ReplyMarkup: InlineKeyboardMarkup{
			InlineKeyboard: keyboardButtons,
		},
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func handleStatus(chatId int64) error {
	return nil
}

func handleMute(ctx context.Context, client *firestore.Client, chatId int64) error {
	isSetup, err := db.Mute(ctx, client, chatId)
	
	var message string = ""
	if !isSetup {
		message = MuteBeforeStartErrorMessage
	} else if err != nil {
		message = DefaultErrorMessage
	} else {
		message = MuteSuccessMessage
	}

	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	payload := struct {
		ChatID    int    `json:"chat_id"`
		ParseMode string `json:"parse_mode"`
		Text      string `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func handleUnmute(ctx context.Context, client *firestore.Client, chatId int64) error {
	isSetup, err := db.Unmute(ctx, client, chatId)
	
	var message string = ""
	if !isSetup {
		message = UnmuteBeforeStartErrorMessage
	} else if err != nil {
		message = DefaultErrorMessage
	} else {
		message = UnmuteSuccessMessage
	}

	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	payload := struct {
		ChatID    int    `json:"chat_id"`
		ParseMode string `json:"parse_mode"`
		Text      string `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func handleUnlink(chatId int64) error {
	return nil
}

func handleHelp(chatId int64) error {
	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	payload := struct {
		ChatID    int    `json:"chat_id"`
		ParseMode string `json:"parse_mode"`
		Text      string `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: HelpMessage,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func handleInvalidCommand(chatId int64) error {
	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "sendMessage")

	payload := struct {
		ChatID    int    `json:"chat_id"`
		ParseMode string `json:"parse_mode"`
		Text      string `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: InvalidCommandMessage,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	return nil
}
