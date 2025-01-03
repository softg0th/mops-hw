package entities

import (
	pb "iotController/internal/proto"
)

type Server struct {
	pb.UnimplementedIotServiceServer
}
