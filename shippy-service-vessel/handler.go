package main

import (
	"context"

	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

// Create - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (h *handler) Create(ctx context.Context, req *vesselpb.Vessel, res *vesselpb.Response) error {
	// Save our vessel
	if err := h.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Created = true
	res.Vessel = req
	return nil
}

// FindAvailable vessels
func (h *handler) FindAvailable(ctx context.Context, req *vesselpb.Specification, res *vesselpb.Response) error {

	// Find the next available vessel
	vessel, err := h.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = UnmarshalVessel(vessel)
	return nil
}
