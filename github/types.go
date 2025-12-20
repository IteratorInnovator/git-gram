package github

import (
    "time"
)

type installationResponse struct {
    ID      int64 `json:"id"`
    Account struct {
        Login string `json:"login"`
    } `json:"account"`
}

type InlineKeyboardMarkup struct {
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type PushEvent struct {
	Ref     string `json:"ref"`
	Compare string `json:"compare"`

	Repository struct {
		Name     string `json:"name"`      // for title %s (optional, but useful)
		FullName string `json:"full_name"` // IteratorInnovator/git-gram
		HTMLURL  string `json:"html_url"`  // repo link button
        PushedAt int64  `json:"pushed_at"` // unix seconds
	} `json:"repository"`

	Pusher struct {
		Name string `json:"name"` // fallback actor name
	} `json:"pusher"`

	Sender struct {
		Login   string `json:"login"`    // actor label
		HTMLURL string `json:"html_url"` // actor profile link
	} `json:"sender"`

	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"commits"`

	HeadCommit struct {
		ID        string    `json:"id"`        // sha for latest
		Message   string    `json:"message"`   // latest commit message
		Timestamp time.Time `json:"timestamp"` // time for "at %s"
		URL       string    `json:"url"`       // optional, for [View commit]
	} `json:"head_commit"`
}