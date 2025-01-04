package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"time"
)

type Message struct {
	DeviceID       int       `json:"device_id"`
	Timestamp      time.Time `json:"timestamp"`
	SomeUsefulFiel int       `json:"some_useful_field"`
}

func (s *Service) PublishToRabbitMQ(deviceID int, timestamp time.Time, someUsefulField int,
	success chan<- bool) {

	publishData := Message{
		DeviceID:       deviceID,
		Timestamp:      timestamp,
		SomeUsefulFiel: someUsefulField,
	}

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
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

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
