package server

import (
	"context"
	"io"
	"iotController/entities"
	pb "iotController/proto"
	"log"
)

func (s *Server) StreamWithAck(stream pb.IotService_StreamWithAckServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed the stream")
				return nil
			}
		}
		log.Printf("Received message: DeviceId=%d, Timestamp=%s, SomeUsefulField=%d",
			in.DeviceId, in.Timestamp.AsTime(), in.SomeUsefulField)

		newEntity := entities.NewDocument(in.DeviceId, in.Timestamp.AsTime(), in.SomeUsefulField)
		ctx := context.Background()
		success := make(chan bool)
		go s.Service.InsertPostsMongoStream(ctx, *newEntity, success)

		result := <-success
		log.Printf("success: %v", result)
		err = stream.Send(&pb.PackageResponse{Success: result})
		if err != nil {
			return err
		}
	}
}
