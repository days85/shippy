package main

import (
	"context"
	"errors"

	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"

	consignmentpb "github.com/days85/shippy/shippy-service-consignment/proto/consignment"
)

type handler struct {
	repository
	vesselClient vesselpb.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (h *handler) CreateConsignment(ctx context.Context, req *consignmentpb.Consignment, res *consignmentpb.Response) error {
	// Here we call a client instance of our vesselpb service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := h.vesselClient.FindAvailable(ctx, &vesselpb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vesselpb, returned nil")
	}

	if err != nil {
		return err
	}

	// We set the VesselId as the vesselpb we got back from our
	// vesselpb service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	if err = h.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (h *handler) GetConsignments(ctx context.Context, req *consignmentpb.GetRequest, res *consignmentpb.Response) error {
	consignments, err := h.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
