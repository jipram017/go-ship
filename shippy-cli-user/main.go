package main

import (
	"context"
	"log"
	"os"

	pb "github.com/jipram017/go-ship/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.cli.user"),
		micro.Version("latest"),
	)

	service.Init()

	client := pb.NewUserService("go.micro.srv.user", service.Client())
	name := "Henry F"
	email := "ppprmn2@gmail.com"
	password := "apjp21"
	company := "TestCompany"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})

	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	log.Printf("Created: %s", r.User.Id)

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
