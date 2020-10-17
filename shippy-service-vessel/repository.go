package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"
)

// Vessel -
type Vessel struct {
	ID        string
	Capacity  int32
	MaxWeight int32
	Name      string
	Available bool
	OwnerID   string
}

// Specification -
type Specification struct {
	Capacity  int32
	MaxWeight int32
}

// MarshalVessel -
func MarshalVessel(vessel *vesselpb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

// UnmarshalVessel -
func UnmarshalVessel(vessel *Vessel) *vesselpb.Vessel {
	return &vesselpb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

// MarshalSpecification -
func MarshalSpecification(spec *vesselpb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// UnarshalSpecification -
func UnarshalSpecification(spec *Specification) *vesselpb.Specification {
	return &vesselpb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

type repository interface {
	Create(ctx context.Context, vessel *Vessel) error
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
}

// MongoRepository -
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (r *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := r.collection.InsertOne(ctx, vessel)
	return err
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (r *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{
			{
				"$lte",
				spec.Capacity,
			},
			{
				"$lte",
				spec.MaxWeight,
			},
		},
	}}
	vessel := &Vessel{}
	if err := r.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}
