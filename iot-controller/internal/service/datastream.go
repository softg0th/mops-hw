package service

import (
	"context"
	"iotController/internal/entities"
)

func (s *Service) InsertPostsMongoStream(ctx context.Context, doc entities.Document, success chan<- bool) {
	collection := s.db.Collection
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		s.Logger.Error(map[string]interface{}{
			"message": err,
			"error":   true,
		})
		success <- false
		return
	}
	s.Logger.Info(map[string]interface{}{"message": "Document inserted successfully"})
	success <- true
}
