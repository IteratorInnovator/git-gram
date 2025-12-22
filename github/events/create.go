package events

import (
	"fmt"
	"net/url"
)

const CreateBranch string = `🔔 *New Branch Created in %s*

[%s](%s) created branch ` + "`%s`" + ` at %s`


const CreateTag string = `🔔 *New Tag Created in %s*

[%s](%s) created tag ` + "`%s`" + ` at %s`


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
					InlineKeyboardButton {
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
				CreateBranch,
				escapeText(createEvent.Repository.FullName),
				escapeText(createEvent.Sender.Login),
				escapeURL(createEvent.Sender.HTMLURL),
				escapeText(createEvent.Ref),
				formatRFC3339Timestamp(createEvent.Repository.PushedAt),
			)
		case "tag":
			message = fmt.Sprintf(
				CreateTag,
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