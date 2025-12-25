package events

import (
	"fmt"
)

const DeleteBranch string = `🔔 *Branch Deleted in %s*

[%s](%s) deleted branch ` + "`%s`" + ` at %s`


const DeleteTag string = `🔔 *New Tag Deleted in %s*

[%s](%s) created tag ` + "`%s`" + ` at %s`


func BuildDeleteInlineKeyboard(deleteEvent *DeleteEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch (deleteEvent.RefType) {
		case "branch":
			refURL := fmt.Sprintf(
				"%s/branches",
				deleteEvent.Repository.HTMLURL,
			)

			keyboardButtons = [][]InlineKeyboardButton {
				{ 
					InlineKeyboardButton { 
						Text: "Repository", 
						URL: deleteEvent.Repository.HTMLURL,
					},
					InlineKeyboardButton {
						Text: "Branches",
						URL: refURL,
					},
				},
			}
		case "tag":
			refURL := fmt.Sprintf(
				"%s/tags",
				deleteEvent.Repository.HTMLURL,
			)
			keyboardButtons = [][]InlineKeyboardButton {
				{ 
					InlineKeyboardButton { 
						Text: "Repository", 
						URL: deleteEvent.Repository.HTMLURL,
					},
					InlineKeyboardButton {
						Text: "Tags",
						URL: refURL,
					},
				},
			}
		default:
			keyboardButtons = [][]InlineKeyboardButton {}
	}
	return keyboardButtons
}


func BuildDeleteMessage(deleteEvent *DeleteEvent) string {
	var message string

	switch (deleteEvent.RefType) {
		case "branch":
			message = fmt.Sprintf(
				DeleteBranch,
				escapeText(deleteEvent.Repository.FullName),
				escapeText(deleteEvent.Sender.Login),
				escapeURL(deleteEvent.Sender.HTMLURL),
				escapeText(deleteEvent.Ref),
				getCurrentTimestamp(),
			)
		case "tag":
			message = fmt.Sprintf(
				DeleteTag,
				escapeText(deleteEvent.Repository.FullName),
				escapeText(deleteEvent.Sender.Login),
				escapeURL(deleteEvent.Sender.HTMLURL),
				escapeText(deleteEvent.Ref),
				getCurrentTimestamp(),
			)
		default:
			message = ""
	}
	return message
}