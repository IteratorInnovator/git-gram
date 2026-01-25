package domain

// User represents a registered user linking Telegram chat to GitHub installation.
type User struct {
	ChatID         int64  `firestore:"chat_id"`
	InstallationID int64  `firestore:"installation_id"`
	AccountLogin   string `firestore:"github_account_username"`
	Muted          bool   `firestore:"muted"`
}

// NewUser creates a new user with default values.
func NewUser(chatID int64) *User {
	return &User{
		ChatID:         chatID,
		InstallationID: 0,
		AccountLogin:   "",
		Muted:          false,
	}
}

// IsLinked returns true if the user has a GitHub installation linked.
func (u *User) IsLinked() bool {
	return u.InstallationID != 0 && u.AccountLogin != ""
}
