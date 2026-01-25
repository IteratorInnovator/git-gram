package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IteratorInnovator/git-gram/internal/config"
)

// Client provides methods to interact with the Telegram Bot API.
type Client struct {
	cfg config.TelegramConfig
}

// NewClient creates a new Telegram API client.
func NewClient(cfg config.TelegramConfig) *Client {
	return &Client{cfg: cfg}
}

// Config returns the Telegram configuration.
func (c *Client) Config() config.TelegramConfig {
	return c.cfg
}

// SendMessage sends a message to a Telegram chat.
func (c *Client) SendMessage(chatID int64, text string, keyboard [][]InlineKeyboardButton) error {
	url := fmt.Sprintf("%s/sendMessage", c.cfg.APIEndpoint)

	req := SendMessageRequest{
		ChatID:    chatID,
		ParseMode: "MarkdownV2",
		Text:      text,
	}

	if keyboard != nil {
		req.ReplyMarkup = &InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		}
	}

	return c.post(url, req)
}

// SetWebhook configures the Telegram webhook.
func (c *Client) SetWebhook() error {
	exists, err := c.getWebhookInfo()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	if err := c.deleteWebhook(); err != nil {
		return err
	}

	url := fmt.Sprintf("%s/setWebhook", c.cfg.APIEndpoint)
	payload := WebhookInfo{
		URL:            c.cfg.WebhookURL,
		SecretToken:    c.cfg.WebhookSecret,
		AllowedUpdates: []string{"message"},
		MaxConnections: 100,
	}

	return c.post(url, payload)
}

// getWebhookInfo checks if the current webhook configuration matches.
func (c *Client) getWebhookInfo() (bool, error) {
	url := fmt.Sprintf("%s/getWebhookInfo", c.cfg.APIEndpoint)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var respBody struct {
		OK     bool        `json:"ok"`
		Result WebhookInfo `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return false, err
	}

	if !respBody.OK {
		return false, fmt.Errorf("telegram getWebhookInfo returned not ok")
	}

	webhook := respBody.Result
	if webhook.URL == "" || webhook.URL != c.cfg.WebhookURL || webhook.SecretToken != c.cfg.WebhookSecret {
		return false, nil
	}
	return true, nil
}

// deleteWebhook removes the current webhook.
func (c *Client) deleteWebhook() error {
	url := fmt.Sprintf("%s/deleteWebhook", c.cfg.APIEndpoint)
	payload := struct {
		DropPendingUpdates bool `json:"drop_pending_updates"`
	}{
		DropPendingUpdates: true,
	}

	return c.post(url, payload)
}

// post sends a POST request with JSON body.
func (c *Client) post(url string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
