package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/v2"

	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultDBHost = "datastore:27017"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.vessel"),
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

	vesselCollection := client.Database("shippy").Collection("vessel")

	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	// Register handlers
	_ = vesselpb.RegisterVesselServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
