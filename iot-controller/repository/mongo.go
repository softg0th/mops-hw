package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoConnection struct {
	client *mongo.Client
}

type DataBase struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func New() *mongoConnection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return &mongoConnection{
		client: client,
	}
}

func InitDatabase(mongo *mongoConnection) (*DataBase, error) {
	datasDB := mongo.client.Database("iot")
	collection := datasDB.Collection("messages")

	return &DataBase{
		db:         datasDB,
		collection: collection,
	}, nil
}

func GetDBInstance() *DataBase {
	db := &DataBase{}
	return db
}
