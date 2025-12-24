package events

import (
	"fmt"
)

const BranchProtectionConfigurationEnabled string = `🔔 *Branch Protection Rules Change in %s*

[%s](%s) enabled branch protection rules at %s`


const BranchProtectionConfigurationDisabled string = `🔔 *Branch Protection Rules Change in %s*

[%s](%s) disabled branch protection rules at %s`


func BuildBranchProtectionConfigurationInlineKeyboard(event *BranchProtectionConfiguration) [][]InlineKeyboardButton {
	return [][]InlineKeyboardButton {
		{ 
			InlineKeyboardButton { 
				Text: "Repository", 
				URL: event.Repository.HTMLURL,
			},
			InlineKeyboardButton {
				Text: "Settings",
				URL: fmt.Sprintf(
					"%s/settings",
					event.Repository.HTMLURL,
				),
			},
		},
	}
}


func BuildBranchProtectionConfigurationMessage(event *BranchProtectionConfiguration) string {
	var message string

	switch (event.Action) {
		case "disabled":
			message = fmt.Sprintf(
				BranchProtectionConfigurationDisabled,
				escapeText(event.Repository.FullName),
				escapeText(event.Sender.Login),
				escapeURL(event.Sender.HTMLURL),
				formatRFC3339Timestamp(event.Repository.UpdatedAt),
			)
		case "enabled":
			message = fmt.Sprintf(
				BranchProtectionConfigurationEnabled,
				escapeText(event.Repository.FullName),
				escapeText(event.Sender.Login),
				escapeURL(event.Sender.HTMLURL),
				formatRFC3339Timestamp(event.Repository.UpdatedAt),
			)
		default:
			message = ""
	}
	return message
}