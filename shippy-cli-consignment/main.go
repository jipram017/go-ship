package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"

	pb "github.com/jipram017/go-ship/shippy-service-consignment/proto/consignment"

	"github.com/micro/go-micro/metadata"
	microclient "github.com/micro/go-micro/v2/client"
)

const (
	defaultFilename = "consignment.json"
	defaultToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImVtYWlsIjoiamlwcmFtMDE3QGdtYWlsLmNvbSIsInBhc3N3b3JkIjoidG9ob2t1MjAxMyJ9LCJleHAiOjE1MDAwLCJpc3MiOiJzaGlwcHkuc2VydmljZS51c2VyIn0.uINU2NZ_BBlPJeY0ExMvXgrOpbHtuJ3ubNJ3vIi"
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
	// Contact the server and print out its response.
	file := defaultFilename
	token := defaultToken

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our consignment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	client := pb.NewShippingService("go.micro.srv.consignment", microclient.DefaultClient)

	// if len(os.Args) < 3 {
	// 	log.Fatal(errors.New("Not enough arguments, expecing file and token."))
	// }

	// file = os.Args[1]
	// token = os.Args[2]

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse error: %v", err)
	}

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
