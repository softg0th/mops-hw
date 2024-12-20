package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "iotController/proto"

	entities "iotController/entities"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listen.Close()

	serv := grpc.NewServer()
	pb.RegisterIotServiceServer(serv, &entities.Server{})

	log.Printf("server listening at %v", listen.Addr())

	if err := serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
