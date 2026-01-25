package telegram

// Update represents a Telegram webhook update.
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

// Message represents a Telegram message.
type Message struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

// Chat represents a Telegram chat.
type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// WebhookInfo represents Telegram webhook configuration.
type WebhookInfo struct {
	URL            string   `json:"url"`
	AllowedUpdates []string `json:"allowed_updates"`
	SecretToken    string   `json:"secret_token"`
	MaxConnections int      `json:"max_connections"`
}

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents a button in an inline keyboard.
type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

// SendMessageRequest represents a request to send a message.
type SendMessageRequest struct {
	ChatID      int64                 `json:"chat_id"`
	ParseMode   string                `json:"parse_mode"`
	Text        string                `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
