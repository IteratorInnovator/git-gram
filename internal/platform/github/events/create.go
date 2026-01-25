package events

import (
	"fmt"
	"net/url"
)

const createBranch string = `🔔 *New Branch Created in %s*

[%s](%s) created branch ` + "`%s`" + ` at %s`

const createTag string = `🔔 *New Tag Created in %s*

[%s](%s) created tag ` + "`%s`" + ` at %s`

// BuildCreateInlineKeyboard builds the inline keyboard for a create event.
func BuildCreateInlineKeyboard(event *CreateEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch event.RefType {
	case "branch":
		refURL := fmt.Sprintf(
			"%s/tree/%s",
			event.Repository.HTMLURL,
			url.PathEscape(event.Ref),
		)

		keyboardButtons = [][]InlineKeyboardButton{
			{
				InlineKeyboardButton{
					Text: "Repository",
					URL:  event.Repository.HTMLURL,
				},
				InlineKeyboardButton{
					Text: "Branch",
					URL:  refURL,
				},
			},
		}
	case "tag":
		refURL := fmt.Sprintf(
			"%s/releases/tag/%s",
			event.Repository.HTMLURL,
			url.PathEscape(event.Ref),
		)
		keyboardButtons = [][]InlineKeyboardButton{
			{
				InlineKeyboardButton{
					Text: "Repository",
					URL:  event.Repository.HTMLURL,
				},
				InlineKeyboardButton{
					Text: "Tag",
					URL:  refURL,
				},
			},
		}
	default:
		keyboardButtons = [][]InlineKeyboardButton{}
	}
	return keyboardButtons
}

// BuildCreateMessage builds the message for a create event.
func BuildCreateMessage(event *CreateEvent) string {
	var message string

	switch event.RefType {
	case "branch":
		message = fmt.Sprintf(
			createBranch,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Ref),
			FormatRFC3339Timestamp(event.Repository.PushedAt),
		)
	case "tag":
		message = fmt.Sprintf(
			createTag,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Ref),
			FormatRFC3339Timestamp(event.Repository.PushedAt),
		)
	default:
		message = ""
	}
	return message
}
