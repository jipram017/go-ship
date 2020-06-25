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

	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	// Initialize service
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

	consignmentCollection := client.Database("shippy").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())

	h := &handler{repository, vesselClient}

	log.Println("successfully connected to mongodb")

	// Register our implementation with
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
