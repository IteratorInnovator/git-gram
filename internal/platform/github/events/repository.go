package events

import (
	"fmt"
	"strings"
)

const repositoryArchived = `🔔 *Repository Archived*

[%s](%s) archived ` + "`%s`" + ` at %s`

const repositoryCreated = `🔔 *New Repository Created*

[%s](%s) created ` + "`%s`" + ` at %s`

const repositoryDeleted = `🔔 *Repository Deleted*

[%s](%s) deleted ` + "`%s`" + ` at %s`

const repositoryEdited = `🔔 *Repository Edited*

[%s](%s) edited ` + "`%s`" + ` at %s

*View Changes*
**>%s||
`

const repositoryPrivatized = `🔔 *Repository Visibility Change*

[%s](%s) changed ` + "`%s`" + ` visibility to private at %s`

const repositoryPublicized = `🔔 *Repository Visibility Change*

[%s](%s) changed ` + "`%s`" + ` visibility to public at %s`

const repositoryRenamed = `🔔 *Repository Renamed*

[%s](%s) renamed ` + "`%s`" + ` to ` + "`%s`" + ` at %s`

const repositoryUnarchived = `🔔 *Repository Unarchived*

[%s](%s) unarchived ` + "`%s`" + ` at %s`

// BuildRepositoryInlineKeyboard builds the inline keyboard for a repository event.
func BuildRepositoryInlineKeyboard(event *RepositoryEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch event.Action {
	case "archived", "created", "edited", "privatized", "publicized", "renamed", "unarchived":
		keyboardButtons = [][]InlineKeyboardButton{
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
	case "deleted":
		keyboardButtons = [][]InlineKeyboardButton{}
	default:
		keyboardButtons = [][]InlineKeyboardButton{}
	}

	return keyboardButtons
}

// BuildRepositoryMessage builds the message for a repository event.
func BuildRepositoryMessage(event *RepositoryEvent) string {
	var message string

	switch event.Action {
	case "archived":
		message = fmt.Sprintf(
			repositoryArchived,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	case "created":
		message = fmt.Sprintf(
			repositoryCreated,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.CreatedAt),
		)
	case "edited":
		changes := parseChanges(event)

		message = fmt.Sprintf(
			repositoryEdited,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
			strings.Join(changes, "\n>"),
		)
	case "deleted":
		message = fmt.Sprintf(
			repositoryDeleted,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	case "privatized":
		message = fmt.Sprintf(
			repositoryPrivatized,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	case "publicized":
		message = fmt.Sprintf(
			repositoryPublicized,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	case "renamed":
		message = fmt.Sprintf(
			repositoryRenamed,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Changes.Repository.Name.From),
			EscapeText(event.Repository.Name),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	case "unarchived":
		message = fmt.Sprintf(
			repositoryUnarchived,
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			EscapeText(event.Repository.FullName),
			FormatRFC3339Timestamp(event.Repository.UpdatedAt),
		)
	default:
		message = ""
	}

	return message
}

func parseChanges(event *RepositoryEvent) []string {
	var changes []string

	if event.Changes.DefaultBranch.From != "" {
		changes = append(
			changes,
			fmt.Sprintf(
				"Default branch: `%s` 🡪 `%s`",
				EscapeText(event.Changes.DefaultBranch.From),
				EscapeText(event.Repository.DefaultBranch),
			),
		)
	}

	if event.Changes.Description.From != nil {
		changes = append(
			changes,
			fmt.Sprintf(
				"Description: `%s` 🡪 `%s`",
				EscapeText(*event.Changes.Description.From),
				EscapeText(*event.Repository.Description),
			),
		)
	}

	if event.Changes.Homepage.From != nil {
		changes = append(
			changes,
			fmt.Sprintf(
				"Homepage: `%s` 🡪 `%s`",
				EscapeText(*event.Changes.Homepage.From),
				EscapeText(*event.Repository.Homepage),
			),
		)
	}

	if event.Changes.Topics.From != nil {
		changes = append(
			changes,
			fmt.Sprintf(
				"Topics: `%s` 🡪 `%s`",
				EscapeText(FormatStringSlice(*event.Changes.Topics.From)),
				EscapeText(FormatInterfaceSlice(event.Repository.Topics)),
			),
		)
	}

	return changes
}
