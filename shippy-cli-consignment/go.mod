module github.com/jipram017/go-ship/shippy-cli-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.13

require (
	github.com/jipram017/go-ship/shippy-service-consignment v0.0.0-20200626190434-ccdd3037eb1f
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1-0.20200625102543-5ab475636ad8
	github.com/micro/micro v1.18.0 // indirect
	github.com/micro/micro/v2 v2.9.2-0.20200625165022-867f44aafdf3
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
)
