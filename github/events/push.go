package events

import (
	"fmt"
)

const SingleCommitPush string = `🔔 *New Push to %s*

[%s](%s) pushed 1 commit to ` + "`%s`" + ` at %s

Commit: ` + "`%s`" + `
%s`


const MultipleCommitsPush string = `🔔 *New Push to %s*

[%s](%s) pushed %d commits to ` + "`%s`" + ` at %s

Latest Commit: ` + "`%s`" + `
%s`


func BuildPushInlineKeyboard(pushEvent *PushEvent) [][]InlineKeyboardButton {
	return [][] InlineKeyboardButton {
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
			MultipleCommitsPush,
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
		SingleCommitPush,
		escapeText(pushEvent.Repository.FullName),
		escapeText(pushEvent.Sender.Login),
		escapeURL(pushEvent.Sender.HTMLURL),
		formatRef(pushEvent.Ref),
		formatUnixTimestamp(pushEvent.Repository.PushedAt),
		shortenSHA(pushEvent.HeadCommit.ID),
		escapeText(pushEvent.HeadCommit.Message),
	)
}
