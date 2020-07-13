package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDB returns db instance
func NewDB() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb://zero:zeropassword@localhost:27017",
	))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client, ctx
}
