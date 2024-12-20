package handler

import (
	"context"
	pb "iotController/proto"
	"sync"
)

func (h *Handler) ReadMessage(ctx context.Context, in *pb.GetPackageRequest) (*pb.PackageResponse, error) {
	var wg sync.WaitGroup

	for {
		wg.Add(1)

	}
	wg.Wait()
	
}
