package events

import (
	"fmt"
)

const branchProtectionConfigurationEnabled string = `🔔 *Branch Protection Rules Change in %s*

[%s](%s) enabled branch protection rules at %s`

const branchProtectionConfigurationDisabled string = `🔔 *Branch Protection Rules Change in %s*

[%s](%s) disabled branch protection rules at %s`

// BuildBranchProtectionConfigurationInlineKeyboard builds the inline keyboard for a branch protection configuration event.
func BuildBranchProtectionConfigurationInlineKeyboard(event *BranchProtectionConfiguration) [][]InlineKeyboardButton {
	return [][]InlineKeyboardButton{
		{
			InlineKeyboardButton{
				Text: "Repository",
				URL:  event.Repository.HTMLURL,
			},
			InlineKeyboardButton{
				Text: "Settings",
				URL: fmt.Sprintf(
					"%s/settings",
					event.Repository.HTMLURL,
				),
			},
		},
	}
}

// BuildBranchProtectionConfigurationMessage builds the message for a branch protection configuration event.
func BuildBranchProtectionConfigurationMessage(event *BranchProtectionConfiguration) string {
	var message string

	switch event.Action {
	case "disabled":
		message = fmt.Sprintf(
			branchProtectionConfigurationDisabled,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			GetCurrentTimestamp(),
		)
	case "enabled":
		message = fmt.Sprintf(
			branchProtectionConfigurationEnabled,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			GetCurrentTimestamp(),
		)
	default:
		message = ""
	}
	return message
}
