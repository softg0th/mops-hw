package service

import (
	"context"
	entities "iotController/entities"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) InsertPostsMongoStream(doc entities.Document, success chan<- bool) {
	collection := s.db.Collection

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"operationType", bson.D{
				{"$in", bson.A{"insert"}},
			}},
		}}},
	}

	streamOptions := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	changeStream, err := collection.Watch(context.Background(), pipeline, streamOptions)
	if err != nil {
		log.Fatalf("Error opening change stream: %v", err)
	}
	defer changeStream.Close(context.Background())

	for changeStream.Next(context.Background()) {
		var event struct {
			FullDocument entities.Document `bson:"fullDocument"`
		}
		if err := changeStream.Decode(&event); err != nil {
			log.Printf("Error decoding change stream document: %v", err)
			success <- false
			continue
		}

		_, err := collection.InsertOne(context.Background(), doc)
		if err != nil {
			log.Printf("Error inserting document: %v", err)
			success <- false
			continue

		}
	}
	success <- true
}
