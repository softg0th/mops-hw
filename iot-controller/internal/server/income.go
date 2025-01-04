package server

import (
	"context"
	"io"
	"iotController/internal/entities"
	"iotController/internal/proto"
	"time"
)

func (s *Server) StreamWithAck(stream iot_controller.IotService_StreamWithAckServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				s.Service.Logger.Error(map[string]interface{}{
					"message": "End of stream",
					"error":   true,
				})
				return nil
			}
			s.Service.Logger.Error(map[string]interface{}{
				"message": err,
				"error":   true,
			})
			return err
		}
		if in == nil {
			s.Service.Logger.Error(map[string]interface{}{
				"message": "Received nil input",
				"error":   true,
			})
			return nil
		}

		s.Service.Logger.Info(map[string]interface{}{
			"message":           "Document received",
			"device_id":         in.DeviceId,
			"timestamp":         in.Timestamp.AsTime().Format(time.RFC3339),
			"some_useful_field": in.SomeUsefulField,
		})

		newEntity := entities.NewDocument(in.DeviceId, in.Timestamp.AsTime(), in.SomeUsefulField)

		ctx := context.Background()
		successDB := make(chan bool)

		go func() {
			defer close(successDB)
			s.Service.InsertPostsMongoStream(ctx, *newEntity, successDB)
		}()

		result, ok := <-successDB
		if !ok {
			s.Service.Logger.Error(map[string]interface{}{
				"message": "Failed to insert document",
				"error":   true,
			})
			return nil
		}

		successRabbitMQ := make(chan bool)

		go func() {
			defer close(successRabbitMQ)
			s.Service.PublishToRabbitMQ(newEntity.DeviceID, newEntity.Timestamp,
				newEntity.SomeUsefulField, successRabbitMQ)
		}()

		err = stream.Send(&iot_controller.PackageResponse{Success: result})
		if err != nil {
			s.Service.Logger.Error(map[string]interface{}{
				"message": err,
				"error":   true,
			})
			return err
		}
	}
}
