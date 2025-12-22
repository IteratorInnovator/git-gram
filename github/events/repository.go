package events

import "fmt"

const RepositoryArchived = `🔔 *Repository Archived*

[%s](%s) archived ` + "`%s`" + ` at %s`


const RepositoryCreated = `🔔 *New Repository Created*

[%s](%s) created ` + "`%s`" + ` at %s`


const RepositoryDeleted = `🔔 *Repository Deleted*

[%s](%s) deleted ` + "`%s`" + ` at %s`


const RepositoryEdited = `🔔 *Repository Edited*

[%s](%s) edited ` + "`%s`" + ` at %s`


const RepositoryPrivatized = `🔔 *Repository Visibility Change*

[%s](%s) changed ` + "`%s`" + ` visibility to private at %s`


const RepositoryPublicized = `🔔 *Repository Visibility Change*

[%s](%s) changed ` + "`%s`" + ` visibility to public at %s`


const RepositoryRenamed = `🔔 *Repository Renamed*

[%s](%s) renamed ` + "`%s`" + ` to ` + "`%s`" + ` at %s`


const RepositoryUnarchived = `🔔 *Repository Unarchived*

[%s](%s) unarchived ` + "`%s`" + ` at %s`


func BuildRepositoryInlineKeyboard(repositoryEvent *RepositoryEvent) [][]InlineKeyboardButton {
	var keyboardButtons [][]InlineKeyboardButton

	switch (repositoryEvent.Action) {
		case "archived", "created", "edited", "privatized", "publicized", "renamed", "unarchived":
			keyboardButtons = [][]InlineKeyboardButton {
				{ 
					InlineKeyboardButton { 
						Text: "Repository", 
						URL: repositoryEvent.Repository.HTMLURL,
					},
					InlineKeyboardButton {
						Text: "Settings",
						URL: fmt.Sprintf(
							"%s/settings",
							repositoryEvent.Repository.HTMLURL,
						),
					},
				},
			}
		case "deleted":
			keyboardButtons = [][]InlineKeyboardButton {}
		default:
			keyboardButtons = [][]InlineKeyboardButton {}
	}

	return keyboardButtons
}

func BuildRepositoryMessage(repositoryEvent *RepositoryEvent) string {
	var message string

	switch (repositoryEvent.Action) {
		case "archived":
			message = fmt.Sprintf(
				RepositoryArchived,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		case "created":
			message = fmt.Sprintf(
				RepositoryCreated,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.CreatedAt),
			)
		case "edited":
			message = "Repository edited" // unimplemented for now
		case "deleted":
			message = fmt.Sprintf(
				RepositoryDeleted,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		case "privatized":
			message = fmt.Sprintf(
				RepositoryPrivatized,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		case "publicized":
			message = fmt.Sprintf(
				RepositoryPublicized,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		case "renamed":
			message = fmt.Sprintf(
				RepositoryRenamed,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Changes.Repository.Name.From),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		case "unarchived":
			message = fmt.Sprintf(
				RepositoryUnarchived,
				escapeText(repositoryEvent.Sender.Login),
				escapeURL(repositoryEvent.Sender.HTMLURL),
				escapeText(repositoryEvent.Repository.FullName),
				formatRFC3339Timestamp(repositoryEvent.Repository.UpdatedAt),
			)
		default:
			message = ""
	}

	return message
}