// shippy-vessel-service/handler.go
package main

import (
	"context"

	pb "github.com/jipram017/go-ship/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

// FindAvailable vessels
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	res.Vessel = UnmarshalVessel(vessel)

	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}

	res.Vessel = req
	return nil
}
