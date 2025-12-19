package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/db"
	"github.com/IteratorInnovator/git-gram/github"
	"github.com/IteratorInnovator/git-gram/telegram"
	"github.com/gofiber/fiber/v2"
)

func GetTelegramWebhook() (bool, error) {
	url := fmt.Sprintf("%v%v", config.TelegramCfg.TELEGRAM_BOT_API_BASE_URL, "getWebhookInfo")

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	
	var respBody struct {
        OK     bool                 `json:"ok"`
        Result telegram.WebHookInfo `json:"result"`
    }

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respBody)
	if err != nil {
		return false, err
	}
	
	if !respBody.OK {
		return false, fmt.Errorf("telegram getWebhookInfo returned not ok")
	}
	var webhook telegram.WebHookInfo = respBody.Result

	// if webhook url is empty, or different from the env tg webhook url, or secret token mismatches
	// return false, indicating to delete webhook and set a new one
	if webhook.URL == "" || webhook.URL != config.TelegramCfg.TELEGRAM_WEBHOOK_URL || webhook.SecretToken != config.TelegramCfg.TELEGRAM_WEBHOOK_SECRET_TOKEN {
		return false, nil
	}
	return true, nil
}

func SetTelegramWebHook() error {
	exists, err := GetTelegramWebhook()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	err = DeleteTelegramWebhook()
	if err != nil {
		return err
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
	token := c.Get("X-Telegram-Bot-Api-Secret-Token", "")
	if token == "" || token != config.TelegramCfg.TELEGRAM_WEBHOOK_SECRET_TOKEN {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error": "unauthorized",
		})
	}

	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()
	
	update := new(telegram.Update)

	if err := c.BodyParser(update); err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
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
			"error":  err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func HandleGitHubWebhook(c *fiber.Ctx, client *firestore.Client) error {
	sig := c.Get("X-Hub-Signature-256")
	body := c.Body()

	ok, err := github.VerifyHMAC256Signature(body, sig, config.GithubCfg.GITHUB_WEBHOOK_SECRET_TOKEN)
	if err != nil || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": fiber.StatusUnauthorized,
			"error":  err.Error(),
		})
	}

	var installation struct {
		Installation struct {
			Id int64 `json:"id"`
		} `json:"installation"`
	}

	if err := c.BodyParser(installation) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	chatId, muted, err := db.FetchChatIdAndMute(ctx, client, installation.Installation.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	if muted {
		return c.SendStatus(fiber.StatusNoContent)
	}

	event := c.Get("X-GitHub-Event")
	err = github.HandleGitHubWebhookEvent(event, chatId, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func HandleSuccessfulInstallation(c *fiber.Ctx, client *firestore.Client) error {
	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	var installation_id int64 = int64(c.QueryInt("installation_id"))
	var stateToken string = c.Query("state")
	err := telegram.HandlePostInstallation(ctx, client, installation_id, stateToken)
	if err != nil {
		c.Set(fiber.HeaderContentType, "application/problem+json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	
	return c.SendStatus(fiber.StatusOK)
}