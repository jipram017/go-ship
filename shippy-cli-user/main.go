package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/jipram017/go-ship/shippy-service-user/proto/user"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config/source/cli"
)

func createUser(ctx context.Context, service micro.Service, user *proto.User) error {
	client := proto.NewUserService("shippy.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// Print the user
	fmt.Println("Response:", rsp.User)
	return nil
}

func main() {
	// create and initialise a new service
	service := micro.NewService()
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			ctx := context.Background()
			user := &proto.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("error creating user: ", err.Error())
				return err
			}

			return nil
		}),
	)
}