module github.com/jipram017/go-ship/shippy-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.13

require (
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jipram017/go-ship/shippy-service-user v0.0.0-20200625183214-ee0485bcb348
	github.com/jipram017/go-ship/shippy-service-vessel v0.0.0-20200625181915-954b432d8aef
	github.com/klauspost/compress v1.10.10 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1-0.20200625102543-5ab475636ad8
	github.com/micro/micro/v2 v2.9.2-0.20200625165022-867f44aafdf3 // indirect
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.3.4
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
)
