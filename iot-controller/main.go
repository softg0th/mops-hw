package main

import (
	"iotController/repository"
	"iotController/server"
	"iotController/service"
	"log"
	"net"

	pb "iotController/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := repository.NewMongoConnection("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	db := repository.NewDataBase(conn, "iot", "messages")

	iotService := service.NewService(db)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listen.Close()

	serv := grpc.NewServer()
	iotServer := server.NewServer(iotService)
	pb.RegisterIotServiceServer(serv, iotServer)

	log.Printf("server listening at %v", listen.Addr())

	if err := serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
