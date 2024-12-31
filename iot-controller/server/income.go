package server

import (
	"iotController/entities"
	pb "iotController/proto"
)

func (h *Handler) StreamWithAck(stream pb.IotService_StreamWithAckServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}

		newEntity := entities.NewDocument(in.DeviceId, in.Timestamp.AsTime(), in.SomeUsefulField)
		success := make(chan bool)

		go h.service.InsertPostsMongoStream(*newEntity, success)

		result := <-success

		err = stream.Send(&pb.PackageResponse{Success: result})
		if err != nil {
			return err
		}
	}
}
