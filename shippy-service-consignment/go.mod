module github.com/days85/shippy/shippy-service-consignment

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/days85/shippy/shippy-service-vessel v0.0.0-20201016163032-1c327c376792
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.4.2 // indirect
	google.golang.org/protobuf v1.25.0
)
