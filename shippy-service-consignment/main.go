// shippy/shippy-service-consignment/main.go
package main

import (
	"context"
	"log"
	"os"

	pb "github.com/jipram017/go-ship/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/jipram017/go-ship/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://127.0.0.1:27017"
)

func main() {

	service := micro.NewService(micro.Name("shippy.service.consignment"))
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())

	h := &handler{repository, vesselClient}

	log.Println("successfully connected to mongodb")

	// Register our implementation with
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// vesselResponse, err := vesselClient.Create(context.Context(), &vesselProto.Vessel{
	// 	MaxWeight: 500000,
	// 	Capacity:  500,
	// })
	// if vesselResponse == nil {
	// 	log.Panic("error fetching vessel, returned nil")
	// }

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
