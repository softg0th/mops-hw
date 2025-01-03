package server

import (
	pb "iotController/internal/proto"
	srv "iotController/internal/service"
)

type Server struct {
	pb.UnimplementedIotServiceServer
	Service *srv.Service
}

func NewServer(s *srv.Service) *Server {
	return &Server{Service: s}
}
