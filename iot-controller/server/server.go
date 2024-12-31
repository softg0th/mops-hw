package server

import (
	pb "iotController/proto"
	"iotController/service"
)

type Server struct {
	pb.UnimplementedIotServiceServer
	Handler *Handler
}

func NewServer(s *service.Service) *Server {
	handler := NewHandler(s)
	return &Server{Handler: handler}
}
