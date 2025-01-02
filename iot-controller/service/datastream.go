package service

import (
	"context"
	entities "iotController/entities"
	"log"
)

func (s *Service) InsertPostsMongoStream(ctx context.Context, doc entities.Document, success chan<- bool) {
	collection := s.db.Collection
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Error inserting document: %v", err)
		success <- false
		return
	}
	log.Printf("Document inserted successfully")
	success <- true
}
