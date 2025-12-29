package services

import (
	"context"

	"github.com/mik-sea/go-firebase-backend/internal/config"
)

type User struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func SaveUser(user User) error {
	ctx := context.Background()
	client, err := config.App.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.UID).Set(ctx, user)
	return err
}
