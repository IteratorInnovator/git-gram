package db

import (
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

func RegisterChat(ctx context.Context, client *firestore.Client, chat_id int64) error {
	data := make(map[string]interface{})

	data["chat_id"] = chat_id
	data["installation_id"] = nil
	data["mute"] = false

	_, _, err := client.Collection("user").Add(ctx, data)

	return err
}

