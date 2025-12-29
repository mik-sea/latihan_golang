package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	ctx := context.Background()

	credPath := os.Getenv("FIREBASE_CREDENTIALS")
	if credPath == "" {
		credPath = "configs/serviceAccountKey.json"
	}

	opt := option.WithCredentialsFile(credPath)

	conf := &firebase.Config{
		ProjectID: "finance-firebase-652de",
	}

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("firebase init error: %v", err)
	}

	App = app
}
