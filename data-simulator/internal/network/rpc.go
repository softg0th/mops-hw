package network

import (
	"context"
	"data-simulator/internal/enteties"
	"data-simulator/internal/exceptions"
	pb "data-simulator/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type RPCConn struct {
	stream pb.IotService_StreamWithAckClient
}

func NewRPCConn() *RPCConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Netrwork connection", Message: "could not connect to" +
			" grpc server"})
	}

	client := pb.NewIotServiceClient(conn)
	stream, err := client.StreamWithAck(context.Background())
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Streaming", Message: "start streaming failed"})
	}

	return &RPCConn{
		stream: stream,
	}
}

func (r RPCConn) StreamRequest(message *enteties.Message) {
	err := r.stream.Send(&pb.GetPackageRequest{
		DeviceId:        int32(message.DeviceID),
		Timestamp:       timestamppb.New(message.Timestamp),
		SomeUsefulField: int32(message.SomeUsefulField),
	})
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Stream message", Message: "could not send message"})
	}

	resp, err := r.stream.Recv()
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Server response", Message: "server not responding"})
	}
	log.Printf("Received response: %v", resp.Success)
}
