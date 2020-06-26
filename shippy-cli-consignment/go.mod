module github.com/jipram017/go-ship/shippy-cli-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.13

require (
	github.com/go-delve/delve v1.4.1
	github.com/go-log/log v0.2.0 // indirect
	github.com/jipram017/go-ship/shippy-service-consignment v0.0.0-20200626121519-16ee802523b2
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lucas-clemente/quic-go v0.17.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1-0.20200625102543-5ab475636ad8
	github.com/micro/micro v1.18.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/genproto v0.0.0-20200626011028-ee7919e894b5 // indirect
	google.golang.org/grpc v1.30.0
)
