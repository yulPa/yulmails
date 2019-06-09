package main

import (
	empty "github.com/golang/protobuf/ptypes/empty"

	pb "gitlab.com/tortuemat/yulmails/services/conservation/v1beta1"
)

type ConservationService struct{}

func (c *ConservationService) ListConservation(in *empty.Empty, stream pb.ConservationService_ListConservationServer) error {
	conservations := []pb.Conservation{
		pb.Conservation{
			ID:               1,
			Created:          "2019-01-25 12:13:14",
			Sent:             10,
			Unsent:           100,
			KeepEmailContent: false,
		},
		pb.Conservation{
			ID:               2,
			Created:          "2020-01-25 12:13:14",
			Sent:             100,
			Unsent:           10,
			KeepEmailContent: true,
		},
	}
	for _, conservation := range conservations {
		if err := stream.Send(&conservation); err != nil {
			return err
		}
	}
	return nil
}
