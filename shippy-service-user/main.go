package main

import (
	"log"
	"sync"

	pb "github.com/jipram017/go-ship/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

const schema = `create table if not exists users (
			  id varchar(36) not null,
			  name varchar (125) not null,
			  email varchar(225) not null unique,
			  password varchar(225) not null,
			  company varchar(125),
			  primary key (id)
);`

func main() {
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %V", err)
	}

	// Run schema query on start-up, as we're using "create if not exists"
	// this will only be ran once. In order to create updates, you'll need to
	// use a migrations library
	db.MustExec(schema)

	repo := NewPostgresRepository(db)
	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Get instance of the broker using our default
	publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher, sync.Mutex{}}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Panic(err)
	}
}
