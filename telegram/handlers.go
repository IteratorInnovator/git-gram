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
	"github.com/IteratorInnovator/git-gram/github"
)

func HandlePostInstallation(ctx context.Context, client *firestore.Client, installation_id int64, stateToken string) error {
	chatId, err := parseAndVerifyStateToken(stateToken)
	if err != nil {
		return err
	}

	account_username, err := github.FetchInstallationAccountUsername(installation_id)
	if err != nil {
		return err
	}

	db.UpdateInstallation(ctx, client, chatId, installation_id, account_username)

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	payload := struct {
		ChatID      int     `json:"chat_id"`
		ParseMode   string  `json:"parse_mode"`
		Text        string  `json:"text"`
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
		err = handleStatus(ctx, client, chatId)
	case "/mute":
		err = handleMute(ctx, client, chatId)
	case "/unmute":
		err = handleUnmute(ctx, client, chatId)
	case "/unlink":
		err = handleUnlink(ctx, client, chatId)
	case "/help":
		err = handleHelp(chatId)
	default:
		err = handleInvalidCommand(chatId)
	}
	return err
}

func handleStart(ctx context.Context, client *firestore.Client, chatId int64) error {
	db.SaveChat(ctx, client, chatId)

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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

func handleStatus(ctx context.Context, client *firestore.Client, chatId int64) error {
	account_login, muted, err := db.FetchUserInfo(ctx, client, chatId)

	var message string
	if err != nil && err.Error() == "user doc does not exist" {
		message = StatusDocNotFoundMessage
	} else if account_login == "" {
		message = StatusNoInstallationMessage
	} else {
		var muted_text string
		var muted_info_text string
		if muted {
			muted_text = "muted"
			muted_info_text = "Use /unmute to resume notifications"
		} else {
			muted_text = "unmuted"
			muted_info_text = "Use /mute to stop notifications"
		}
		
		message = fmt.Sprintf(
			StatusInstalledTemplateMessage,
			account_login,
			muted_text,
			muted_info_text,
		)
	}

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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

func handleUnlink(ctx context.Context, client *firestore.Client, chatId int64) error {
	var message string = UnlinkSuccessMessage
	installation_id, err := db.FetchInstallationId(ctx, client, chatId)
	if err != nil {
		message = UnlinkFailedMessage
	} else if installation_id == 0 {
		message = UnlinkNotInstalledMessage
	} else {
		err = github.DeleteAppInstallation(installation_id)
		if err != nil {
			message = UnlinkFailedMessage
		}

		db.UpdateInstallation(ctx, client, chatId, 0, "")
	}

	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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

func handleHelp(chatId int64) error {
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

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
