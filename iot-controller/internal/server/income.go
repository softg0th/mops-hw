package server

import (
	"context"
	"io"
	"iotController/internal/entities"
	"iotController/internal/proto"
	"log"
	"time"
)

func (s *Server) StreamWithAck(stream iot_controller.IotService_StreamWithAckServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				s.Service.Logger.Error(map[string]interface{}{
					"message": err,
					"error":   true,
				})
				return nil
			}
		}
		log.Printf("exampllllee")
		s.Service.Logger.Info(map[string]interface{}{
			"message":           "Document inserted successfully",
			"device_id":         in.DeviceId,
			"timestamp":         in.Timestamp.AsTime().Format(time.RFC3339),
			"some_useful_field": in.SomeUsefulField,
		})

		newEntity := entities.NewDocument(in.DeviceId, in.Timestamp.AsTime(), in.SomeUsefulField)
		ctx := context.Background()
		success := make(chan bool)
		go s.Service.InsertPostsMongoStream(ctx, *newEntity, success)

		result := <-success

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
