package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IteratorInnovator/git-gram/config"
	"github.com/IteratorInnovator/git-gram/github/message_templates"
	"github.com/gofiber/fiber/v2"
)

func HandleGitHubWebhookEvent(event string, chatId int64, ctx *fiber.Ctx) error {
	var err error = nil

	switch (event) {
		case "push":
			fmt.Printf("github event handler: push event for chat id=%d\n", chatId)
			err = handlePushEvent(ctx, chatId)
		case "create":
			fmt.Printf("github event handler: create event for chat id=%d\n", chatId)
			fmt.Println("create event")
		case "delete":
			fmt.Printf("github event handler: delete event for chat id=%d\n", chatId)
			fmt.Println("delete event")
		case "pull_request":
			fmt.Printf("github event handler: pull_request event for chat id=%d\n", chatId)
			fmt.Println("pull request event")
		default:
			fmt.Printf("github event handler: unknown event=%v chat id=%d\n", event, chatId)
			fmt.Printf("%v event", event)
	}
	return err
}

func handlePushEvent(ctx *fiber.Ctx, chatId int64) error {	
	fmt.Printf("push event: building notification for chat id=%d\n", chatId)
	url := fmt.Sprintf("%v/%v", config.TelegramCfg.TELEGRAM_BOT_API_ENDPOINT, "sendMessage")

	var pushEvent PushEvent
	err := ctx.BodyParser(&pushEvent)
	if err != nil {
		fmt.Printf("push event: body parse failed: %v\n", err)
		return err
	}
	fmt.Printf("push event: repo=%s ref=%s commits=%d compare=%s\n", pushEvent.Repository.FullName, pushEvent.Ref, len(pushEvent.Commits), pushEvent.Compare)

	keyboardButtons := [][]InlineKeyboardButton {
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
			InlineKeyboardButton{
				Text: "Branch",
				URL: pushEvent.Repository.HTMLURL + "/tree/" + formatRef(pushEvent.Ref),
			},
		},
	}

	var messageText string = ""
	var commitCount int = len(pushEvent.Commits)
	if commitCount > 1 {
		fmt.Println("push event: using multiple commits template")
		messageText = fmt.Sprintf(
			message_templates.MultipleCommitsPush,
			pushEvent.Repository.FullName,
			pushEvent.Sender.Login,
			pushEvent.Sender.HTMLURL,
			commitCount,
			formatRef(pushEvent.Ref),
			formatUnixTimestamp(pushEvent.Repository.PushedAt),
			shortenSHA(pushEvent.HeadCommit.ID),
			pushEvent.HeadCommit.Message,
		)
	} else {
		fmt.Println("push event: using single commit template")
		messageText = fmt.Sprintf(
			message_templates.SingleCommitPush,
			pushEvent.Repository.FullName,
			pushEvent.Sender.Login,
			pushEvent.Sender.HTMLURL,
			formatRef(pushEvent.Ref),
			formatUnixTimestamp(pushEvent.Repository.PushedAt),
			shortenSHA(pushEvent.HeadCommit.ID),
			pushEvent.HeadCommit.Message,
		)
	}

	payload := struct {
		ChatID      int                  `json:"chat_id"`
		ParseMode   string               `json:"parse_mode"`
		Text        string               `json:"text"`
		ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
	} {
		ChatID: int(chatId),
		ParseMode: "MarkdownV2",
		Text: messageText,
		ReplyMarkup: InlineKeyboardMarkup{
			InlineKeyboard: keyboardButtons,
		},
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("push event: payload marshal failed: %v\n", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		fmt.Printf("push event: telegram send failed: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("push event: telegram response status=%s\n", resp.Status)
	return nil
}
