package events

import (
	"fmt"
)

const deleteBranch string = `🔔 *Branch Deleted in %s*

[%s](%s) deleted branch ` + "`%s`" + ` at %s`

const deleteTag string = `🔔 *New Tag Deleted in %s*

[%s](%s) created tag ` + "`%s`" + ` at %s`

// BuildDeleteInlineKeyboard builds the inline keyboard for a delete event.
func BuildDeleteInlineKeyboard(event *DeleteEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch event.RefType {
	case "branch":
		refURL := fmt.Sprintf(
			"%s/branches",
			event.Repository.HTMLURL,
		)

		keyboardButtons = [][]InlineKeyboardButton{
			{
				InlineKeyboardButton{
					Text: "Repository",
					URL:  event.Repository.HTMLURL,
				},
				InlineKeyboardButton{
					Text: "Branches",
					URL:  refURL,
				},
			},
		}
	case "tag":
		refURL := fmt.Sprintf(
			"%s/tags",
			event.Repository.HTMLURL,
		)
		keyboardButtons = [][]InlineKeyboardButton{
			{
				InlineKeyboardButton{
					Text: "Repository",
					URL:  event.Repository.HTMLURL,
				},
				InlineKeyboardButton{
					Text: "Tags",
					URL:  refURL,
				},
			},
		}
	default:
		keyboardButtons = [][]InlineKeyboardButton{}
	}
	return keyboardButtons
}

// BuildDeleteMessage builds the message for a delete event.
func BuildDeleteMessage(event *DeleteEvent) string {
	var message string

	switch event.RefType {
	case "branch":
		message = fmt.Sprintf(
			deleteBranch,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Ref),
			GetCurrentTimestamp(),
		)
	case "tag":
		message = fmt.Sprintf(
			deleteTag,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Ref),
			GetCurrentTimestamp(),
		)
	default:
		message = ""
	}
	return message
}
