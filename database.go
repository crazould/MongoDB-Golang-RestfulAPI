package main

import (
	"context"
	f "fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DbServerAddress ...
const DbServerAddress = "127.0.0.1"

//DbServerPort ...
const DbServerPort = "27017"

//Database ...
const Database = "mongodb"

//DbHandler is
type DbHandler struct {
	connectionString string
}

func (db *DbHandler) connectWithMongoDB() *mongo.Client {

	db.connectionString = f.Sprintf("%s://%s:%s",
		Database,
		DbServerAddress,
		DbServerPort,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(db.connectionString))

	return client
}
