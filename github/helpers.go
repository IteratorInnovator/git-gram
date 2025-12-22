package github

import (
	"strings"
	"time"
	"fmt"
	"net/url"

	"github.com/IteratorInnovator/git-gram/github/message_templates"
)


// unixSec is seconds since epoch (GitHub repository.pushed_at).
// Example output: "Thu, 18 Dec 2025, 1:03 AM SGT"
func formatUnixTimestamp(unixSec int64) string {
	if unixSec <= 0 {
		return ""
	}

	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		loc = time.FixedZone("SGT", 8*60*60)
	}

	t := time.Unix(unixSec, 0).In(loc)
	return t.Format("Mon, 2 Jan 2006, 3:04 PM") + " SGT"
}


func formatRFC3339Timestamp(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		loc = time.FixedZone("SGT", 8*60*60)
	}

	return t.In(loc).Format("Mon, 2 Jan 2006, 3:04 PM") + " SGT"
}


func formatRef(ref string) string {
	if ref == "" {
		return ""
	}
	if strings.HasPrefix(ref, "refs/heads/") {
		return strings.TrimPrefix(ref, "refs/heads/")
	}
	if strings.HasPrefix(ref, "refs/tags/") {
		return strings.TrimPrefix(ref, "refs/tags/")
	}
	if strings.HasPrefix(ref, "refs/") {
		return strings.TrimPrefix(ref, "refs/")
	}
	return ref
}

func shortenSHA(sha string) string {
	if len(sha) <= 7 {
		return sha
	}
	return sha[:7]
}


// escapeText escapes a string for Telegram MarkdownV2 "normal text" context.
func escapeText(s string) string {
	// Escape backslash first, then the rest.
	r := strings.NewReplacer(
		`\\`, `\\\\`, // if your input can already contain escapes, keep this; otherwise use "\" -> "\\"
		`\`, `\\`,
		`_`, `\_`,
		`*`, `\*`,
		`[`, `\[`,
		`]`, `\]`,
		`(`, `\(`,
		`)`, `\)`,
		`~`, `\~`,
		"`", "\\`",
		`>`, `\>`,
		`#`, `\#`,
		`+`, `\+`,
		`-`, `\-`,
		`=`, `\=`,
		`|`, `\|`,
		`{`, `\{`,
		`}`, `\}`,
		`.`, `\.`,
		`!`, `\!`,
	)
	return r.Replace(s)
}


// escapeURL escapes the URL part inside [text](url) in MarkdownV2.
func escapeURL(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		`)`, `\)`,
	)
	return r.Replace(s)
}


func BuildPushInlineKeyboard(pushEvent *PushEvent) [][]InlineKeyboardButton {
	return [][]InlineKeyboardButton {
		{ 
			InlineKeyboardButton { 
				Text: "View Commit", 
				URL: pushEvent.HeadCommit.URL,
			},
			InlineKeyboardButton{
				Text: "Changes",
				URL: pushEvent.Compare,
			},
		},
		{ 
			InlineKeyboardButton { 
				Text: "Repository", 
				URL: pushEvent.Repository.HTMLURL,
			},
			InlineKeyboardButton {
				Text: "Branch",
				URL: pushEvent.Repository.HTMLURL + "/tree/" + formatRef(pushEvent.Ref),
			},
		},
	}
}


func BuildPushMessage(pushEvent *PushEvent) string {
	var commitCount int = len(pushEvent.Commits)

	if commitCount > 1 {
		return fmt.Sprintf(
			message_templates.MultipleCommitsPush,
			escapeText(pushEvent.Repository.FullName),
			escapeText(pushEvent.Sender.Login),
			escapeURL(pushEvent.Sender.HTMLURL),
			commitCount,
			formatRef(pushEvent.Ref),
			formatUnixTimestamp(pushEvent.Repository.PushedAt),
			shortenSHA(pushEvent.HeadCommit.ID),
			escapeText(pushEvent.HeadCommit.Message),
		)
	}
	return fmt.Sprintf(
		message_templates.SingleCommitPush,
		escapeText(pushEvent.Repository.FullName),
		escapeText(pushEvent.Sender.Login),
		escapeURL(pushEvent.Sender.HTMLURL),
		formatRef(pushEvent.Ref),
		formatUnixTimestamp(pushEvent.Repository.PushedAt),
		shortenSHA(pushEvent.HeadCommit.ID),
		escapeText(pushEvent.HeadCommit.Message),
	)
}


func BuildCreateInlineKeyboard(createEvent *CreateEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch (createEvent.RefType) {
		case "branch":
			refURL := fmt.Sprintf(
				"%s/tree/%s",
				createEvent.Repository.HTMLURL,
				url.PathEscape(createEvent.Ref),
			)

			keyboardButtons = [][]InlineKeyboardButton {
				{ 
					InlineKeyboardButton { 
						Text: "Repository", 
						URL: createEvent.Repository.HTMLURL,
					},
					InlineKeyboardButton{
						Text: "Branch",
						URL: refURL,
					},
				},
			}
		case "tag":
			refURL := fmt.Sprintf(
				"%s/releases/tag/%s",
				createEvent.Repository.HTMLURL,
				url.PathEscape(createEvent.Ref),
			)
			keyboardButtons = [][]InlineKeyboardButton {
				{ 
					InlineKeyboardButton { 
						Text: "Repository", 
						URL: createEvent.Repository.HTMLURL,
					},
					InlineKeyboardButton{
						Text: "Tag",
						URL: refURL,
					},
				},
			}
		default:
			keyboardButtons = [][]InlineKeyboardButton {}
	}
	return keyboardButtons
}


func BuildCreateMessage(createEvent *CreateEvent) string {
	var message string

	switch (createEvent.RefType) {
		case "branch":
			message = fmt.Sprintf(
				message_templates.CreateBranch,
				escapeText(createEvent.Repository.FullName),
				escapeText(createEvent.Sender.Login),
				escapeURL(createEvent.Sender.HTMLURL),
				escapeText(createEvent.Ref),
				formatRFC3339Timestamp(createEvent.Repository.PushedAt),
			)
		case "tag":
			message = fmt.Sprintf(
				message_templates.CreateTag,
				escapeText(createEvent.Repository.FullName),
				escapeText(createEvent.Sender.Login),
				escapeURL(createEvent.Sender.HTMLURL),
				escapeText(createEvent.Ref),
				formatRFC3339Timestamp(createEvent.Repository.PushedAt),
			)
		default:
			message = ""
	}
	return message
}