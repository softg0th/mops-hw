package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

func (s *Service) PublishToRabbitMQ(publishData map[int]int, success chan<- bool) {
	body, err := json.Marshal(publishData)
	if err != nil {
		s.Logger.Error(map[string]interface{}{
			"message": err,
			"error":   true,
		})
		success <- false
		return
	}
	
	err = s.Rabbit.Channel.Publish(
		"",
		s.Rabbit.MainQueue.Name,
		false,
		false,
		amqp.Publishing{ContentType: "application/json",
			Body: body})

	if err != nil {
		s.Logger.Error(map[string]interface{}{
			"message": err,
			"error":   true,
		})
		success <- false
		return
	}
	s.Logger.Info(map[string]interface{}{"message": "Message published successfully"})
	success <- true
}
