package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	consignmentpb "github.com/days85/shippy/shippy-service-consignment/proto/consignment"
	vesselpb "github.com/days85/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	address         = "localhost:50051"
	consignmentFile = "consignment.json"
	vesselFile      = "vessel.json"
)

func parseConsignement(file string) (*consignmentpb.Consignment, error) {
	var consignment *consignmentpb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(data, &consignment)
	return consignment, err
}

func parseVessel(file string) (*vesselpb.Vessel, error) {
	var vessel *vesselpb.Vessel
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(data, &vessel)
	return vessel, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.cli.consignment"))
	service.Init()

	consignmentClient := consignmentpb.NewShippingService("shippy.service.consignment", service.Client())
	vesselClient := vesselpb.NewVesselService("shippy.service.vessel", service.Client())

	cFile := consignmentFile
	if len(os.Args) > 1 {
		cFile = os.Args[1]
	}
	consignment, err := parseConsignement(cFile)
	if err != nil {
		log.Fatalf("Could not parse consignment file: %v", err)
	}

	vFile := consignmentFile
	if len(os.Args) > 1 {
		vFile = os.Args[1]
	}
	vessel, err := parseVessel(vFile)
	if err != nil {
		log.Fatalf("Could not parse vessel file: %v", err)
	}

	vesselResponse, err := vesselClient.Create(context.Background(), vessel)
	if err != nil {
		log.Fatalf("Could not create a vessel: %v", err)
	}
	log.Printf("Created: %t", vesselResponse.Created)

	consignmentResponse, err := consignmentClient.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not create a consignment: %v", err)
	}
	log.Printf("Created: %t", consignmentResponse.Created)

	getAll, err := consignmentClient.GetConsignments(context.Background(), &consignmentpb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
