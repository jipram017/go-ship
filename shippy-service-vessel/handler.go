// shippy-vessel-service/handler.go
package main

import (
	"context"
	"log"

	pb "github.com/jipram017/go-ship/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

// FindAvailable vessels
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	log.Println("Req:")
	log.Println(req)
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	test1 := UnmarshalVessel(vessel)
	test2 := vessel

	log.Println("This is to test unmarshalling struct : ")
	log.Println(test1)

	log.Println("This is to test original struct : ")
	log.Println(test2)

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
