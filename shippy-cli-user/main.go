package main

import (
	"context"
	"log"
	"os"

	pb "github.com/jipram017/go-ship/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	client := pb.NewUserService("shippy.service.user", srv.Client())
	//name := "Aji Pramono"
	email := "jipram017@gmail.com"
	password := "tohoku2013"
	//company := "NIAGA"

	// r, err := client.Create(context.TODO(), &pb.User{
	// 	Name:     name,
	// 	Email:    email,
	// 	Password: password,
	// 	Company:  company,
	// })

	// if err != nil {
	// 	log.Fatalf("Could not create: %v", err)
	// }

	// log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
