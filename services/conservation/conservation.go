package main

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	pb "github.com/yulpa/yulmails/services/conservation/v1beta1"
)

type ConservationService struct{ DaoService *Dao }

// ListConservation returns list of conservation law
func (c *ConservationService) ListConservation(in *empty.Empty, stream pb.ConservationService_ListConservationServer) error {
	conservations, err := c.DaoService.GetConservations()
	if err != nil {
		return err
	}
	for _, conservation := range conservations {
		if err := stream.Send(&conservation); err != nil {
			return err
		}
	}
	return nil
}

// CreateConservation add a conservation into the DB
func (c *ConservationService) CreateConservation(ctx context.Context, in *pb.Conservation) (*pb.Conservation, error) {
	if err := c.DaoService.CreateConservation(in); err != nil {
		return in, err
	}
	return in, nil
}
