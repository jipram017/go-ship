module github.com/jipram017/go-ship/shippy-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/jipram017/go-ship/shippy-service-vessel v0.0.0-20200621065520-d9cb57377789
	github.com/micro/go-micro/v2 v2.9.0
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.3.4
)
