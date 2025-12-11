package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/db"
)

func HandleCommand(ctx context.Context, client *firestore.Client, command string, chatId int64) error {
	var err error = nil

	switch command {
	case "/start":
		err = handleStart(ctx, client, chatId)
	case "/status":
		err = handleStatus(chatId)
	case "/mute":
		err = handleMute(chatId)
	case "/unmute":
		err = handleUnmute(chatId)
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

	payload := struct {
		ChatID    int    `json:"chat_id"`
		ParseMode string `json:"parse_mode"`
		Text      string `json:"text"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: InstallationMessage,
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

func handleMute(chatId int64) error {
	return nil
}

func handleUnmute(chatId int64) error {
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
