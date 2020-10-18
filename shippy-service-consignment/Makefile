protobuf: #Meke protobuf file
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto

update-dependencies:
	go get -u github.com/days85/shippy/shippy-service-vessel@main
	go mod tidy