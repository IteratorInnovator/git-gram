package db

import (
	"context"
	"strconv"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/IteratorInnovator/git-gram/config"
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
	data["installation_id"] = nil
	data["mute"] = false

	_, err := client.Collection("users").Doc(docId).Create(ctx, data)

	return err
}

func SaveInstallation(ctx context.Context, client *firestore.Client, chat_id int64, installation_id int64) error {
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
    	{ Path: "mute", Value: true },
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
    	{ Path: "mute", Value: false },
	})
	if (err != nil) {
		return true, errors.New("failed to unmute")
	}
	return true, nil
}

