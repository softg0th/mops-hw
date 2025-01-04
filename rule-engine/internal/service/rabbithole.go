package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"ruleEngine/internal/core"
	"time"
)

type Message struct {
	DeviceID        int       `json:"device_id"`
	Timestamp       time.Time `json:"timestamp"`
	SomeUsefulField int       `json:"some_useful_field"`
}

func (s *Service) ReadFromRabbitMQ(targetValue int, durationValue int) {
	var msgs <-chan amqp.Delivery
	var err error

	retry := func() <-chan amqp.Delivery {
		for attempts := 1; ; attempts++ {
			msgs, err = s.Rabbit.Channel.Consume(
				s.Rabbit.MainQueue.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err == nil {
				return msgs
			}
			s.Logger.Error(map[string]interface{}{
				"message": err,
				"attempt": attempts,
				"error":   true,
			})
			time.Sleep(time.Second * time.Duration(attempts))
		}
	}
	msgs = retry()

	observedDevices := core.NewObservedDevices(targetValue, durationValue)

	instantChannel := make(chan bool)
	durationChannel := make(chan bool)

	go func() {
		for d := range msgs {
			var message Message
			if err := json.Unmarshal(d.Body, &message); err != nil {
				s.Logger.Error(map[string]interface{}{
					"message": "Failed to unmarshal JSON",
					"error":   err.Error(),
				})
				continue
			}

			s.Logger.Info(map[string]interface{}{
				"message":   "Message received successfully",
				"device_id": message.DeviceID,
				"timestamp": message.Timestamp.Format(time.RFC3339),
				"value":     message.SomeUsefulField,
			})

			isMatchInstant, isMatchWithDuration := observedDevices.ProcessMessage(message.DeviceID, message.SomeUsefulField)

			instantChannel <- isMatchInstant
			durationChannel <- isMatchWithDuration
		}
	}()

	go func() {
		for {
			select {
			case instant := <-instantChannel:
				if instant {
					s.Logger.Info(map[string]interface{}{
						"message": "Instant match detected",
					})
				}
			case duration := <-durationChannel:
				if duration {
					s.Logger.Info(map[string]interface{}{
						"message": "Duration match detected",
					})
				}
			}
		}
	}()

	select {}
}
