package db

import (
	"context"
	"errors"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
	"google.golang.org/api/iterator"
)

func CreateClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClientWithDatabase(ctx, config.FirestoreCfg.GOOGLE_CLOUD_PROJECT_ID, config.FirestoreCfg.FIRESTORE_DATABASE_ID)

	if (err != nil) {
		return nil, err
	}
	return client, nil
}

func SaveChat(ctx context.Context, client *firestore.Client, chat_id int64) error {
	var docId string = strconv.FormatInt(chat_id, 10)

	data := make(map[string]interface{})

	data["chat_id"] = chat_id
	data["installation_id"] = 0
	data["github_account_username"] = ""
	data["muted"] = false

	_, err := client.Collection("users").Doc(docId).Create(ctx, data)

	return err
}

func UpdateInstallation(ctx context.Context, client *firestore.Client, chat_id int64, installation_id int64, account_username string) error {
	var chatId string = strconv.FormatInt(chat_id, 10)

	docRef := client.Collection("users").Doc(chatId)
	snap, _ := docRef.Get(ctx)

	if !snap.Exists() {
		return errors.New("doc does not exist")
	}

	_, err := docRef.Update(ctx, []firestore.Update {
    	{ Path: "installation_id", Value: installation_id },
		{ Path: "github_account_username", Value: account_username },
	})
	if (err != nil) {
		return errors.New("failed to save installation")
	}
	return nil
}

func Mute(ctx context.Context, client *firestore.Client, chat_id int64) (bool, error) {
	var chatId string = strconv.FormatInt(chat_id, 10)

	docRef := client.Collection("users").Doc(chatId)
	snap, _ := docRef.Get(ctx)

	if !snap.Exists() {
		return false, nil
	}

	_, err := docRef.Update(ctx, []firestore.Update{
    	{ Path: "muted", Value: true },
	})
	if (err != nil) {
		return true, errors.New("failed to mute")
	}
	return true, nil
}

func Unmute(ctx context.Context, client *firestore.Client, chat_id int64) (bool, error) {
	var chatId string = strconv.FormatInt(chat_id, 10)

	docRef := client.Collection("users").Doc(chatId)
	snap, _ := docRef.Get(ctx)

	if !snap.Exists() {
		return false, nil
	}

	_, err := docRef.Update(ctx, []firestore.Update{
    	{ Path: "muted", Value: false },
	})
	if (err != nil) {
		return true, errors.New("failed to unmute")
	}
	return true, nil
}

func FetchUserInfo(ctx context.Context, client *firestore.Client, chat_id int64) (string, bool, error) {
	var chatId string = strconv.FormatInt(chat_id, 10)

	docRef := client.Collection("users").Doc(chatId)
	snap, _ := docRef.Get(ctx)

	if !snap.Exists() {
		return "", false, errors.New("user doc does not exist")
	}

	var u struct {
		AccountLogin string `firestore:"github_account_username"`
		Muted        bool   `firestore:"muted"`
	}
	err := snap.DataTo(&u)
	if err != nil {
		return "", false, err
	}

	return u.AccountLogin, u.Muted, nil
}

func FetchChatIdAndMute(ctx context.Context, client *firestore.Client, installation_id int64) (int64, bool, error) {
	query := client.Collection("users").Select("installation_id").Where("installation_id", "=", installation_id)
	documentIterator := query.Documents(ctx)

	snap, err := documentIterator.Next()
	if err != nil || err == iterator.Done {
		return 0, false, err
	}

	var data struct {
		ChatId int64 `firestore:"chat_id"`
		Muted  bool  `firestore:"muted"` 
	}
	err = snap.DataTo(&data)
	if err != nil {
		return 0, false, err
	}

	return data.ChatId, data.Muted, nil
}

func FetchInstallationId(ctx context.Context, client *firestore.Client, chat_id int64) (int64, error) {
	var chatId string = strconv.FormatInt(chat_id, 10)

	docRef := client.Collection("users").Doc(chatId)
	snap, _ := docRef.Get(ctx)

	if !snap.Exists() {
		return 0, errors.New("user doc does not exist")
	}

	var data struct {
		InstallationId int64 `firestore:"installation_id"`
	}
	err := snap.DataTo(&data)
	if err != nil {
		return 0, err
	}

	return data.InstallationId, nil
}