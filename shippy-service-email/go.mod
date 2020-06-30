module github.com/jipram017/go-ship/shippy-service-email

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

go 1.13

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/jipram017/go-ship/shippy-service-user v0.0.0-20200628093047-e865d4b4ac53
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.0
)
