package server

import (
	pb "iotController/proto"
	srv "iotController/service"
)

type Server struct {
	pb.UnimplementedIotServiceServer
	Service *srv.Service
}

func NewServer(s *srv.Service) *Server {
	return &Server{Service: s}
}
