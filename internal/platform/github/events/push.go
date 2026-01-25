package events

import (
	"fmt"
)

const singleCommitPush string = `🔔 *New Push to %s*

[%s](%s) pushed 1 commit to ` + "`%s`" + ` at %s

Commit: ` + "`%s`" + `
%s`

const multipleCommitsPush string = `🔔 *New Push to %s*

[%s](%s) pushed %d commits to ` + "`%s`" + ` at %s

Latest Commit: ` + "`%s`" + `
%s`

// BuildPushInlineKeyboard builds the inline keyboard for a push event.
func BuildPushInlineKeyboard(event *PushEvent) [][]InlineKeyboardButton {
	return [][]InlineKeyboardButton{
		{
			InlineKeyboardButton{
				Text: "View Commit",
				URL:  event.HeadCommit.URL,
			},
			InlineKeyboardButton{
				Text: "Changes",
				URL:  event.Compare,
			},
		},
		{
			InlineKeyboardButton{
				Text: "Repository",
				URL:  event.Repository.HTMLURL,
			},
			InlineKeyboardButton{
				Text: "Branch",
				URL:  event.Repository.HTMLURL + "/tree/" + FormatRef(event.Ref),
			},
		},
	}
}

// BuildPushMessage builds the message for a push event.
func BuildPushMessage(event *PushEvent) string {
	commitCount := len(event.Commits)

	if commitCount > 1 {
		return fmt.Sprintf(
			multipleCommitsPush,
			EscapeText(event.Repository.FullName),
			EscapeText(event.Sender.Login),
			EscapeURL(event.Sender.HTMLURL),
			commitCount,
			FormatRef(event.Ref),
			FormatUnixTimestamp(event.Repository.PushedAt),
			ShortenSHA(event.HeadCommit.ID),
			EscapeText(event.HeadCommit.Message),
		)
	}
	return fmt.Sprintf(
		singleCommitPush,
		EscapeText(event.Repository.FullName),
		EscapeText(event.Sender.Login),
		EscapeURL(event.Sender.HTMLURL),
		FormatRef(event.Ref),
		FormatUnixTimestamp(event.Repository.PushedAt),
		ShortenSHA(event.HeadCommit.ID),
		EscapeText(event.HeadCommit.Message),
	)
}
