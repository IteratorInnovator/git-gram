package github

import (
	"fmt"
)

func HandleGitHubWebhookEvent(event string) error {
	switch (event) {
		case "push":
			fmt.Println("push event")
		case "create":
			fmt.Println("create event")
		case "delete":
			fmt.Println("delete event")
		case "pull_request":
			fmt.Println("pull request event")
		default:
			fmt.Printf("%v event", event)
	}
	return nil
}