package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/telegram"
	"github.com/gofiber/fiber/v2"
)

func GetTelegramWebhook() (bool, error) {
	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "getWebhook")

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	
	var webhook telegram.WebHookInfo

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&webhook)
	if err != nil {
		return false, err
	}

	if webhook.URL != "" || webhook.SecretToken != config.TelegramCfg.TELEGRAM_WEBHOOK_SECRET_TOKEN {
		return true, nil
	}
	return false, nil
}

func SetTelegramWebHook() error {
	exists, err := GetTelegramWebhook()
	if err != nil {
		return err
	}
	if exists {
		err = DeleteTelegramWebhook()
	}
	if err != nil {
		return nil
	}

	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "setWebhook")

	payload := telegram.WebHookInfo {
		URL: config.TelegramCfg.TELEGRAM_WEBHOOK_URL,
		SecretToken: config.TelegramCfg.TELEGRAM_WEBHOOK_SECRET_TOKEN,
		AllowedUpdates: []string{"message"},
		MaxConnections: 100,
	}

	reqBody, err := json.Marshal(payload)
	if (err != nil) {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil  {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func DeleteTelegramWebhook() error {
	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "deleteWebhook")

	payload := struct {
		DropPendingUpdates bool `json:"drop_pending_updates"`
	} {
		DropPendingUpdates: true,
	}

	reqBody, err := json.Marshal(payload)
	if (err != nil) {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil  {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func HandleTelegramWebhook(c *fiber.Ctx, client *firestore.Client) error {
	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()
	
	update := new(telegram.Update)

	if err := c.BodyParser(update); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "could not parse request body",
		})
	}

	if update.Message == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	var message *telegram.Message = update.Message
	var chat *telegram.Chat = &message.Chat

	if err := telegram.HandleCommand(ctx, client, message.Text, chat.ID); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  "could not send reply",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func HandleGitHubWebhook(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
