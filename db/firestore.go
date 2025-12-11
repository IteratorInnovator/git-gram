package db

import (
	"strconv"
	"context"
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

