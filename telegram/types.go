package telegram

type Update struct {
    UpdateID int      `json:"update_id"`
    Message  *Message `json:"message,omitempty"`
}

type Message struct {
    MessageID int             `json:"message_id"`
    From      *User           `json:"from"`
    Chat      Chat            `json:"chat"`
    Date      int64           `json:"date"`
    Text      string          `json:"text"`
    Entities  []MessageEntity `json:"entities,omitempty"`
}

type User struct {
    ID       int64  `json:"id"`
    IsBot    bool   `json:"is_bot"`
    Username string `json:"username"`
}

type Chat struct {
    ID   int64  `json:"id"`
    Type string `json:"type"`
}

type MessageEntity struct {
    Offset int    `json:"offset"`
    Length int    `json:"length"`
    Type   string `json:"type"`
}

type WebHookInfo struct {
	URL            string   `json:"url"`
	AllowedUpdates []string `json:"allowed_updates"`
	SecretToken    string   `json:"secret_token"`
	MaxConnections int      `json:"max_connections"`
}

type InlineKeyboardMarkup struct {
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}