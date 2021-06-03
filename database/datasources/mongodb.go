package datasources

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Context context.Context
	MongoDB *mongo.Client
}

type IMongoDB interface {
	Disconnect()
}

func NewMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	option := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(ctx, option)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDB{
		Context: ctx,
		MongoDB: client,
	}
}

func (db MongoDB) Disconnect() {
	db.MongoDB.Disconnect(db.Context)
}
