package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/github/events"
	"github.com/gofiber/fiber/v2"
)


func HandleGitHubWebhookEvent(event string, chatId int64, ctx *fiber.Ctx) error {
	var err error = nil

	switch (event) {
		case "branch_protection_configuration":
			err = handleBranchProtectionConfigurationEvent(ctx, chatId)
		case "create":
			err = handleCreateEvent(ctx, chatId)
		case "delete":
			err = handleDeleteEvent(ctx, chatId)
		case "push":
			err = handlePushEvent(ctx, chatId)
		case "pull_request":
			fmt.Printf("github event handler: pull_request event for chat id=%d\n", chatId)
			fmt.Println("pull request event")
		case "repository":
			err = handleRepositoryEvent(ctx, chatId)
		default:
			fmt.Printf("github event handler: unknown event=%v chat id=%d\n", event, chatId)
			fmt.Printf("%v event", event)
	}
	return err
}


func handleBranchProtectionConfigurationEvent(ctx *fiber.Ctx, chatId int64) error {	
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var event events.BranchProtectionConfiguration
	err := ctx.BodyParser(&event)
	if err != nil {
		return err
	}

	keyboardButtons := events.BuildBranchProtectionConfigurationInlineKeyboard(&event)
	message := events.BuildBranchProtectionConfigurationMessage(&event)

	payload := struct {
		ChatID      int                         `json:"chat_id"`
		ParseMode   string                      `json:"parse_mode"`
		Text        string                      `json:"text"`
		ReplyMarkup events.InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
		ReplyMarkup: events.InlineKeyboardMarkup{
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

	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("telegram response status code: %d", resp.StatusCode)
	}

	return nil
}


func handleCreateEvent(ctx *fiber.Ctx, chatId int64) error {	
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var createEvent events.CreateEvent
	err := ctx.BodyParser(&createEvent)
	if err != nil {
		return err
	}

	keyboardButtons := events.BuildCreateInlineKeyboard(&createEvent)
	message := events.BuildCreateMessage(&createEvent)

	payload := struct {
		ChatID      int                         `json:"chat_id"`
		ParseMode   string                      `json:"parse_mode"`
		Text        string                      `json:"text"`
		ReplyMarkup events.InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
		ReplyMarkup: events.InlineKeyboardMarkup{
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

	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("telegram response status code: %d", resp.StatusCode)
	}

	return nil
}


func handleDeleteEvent(ctx *fiber.Ctx, chatId int64) error {	
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var deleteEvent events.DeleteEvent
	err := ctx.BodyParser(&deleteEvent)
	if err != nil {
		return err
	}

	keyboardButtons := events.BuildDeleteInlineKeyboard(&deleteEvent)
	message := events.BuildDeleteMessage(&deleteEvent)
	
	payload := struct {
		ChatID      int                         `json:"chat_id"`
		ParseMode   string                      `json:"parse_mode"`
		Text        string                      `json:"text"`
		ReplyMarkup events.InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
		ReplyMarkup: events.InlineKeyboardMarkup{
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

	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("telegram response status code: %d", resp.StatusCode)
	}

	return nil
}


func handlePushEvent(ctx *fiber.Ctx, chatId int64) error {	
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var pushEvent events.PushEvent
	err := ctx.BodyParser(&pushEvent)
	if err != nil {
		return err
	}

	keyboardButtons := events.BuildPushInlineKeyboard(&pushEvent)
	message := events.BuildPushMessage(&pushEvent)
	
	payload := struct {
		ChatID      int                         `json:"chat_id"`
		ParseMode   string                      `json:"parse_mode"`
		Text        string                      `json:"text"`
		ReplyMarkup events.InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
		ReplyMarkup: events.InlineKeyboardMarkup{
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

	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("telegram response status code: %d", resp.StatusCode)
	}

	return nil
}


func handleRepositoryEvent(ctx *fiber.Ctx, chatId int64) error {
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var repositoryEvent events.RepositoryEvent
	err := ctx.BodyParser(&repositoryEvent)
	if err != nil {
		return err
	}

	keyboardButtons := events.BuildRepositoryInlineKeyboard(&repositoryEvent)
	message := events.BuildRepositoryMessage(&repositoryEvent)

	payload := struct {
		ChatID      int                         `json:"chat_id"`
		ParseMode   string                      `json:"parse_mode"`
		Text        string                      `json:"text"`
		ReplyMarkup events.InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: message,
		ReplyMarkup: events.InlineKeyboardMarkup{
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

	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("telegram response status code: %d", resp.StatusCode)
	}

	return nil
}