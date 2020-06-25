// shippy/shippy-service-consignment/main.go
package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/jipram017/go-ship/shippy-service-consignment/proto/consignment"
	userService "github.com/jipram017/go-ship/shippy-service-user/proto/user"
	vesselProto "github.com/jipram017/go-ship/shippy-service-vessel/proto/vessel"
	"github.com/pkg/errors"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
)

const (
	defaultHost = "mongodb://127.0.0.1:27017"
)

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserService("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	log.Println("masuk6")
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

	log.Println("masuk7")
	consignmentCollection := client.Database("shippy").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", service.Client())
	log.Println("masuk8")

	h := &handler{repository, vesselClient}

	// Register our implementation with
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Println("masuk9")
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Println("masuk10")
		log.Panic(err)
	}
}
