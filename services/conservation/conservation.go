package main

import (
	empty "github.com/golang/protobuf/ptypes/empty"

	pb "gitlab.com/tortuemat/yulmails/services/conservation/v1beta1"
)

type ConservationService struct{ DaoService *Dao}

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
