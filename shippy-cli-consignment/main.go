package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/jipram017/go-ship/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {

	// Set up a connection to the server.
	service := micro.NewService(micro.Name("shippy.cli.consignment"))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse error: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{"token": token})

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Second call
	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
