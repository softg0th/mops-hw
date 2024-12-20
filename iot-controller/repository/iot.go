package repository

import (
	"context"
	entities "iotController/entities"
)

func InsertPost(mongo *DataBase, d entities.Document) error {
	ctx := context.Background()
	_, err := mongo.collection.InsertOne(ctx, d)

	return err
}
