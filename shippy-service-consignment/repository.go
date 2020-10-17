package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	consignmentpb "github.com/days85/shippy/shippy-service-consignment/proto/consignment"
)

// Consignment -
type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

// Container -
type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

// Containers -
type Containers []*Container

// MarshalConsignment -
func MarshalConsignment(consignment *consignmentpb.Consignment) *Consignment {
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  MarshalContainerCollection(consignment.Containers),
		VesselID:    consignment.VesselId,
	}
}

// UnmarshalConsignment -
func UnmarshalConsignment(consignment *Consignment) *consignmentpb.Consignment {
	return &consignmentpb.Consignment{
		Id:          consignment.ID,
		Description: consignment.Description,
		Weight:      consignment.Weight,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

// UnmarshalConsignmentCollection -
func UnmarshalConsignmentCollection(consignments []*Consignment) []*consignmentpb.Consignment {
	collection := make([]*consignmentpb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}

// MarshalContainer -
func MarshalContainer(container *consignmentpb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

// UnmarshalContainer -
func UnmarshalContainer(container *Container) *consignmentpb.Container {
	return &consignmentpb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

// MarshalContainerCollection -
func MarshalContainerCollection(containers []*consignmentpb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

// UnmarshalContainerCollection -
func UnmarshalContainerCollection(containers []*Container) []*consignmentpb.Container {
	collection := make([]*consignmentpb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

// MongoRepository -
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (r *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := r.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll -
func (r *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := r.collection.Find(ctx, nil, nil)
	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
