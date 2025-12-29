package services

import (
	"context"

	"github.com/mik-sea/go-firebase-backend/internal/config"
)

type Transaction struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	CreatedAt int64   `json:"created_at"`
}
type TransactionListResponse struct {
	Count int           `json:"count"`
	Data  []Transaction `json:"data"`
}

func GetAllTransactions() (*TransactionListResponse, error) {
	ctx := context.Background()

	client, err := config.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iter := client.Collection("transactions").Documents(ctx)

	var results []Transaction

	for {
		doc, err := iter.Next()
		if err != nil {
			break // selesai
		}

		var trx Transaction
		doc.DataTo(&trx)
		trx.ID = doc.Ref.ID // ambil document ID

		results = append(results, trx)
	}

	return &TransactionListResponse{
		Count: len(results),
		Data:  results,
	}, nil
}
