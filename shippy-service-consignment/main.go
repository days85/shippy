package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"

	consignmentpb "github.com/days85/shippy/shippy-service-consignment/proto/consignment"
)

const (
	defaultDBHost = "datastore:27017"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags.
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultDBHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselpb.NewVesselService("shippy.service.vessel", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	_ = consignmentpb.RegisterShippingServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
